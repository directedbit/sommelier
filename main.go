package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	"github.com/peggyjv/sommelier/v6/app/params"
	corkTypes "github.com/peggyjv/sommelier/v6/x/cork/types"
	"github.com/tendermint/tendermint/store"
	cometbft "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	gravitytypes "github.com/peggyjv/gravity-bridge/module/v3/x/gravity/types"
	//authzTypes "github.com/cosmos/cosmos-sdk/x/authz/types"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	//ibcTypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	stakeTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/writer"
	"go.uber.org/zap"
)

// need to use as raw data as possible, can always do transforms on
// the data later.
type Transaction struct {
	BlockNumber     int64   `parquet:"name=block_number, type=INT64"`
	Time            int64   `parquet:"name=time, type=INT64"`
	MessageType     string  `parquet:"name=message_type, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Source          string  `parquet:"name=source, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Destination     string  `parquet:"name=destination, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Amount          int64   `parquet:"name=amount, type=INT64"`
	OsmoToEthPrice  float32 `parquet:"name=osmo_eth_price, type=FLOAT"`
	SommToOsmoPrice float32 `parquet:"name=somm_osmo_price, type=FLOAT"`
}

func IndexMsg(logger *zap.Logger, pw *writer.ParquetWriter, encodingConfig *params.EncodingConfig, msg sdk.Msg, msgIndex int, height int64, blockTime int64) bool {
	var newTrans = Transaction{}
	switch m := msg.(type) {
	case *bankTypes.MsgSend:
		print("MsgSend ", m.Amount[0].Amount.String(), " of ", m.Amount[0].Denom, " from ", m.FromAddress, " to ", m.ToAddress)
		newTrans = Transaction{BlockNumber: height,
			Time:        blockTime,
			MessageType: "bank.MsgSend",
			Source:      m.FromAddress,
			Destination: m.ToAddress,
			Amount:      m.Amount[0].Amount.Int64(),
		}
	case *bankTypes.MsgMultiSend:
		print("MsgMultiSend ", m.Inputs[0].Coins[0].Amount.String(), " of ", m.Inputs[0].Coins[0].Denom, " from ", m.Inputs[0].Address, " to ", m.Outputs[0].Address)
		for i := range m.Inputs {
			newTrans = Transaction{
				BlockNumber: height,
				Time:        blockTime,
				MessageType: "bank.MsgMultiSend",
				Source:      m.Inputs[i].Address,
				Destination: m.Outputs[i].Address,
				Amount:      m.Inputs[i].Coins[0].Amount.Int64(),
			}

		}
	case *stakeTypes.MsgCreateValidator:
		print("ignoring MsgCreateValidator ", m.Value.Amount.String(), " from ", m.DelegatorAddress, " to ", m.ValidatorAddress)
	case *stakeTypes.MsgEditValidator:
		print("ignoring MsgEditValidator ")
	case *vestingTypes.MsgCreateVestingAccount:
		print("ignoring MsgCreateVestingAccount ", m.Amount.AmountOf("usomm").String(), " from ", m.FromAddress, " to ", m.ToAddress)
	case *transfertypes.MsgTransfer:
		print("IbcMsgTransfer ", m.Token.Amount.String(), " from ", m.Sender, " to ", m.Receiver)
		newTrans = Transaction{BlockNumber: height,
			Time:        blockTime,
			MessageType: "ibc.MsgTransfer",
			Source:      m.Sender,
			Destination: m.Receiver,
			Amount:      m.Token.Amount.Int64(),
		}
	case *distTypes.MsgFundCommunityPool:
		print("MsgFundCommunityPool ", m.Amount.AmountOf("usomm").Int64(), " from ", m.Depositor)
		newTrans = Transaction{BlockNumber: height,
			Time:        blockTime,
			MessageType: "distribution.MsgFundCommunityPool",
			Source:      m.Depositor,
			Destination: "community",
			Amount:      m.Amount.AmountOf("usomm").Int64(),
		}
	case *distTypes.MsgSetWithdrawAddress:
		print("MsgSetWithdrawAddress ", m.DelegatorAddress, " to ", m.WithdrawAddress)
		newTrans = Transaction{BlockNumber: height,
			Time:        blockTime,
			MessageType: "distribution.MsgSetWithdrawAddress",
			Source:      m.DelegatorAddress,
			Destination: m.WithdrawAddress,
			Amount:      0,
		}

	case *distTypes.MsgWithdrawDelegatorReward:
		print("MsgWithdrawDelegatorReward ", m.DelegatorAddress, " amount ", m.Size)
		newTrans = Transaction{BlockNumber: height,
			Time:        blockTime,
			MessageType: "distribution.MsgWithdrawDelegatorReward",
			Source:      m.ValidatorAddress,
			Destination: m.DelegatorAddress,
			Amount:      0,
		}
	case *distTypes.MsgWithdrawValidatorCommission:
		print("MsgWithdrawValidatorCommission ", m.ValidatorAddress)
		newTrans = Transaction{BlockNumber: height,
			Time:        blockTime,
			MessageType: "distribution.MsgWithdrawValidatorCommission",
			Source:      m.ValidatorAddress,
			Destination: m.ValidatorAddress,
			Amount:      0}
	case *stakeTypes.MsgDelegate:
		print("MsgDelegate ", m.Amount.String(), " from ", m.DelegatorAddress, " to ", m.ValidatorAddress)
		newTrans = Transaction{BlockNumber: height,
			Time:        blockTime,
			MessageType: "staking.MsgDelegate",
			Source:      m.DelegatorAddress,
			Destination: m.ValidatorAddress,
			Amount:      m.Amount.Amount.Int64(),
		}
	case *stakeTypes.MsgUndelegate:
		print("MsgUndelegate ", m.Amount.String(), " from ", m.DelegatorAddress, " to ", m.ValidatorAddress)
		newTrans = Transaction{BlockNumber: height,
			Time:        blockTime,
			MessageType: "staking.MsgUndelegate",
			Source:      m.ValidatorAddress,
			Destination: m.DelegatorAddress,
			Amount:      m.Amount.Amount.Int64(),
		}
	case *corkTypes.MsgSubmitCorkRequest:
		print("MsgSubmitCorkRequest ", m.Signer, " from ", m.Cork.TargetContractAddress)
		newTrans = Transaction{
			BlockNumber: height,
			Time:        blockTime,
			MessageType: "cork.MsgSubmitCorkRequest",
			Source:      m.Signer,
			Destination: m.Cork.TargetContractAddress,
			Amount:      0,
		}
	case *authz.MsgRevoke:
		print("Ignoring MsgRevoke ", m.Grantee, " from ", m.Granter, " don't think it will impact price")
	case *authz.MsgGrant:
		print("ignoring MsgGrant ", m.Granter, " to ", m.Grantee, " don't think it will impact price")
	case *stakeTypes.MsgBeginRedelegate:
		print("ignoring MsgBeginRedelegate ", m.Amount.String(), " from ", m.DelegatorAddress, " to ", m.ValidatorDstAddress)
	case *authz.MsgExec:
		wroteValue := false
		for i := range m.Msgs {
			var innerMsg sdk.Msg
			if err := encodingConfig.InterfaceRegistry.UnpackAny(m.Msgs[i], &innerMsg); err != nil {
				logger.Error("failed to unpack msg", zap.Error(err))
				panic(err)
			}
			IndexMsg(logger, pw, encodingConfig, innerMsg, i, height, blockTime)
			if err := pw.Write(newTrans); err != nil {
				log.Fatal("Failed writing to parquet file:", err)
			}
			wroteValue = true
		}
		return wroteValue
	case *gravitytypes.MsgSendToEthereum:
		print("gravity.MsgSendToEthereum ", m.Amount.Amount.String(), " from ", m.Sender, " to ", m.EthereumRecipient)
		newTrans = Transaction{BlockNumber: height,
			Time:        blockTime,
			MessageType: "gravity.MsgSendToEthereum",
			Source:      m.Sender,
			Destination: m.EthereumRecipient,
			Amount:      m.Amount.Amount.Int64(),
		}
	case *gravitytypes.MsgSubmitEthereumTxConfirmation:
		print("ignoring gravity.MsgSubmitEthereumTxConfirmation ", m.Signer)
	case *gravitytypes.MsgSubmitEthereumEvent:
		print("ignoring gravity.MsgSubmiteEthereumEvent ", m.Signer)
	case *gravitytypes.MsgRequestBatchTx:
		print("ignoring gravity.MsgRequestBatchTx ", m.Signer)
	case *gravitytypes.MsgEthereumHeightVote:
		print("ignoring gravity.MsgEthereumHeightVote ", m.Signer)
	case *gravitytypes.MsgDelegateKeys:
		print("ignoring gravity.MsgDelegateKeys ", m.EthereumAddress)
	case *govTypes.MsgDeposit:
		print("ignoring gov.MsgDeposit ", m.Depositor)
	case *govTypes.MsgSubmitProposal:
		newTrans = Transaction{
			BlockNumber: height,
			Time:        blockTime,
			MessageType: "gov.MsgSubmitProposal",
			Source:      m.Proposer,
			Destination: "community",
			Amount:      m.InitialDeposit[0].Amount.Int64(),
		}
	case *govTypes.MsgVote:
		newTrans = Transaction{BlockNumber: height,
			Time:        blockTime,
			MessageType: "gov.MsgVote",
			Source:      m.Voter,
			Destination: "community",
			Amount:      0,
		}
	case *govTypes.MsgVoteWeighted:
		newTrans = Transaction{BlockNumber: height,
			Time:        blockTime,
			MessageType: "gov.MsgVoteWeighted",
			Source:      m.Voter,
			Destination: "community",
			Amount:      m.Options[0].Weight.RoundInt64(),
		}

	/*case *upgradeTypes.SoftwareUpgradeProposal:
		newTrans = Transaction{BlockNumber: height,
			Time:        blockTime,
			MessageType: "upgrade.MsgSoftwareUpgrade",
			Source:      m.Authority,
			Destination: "community",
			Amount:      0,
		}
	case *upgradeTypes.MsgCancelUpgrade:
		newTrans = Transaction{BlockNumber: height,
			Time:        blockTime,
			MessageType: "upgrade.MsgCancelUpgrade",
			Source:      m.Authority,
			Destination: "community",
			Amount:      0,
		}
	*/
	default:
		logger.Fatal("unknown message type", zap.String("type", fmt.Sprintf("%T", m)))
		return false
	}
	if (newTrans != Transaction{}) {
		if err := pw.Write(newTrans); err != nil {
			log.Fatal("Failed writing to parquet file:", err)
		}
		return true
	}
	return false
}

func verifyParquetDB(logger *zap.Logger, deleteAtFinish bool) {

	filename := "/Users/richard/workspace/training-data/test.parquet"
	pw := initParquetWriter(filename)
	data := []Transaction{
		{BlockNumber: 1,
			Time:            2,
			MessageType:     "3",
			Source:          "4",
			Destination:     "5",
			Amount:          6,
			OsmoToEthPrice:  -1.0,
			SommToOsmoPrice: -1.0},
	}

	for _, testTransaction := range data {
		if err := pw.Write(testTransaction); err != nil {
			log.Fatal("Failed writing to parquet file:", err)
		}
	}
	if err := pw.Flush(true); err != nil {
		log.Fatal("Failed flushing to parquet file:", err)
	}
	if err := pw.WriteStop(); err != nil {
		log.Fatal("Failed writing to stop parquet file:", err)
	}
	pw.PFile.Close()

	pr := initParquetReader(filename)
	rows := int(pr.GetNumRows())
	for i := 0; i < rows; i++ {
		logger.Info("reading row", zap.Int("row", i))
		testTransaction := Transaction{}
		transactionRow := make([]Transaction, 1)
		if err := pr.Read(&transactionRow); err != nil {
			log.Fatal("Failed reading:", err)
		}
		for _, testTransaction = range transactionRow {
			logger.Info("read row", zap.Int64("block", testTransaction.BlockNumber), zap.Int64("time", testTransaction.Time), zap.String("message", testTransaction.MessageType), zap.String("source", testTransaction.Source), zap.String("destination", testTransaction.Destination), zap.Int64("amount", testTransaction.Amount))
			testTransaction.OsmoToEthPrice = 12.34
			//testTransaction.OsmoToEthPrice = 12.34
			testTransaction.SommToOsmoPrice = 23.78
		}
	}
	pr.ReadStop()
	if err := pr.PFile.Close(); err != nil {
		log.Fatal("Failed closing file:", err)
	}
	//now write the price data

	if deleteAtFinish {
		if err := os.Remove(filename); err != nil {
			log.Fatal("Failed removing parquet file:", err)
		}
	}
}

func initParquetReader(fileName string) *reader.ParquetReader {
	fr, err := local.NewLocalFileReader(fileName)
	if err != nil {
		log.Fatal("Can't open file", err)
	}
	pr, err := reader.NewParquetReader(fr, new(Transaction), 2)
	if err != nil {
		log.Fatal("Can't create parquet reader", err)
	}
	return pr
}

func initParquetWriter(fileName string) *writer.ParquetWriter {
	fw, err := local.NewLocalFileWriter(fileName)
	if err != nil {
		log.Fatalf("Failed to create Parquet writer: %v", err)
	}

	pw, err := writer.NewParquetWriter(fw, new(Transaction), 2)
	if err != nil {
		log.Fatalf("Failed to create Parquet writer: %v", err)
	}
	pw.RowGroupSize = 128 * 1024 * 1024 // 128M
	pw.CompressionType = parquet.CompressionCodec_SNAPPY
	return pw
}

func getFileWithPrefix(folderPath string, prefix string) string {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		log.Fatalf("Failed to read directory: %s", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue // Skip directories
		}

		// Using file.Name() to get the name of the file and then checking if it starts with the prefix
		if strings.HasPrefix(file.Name(), prefix) {
			fmt.Printf("Found a match: %s\n", file.Name())
			return file.Name()
		}
	}
	return ""
}

func getParquetWriter(rootDirectory string, block *cometbft.Block, logger *zap.Logger, curPw *writer.ParquetWriter, curFile string, curPrefix string) (*writer.ParquetWriter, string, string) {
	//ensure it's a UTC time
	t := time.Unix(block.Time.Unix(), 0).UTC()
	// Format the time to construct the folder and file name

	filePrefix := t.Format("02")
	if curPrefix == filePrefix && curPw != nil && curFile != "" {
		return curPw, curFile, curPrefix
	}
	folderPath := filepath.Join(rootDirectory, t.Format("./2006/01/"))
	// Check if the folder exists
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// Create missing directories
		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			log.Fatalf("Could not create directory: %v", err)
		}
	}
	fileName := getFileWithPrefix(folderPath, filePrefix)
	if fileName == "" {
		fileName = filePrefix + "-" + fmt.Sprint(block.Height) + ".parquet"
	}
	// Combine folder and file names to get the full path
	fullPath := filepath.Join(folderPath, fileName)
	if curPw != nil {
		curPw.Flush(true)
		curPw.WriteStop()
		curPw.PFile.Close()
	}

	pw := initParquetWriter(fullPath)
	return pw, fullPath, filePrefix

}

func indexDbStore(logger *zap.Logger) {
	//storedb, err := dbm.NewGoLevelDB("blockstore", "/Users/richard/workspace/somm-bucket/.sommelier/data")
	//storedb, err := dbm.NewGoLevelDB("blockstore", "/Users/richard/workspace/cometbft")
	encodingConfig := params.MakeCodec()
	storedb, err := dbm.NewGoLevelDB("blockstore", "/Volumes/harry/sommelier")
	partquetFolder := "/Users/richard/workspace/flappy_trade/parquet-training-data"
	if err != nil {
		logger.Error("failed to open blockstore", zap.Error(err))
		log.Fatal(err)
	}
	defer storedb.Close()

	store := store.NewBlockStore(storedb)
	// 25/1/2022 1936108
	// 26/1/2022 1938785
	// 27/1/2022 1955417
	// 2022/02/05-2104778.parquet
	// 2022/02/17-2304669.parquet
	// 2022/02/20-2354455.parquet
	// 2022/02/21-2370941.parquet
	// 2022/03/05-2566098.parquet
	// 2022/03/07-2597957.parquet
	// 2022/03/19-2790436.parquet
	// 2022/04/01-2985179.parquet
	// 2022/04/20-3263356.parquet
	// 2022/04/22-3292600.parquet
	// 2022/05/06-3496897.parquet
	// 2022/05/13-3598939.parquet
	// 2022/06/05-3928061.parquet
	// 2022/06/17-4098228.parquet
	// 2022/07/12-4450435.parquet
	// 2022/08/13-4903515.parquet
	// 2022/09/07-5263456.parquet
	// 2022/10/28-6017364.parquet
	// 2023/02/13-7622094.parquet
	// 2023/03/28-8255283.parquet
	// 2023/11/21-11826331.parquet
	start_block := int64(11826331)
	//end_block := start_block + 200000 //should be about a week
	//end_block := start_block + 30000 //should be about a day
	//end_block := start_block + 1000 //should be about an hour
	end_block := start_block + 5000000
	var pw *writer.ParquetWriter
	var curPwFile string
	var curPrefix string
	for i := start_block; i < end_block; i++ {
		block := store.LoadBlock(i)
		if block == nil {
			logger.Fatal("block is nil")
		} else {
			blockTime := block.Time.Unix()
			pw, curPwFile, curPrefix = getParquetWriter(partquetFolder, block, logger, pw, curPwFile, curPrefix)
			logger.Info("found block", zap.Int64("bh", block.Height), zap.Time("bt", time.Unix(blockTime, 0).UTC()))
			valid_blocks := 0
			for index, tx := range block.Txs {
				sdkTx, err := encodingConfig.TxConfig.TxDecoder()(tx)
				if err != nil {
					// /gravity.v1.SignerSetTxConfirmation
					// /gravity.v1.MsgSubmitEthereumEvent
					//
					if strings.Contains(err.Error(), "/cosmos.slashing.v1beta1.MsgUnjail") {
						logger.Info("Ignoring MsgUnjail")
					} else if strings.Contains(err.Error(), "/gravity.v1.BatchTxConfirmation") {
						logger.Info("Ignoring BatchTxConfirmation")
					} else if strings.Contains(err.Error(), "/ibc.core.client.v1.MsgCreateClient") {
						logger.Info("Ignoring IBC MsgCreateClient")
					} else if strings.Contains(err.Error(), "/ibc.core.client.v1.MsgUpdateClient") {
						logger.Info("Ignoring IBC MsgCreateClient")
					} else if strings.Contains(err.Error(), "ibc.core.channel.v1.MsgAcknowledgement") {
						logger.Info("Ignoring IBC MsgAcknowledgement")
					} else if strings.Contains(err.Error(), "/ibc.core.channel.v1.MsgTimeout") {
						logger.Info("Ignoring IBC MsgTimeout")
						///ibc.core.connection.v1.MsgConnectionOpenInit:
					} else if strings.Contains(err.Error(), "/ibc.core.connection.v1.MsgConnectionOpenInit") {
						logger.Info("Ignoring IBC MsgConnectionOpenInit")
						///ibc.core.channel.v1.MsgChannelOpenInit:
					} else if strings.Contains(err.Error(), "/ibc.core.channel.v1.MsgChannelOpenInit") {
						logger.Info("Ignoring IBC MsgChannelOpenInit")
					} else if strings.Contains(err.Error(), "/ibc.core.channel.v1.MsgRecvPacket:") {
						logger.Info("Ignoring IBC MsgRecvPacket, #TODO inspect the data packet to see if it has amounts etc")
					} else if strings.Contains(err.Error(), "/cosmos.authz.v1beta1.MsgGrant:") {
						logger.Info("Ignoring Authz MsgGrant, authorises  someone to do something on someone else's behalf seems unlikely to impact price")
					} else if strings.Contains(err.Error(), "cosmos.feegrant.v1beta1.MsgGrantAllowance") {
						logger.Info("Ignoring feegrant MsgGrantAllowance")
					} else if strings.Contains(err.Error(), "gravity.v1.CommunityPoolEthereumSpendProposal") {
						logger.Info("Ignoring gravity CommunityPoolEthereumSpendProposal")
					} else if strings.Contains(err.Error(), "gravity.v1.MsgSubmitEthereumTxConfirmation") {
						logger.Info("Ignoring gravity MsgSubmitEthereumTxConfirmation")
						///cosmos.params.v1beta1.ParameterChangeProposal
					} else if strings.Contains(err.Error(), "/cosmos.params.v1beta1.ParameterChangeProposal") {
						logger.Info("Ignoring gravity ParameterChangeProposal")
					} else {
						// TODO application specific txs fail here (e.g. DEX swaps, Akash deployments, etc.)
						fmt.Printf("[Height %d] {%d/%d txs} - Failed to decode tx. Err: %s \n", block.Height, index+1, len(block.Data.Txs), err.Error())
						// gravity.v1.MsgSubmitEthereumTxConfirmation:
						//https://github.com/PeggyJV/gravity-bridge/tree/main/module/proto/gravity/v1
						logger.Fatal("[Height %d] {%d/%d txs} - Failed to decode tx. Err: %s \n", zap.Int64("height", block.Height), zap.Int("index", index+1), zap.Int("total", len(block.Data.Txs)), zap.Error(err))
					}
				} else {
					fmt.Printf("block %d tx %d: %s\n", block.Height, index, tx)
					for msgIndex, msg := range sdkTx.GetMsgs() {
						print(msg)
						if IndexMsg(logger, pw, &encodingConfig, msg, msgIndex, block.Height, blockTime) {
							valid_blocks++
						}
					}
				}
			}
			if valid_blocks == 0 {
				emptyTransaction := Transaction{BlockNumber: block.Height,
					Time:        blockTime,
					MessageType: "empty_block",
					Source:      "",
					Destination: "",
					Amount:      0,
				}
				if err := pw.Write(emptyTransaction); err != nil {
					log.Fatal("Failed writing to parquet file:", err)
				}
			}
		}
		//pw.Flush(true)
	}
	//pw.WriteStop()
}

func main() {
	var logger, _ = zap.NewDevelopment()
	/// read the blockchain myself.
	verifyParquetDB(logger, true)

	indexDbStore(logger)
	//readCorruptDatabase()
	logger.Info("done")

	//lets find if anyone has block 1.
	//var blocks = []int64{10760228, 10760229, 10760230, 10760231, 10760232, 10760233, 10760234, 10760235, 10760236}
	//a := int64(10760228)
	// run the indexer
	//start_index := 1936000
}
