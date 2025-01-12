package flappy_trade

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	//authzTypes "github.com/cosmos/cosmos-sdk/x/authz/types"

	"github.com/tendermint/tendermint/libs/log"
	"go.uber.org/zap"

	//ibcTypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	gravitytypes "github.com/peggyjv/gravity-bridge/module/v3/x/gravity/types"
	"github.com/tendermint/tendermint/store"

	//authzTypes "github.com/cosmos/cosmos-sdk/x/authz/types"
	vestingTypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distTypes "github.com/cosmos/cosmos-sdk/x/distribution/types"

	govTypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	corkTypes "github.com/peggyjv/sommelier/v6/x/cork/types"

	//ibcTypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	stakeTypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	"github.com/peggyjv/sommelier/v6/app/params"
	cometbft "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/parquet"
	"github.com/xitongsys/parquet-go/reader"
	"github.com/xitongsys/parquet-go/writer"
	//"go.uber.org/zap"
)

var ParquetFolder string

const SOMM_OSMO_POOL = 627
const OSMO_ETH_POOL = 704

var encodingConfig params.EncodingConfig
var blockstore *store.BlockStore
var curPw *writer.ParquetWriter
var curPwFile string
var curPrefix string
var storedb *dbm.GoLevelDB

func init() {
	userHome, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	ParquetFolder = path.Join(userHome, "workspace", "flappy_trade", "parquet-training-data")
	encodingConfig = params.MakeCodec()
}

/*
*[

	{
	  "symbol": "SOMM",
	  "amount": 383141.229872,
	  "denom": "ibc/9BBA9A1C257E971E38C1422780CE6F0B0686F0A3085E2D61118D904BFE0F5F5E",
	  "coingecko_id": "sommelier",
	  "liquidity": 149644.314156042,
	  "liquidity_24h_change": 12.47328510471743,
	  "volume_24h": 56001.03135882975,
	  "volume_24h_change": 263.5760991782874,
	  "price": 0.1952861028021594,
	  "price_24h_change": 20.487449847446115,
	  "fees": "0.2%"
	},
	{
	  "symbol": "OSMO",
	  "amount": 98205.172454,
	  "denom": "uosmo",
	  "coingecko_id": "osmosis",
	  "liquidity": 149644.314156042,
	  "liquidity_24h_change": 12.47328510471743,
	  "volume_24h": 56001.03135882975,
	  "volume_24h_change": 263.5760991782874,
	  "price": 0.76189629,
	  "price_24h_change": 4.982722824158406,
	  "fees": "0.2%"
	}

]
*/

type Asset struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	// Include other fields if necessary
}

func GetPoolPrice(poolId int64, baseAsset string, denomAsset string) float64 {
	url := fmt.Sprintf("https://api-osmosis.imperator.co/pools/v2/%d", poolId)

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// Add header to the request: 'accept: application/json'
	req.Header.Add("accept", "application/json")

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var assets []Asset
	err = json.Unmarshal([]byte(body), &assets)
	if err != nil {
		return 0.0
	}

	// Initialize variables to store prices
	var baseAssetPrice, denomAssetPrice float64

	// Iterate over the slice to find the prices
	for _, asset := range assets {
		switch asset.Symbol {
		case baseAsset:
			baseAssetPrice = asset.Price
		case denomAsset:
			denomAssetPrice = asset.Price
		}
	}

	// Calculate the SOMM/OSMO ratio
	if denomAssetPrice != 0 { // Ensure osmoPrice is not zero to avoid division by zero
		baseToDenomRatio := baseAssetPrice / denomAssetPrice
		fmt.Printf("%s/%s: %f\n", baseAsset, denomAsset, baseToDenomRatio)
		return baseToDenomRatio
	} else {
		fmt.Println("OSMO price is zero, cannot calculate ratio")
	}
	return 0.0

}

func EncodeMsg(logger log.Logger, encodingConfig *params.EncodingConfig, msg sdk.Msg, msgIndex int, height int64, blockTime int64, somm_to_osmo float64, osmo_to_eth float64) []Transaction {
	var foundTransactions []Transaction
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
		for i := range m.Msgs {
			var innerMsg sdk.Msg
			if err := encodingConfig.InterfaceRegistry.UnpackAny(m.Msgs[i], &innerMsg); err != nil {
				logger.Error("failed to unpack msg", zap.Error(err))
				panic(err)
			}
			newTransactions := EncodeMsg(logger, encodingConfig, innerMsg, i, height, blockTime, somm_to_osmo, osmo_to_eth)
			foundTransactions = append(foundTransactions, newTransactions...)
		}
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
		emsg := fmt.Sprintf("unknown message type %T", m)
		logger.Error(emsg)
		panic(emsg)
	}
	if (newTrans != Transaction{}) {
		newTrans.SommToOsmoPrice = float32(somm_to_osmo)
		newTrans.OsmoToEthPrice = float32(osmo_to_eth)

		foundTransactions = append(foundTransactions, newTrans)
	}
	return foundTransactions
}

// need to use as raw data as possible, can always do transforms on
// the data later.
type Transaction struct {
	BlockNumber     int64   `json:"block_number" parquet:"name=block_number, type=INT64"`
	Time            int64   `json:"time" parquet:"name=time, type=INT64"`
	MessageType     string  `json:"message_type" parquet:"name=message_type, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Source          string  `json:"source" parquet:"name=source, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Destination     string  `json:"destination" parquet:"name=destination, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Amount          int64   `json:"amount" parquet:"name=amount, type=INT64"`
	OsmoToEthPrice  float32 `json:"osmo_eth_price" parquet:"name=osmo_eth_price, type=FLOAT"`
	SommToOsmoPrice float32 `json:"somm_osmo_price" parquet:"name=somm_osmo_price, type=FLOAT"`
}

func IndexBlockHeight(height int64, somm_to_osmo float64, osmo_to_eth float64, logger log.Logger) {
	logger.Error("Unfinished / not working")
	panic("Unfinished")

	/*
		storedb, err = dbm.NewGoLevelDB("blockstore", BlockStoreFolder)
		if err != nil {
			panic(err)
		}
		blockstore = store.NewBlockStore(storedb)
		block := blockstore.LoadBlock(height)
		pw, curPwFile, curPrefix = GetParquetWriter(ParquetFolder, block, logger, pw, curPwFile, curPrefix)
		IndexBlock(block, somm_to_osmo, osmo_to_eth, &encodingConfig, pw, logger)
		if err := pw.Flush(true); err != nil {
			log.Fatal("Failed flushing to parquet file:", err)
		}
		if err := pw.WriteStop(); err != nil {
			log.Fatal("Failed writing to stop parquet file:", err)
		}
		pw.PFile.Close()
	*/
}

func IndexTransactions(transactions []Transaction, block *cometbft.Block, logger log.Logger) {
	curPw, curPwFile, curPrefix = GetParquetWriter(ParquetFolder, block, logger, curPw, curPwFile, curPrefix)
	for _, transaction := range transactions {
		if err := curPw.Write(transaction); err != nil {
			panic(err)
		}
	}
}

func EncodeBlockTransactions(block *cometbft.Block, somm_to_osmo float64, osmo_to_eth float64, logger log.Logger) []Transaction {
	var foundTransactions []Transaction
	//blocktime is in UTC time
	blockTime := block.Time.UTC().Unix()
	blockTimeString := block.Time.UTC().Format("2006-01-02 15:04:05")
	logger.Info("Indexig block ", " height ", block.Height, "time ", blockTimeString)
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
				emsg := fmt.Sprintf("[Height %d] {%d/%d txs} - Failed to decode tx. Err: %s \n", block.Height, index+1, len(block.Data.Txs), err.Error())
				// gravity.v1.MsgSubmitEthereumTxConfirmation:
				//https://github.com/PeggyJV/gravity-bridge/tree/main/module/proto/gravity/v1
				logger.Error(emsg)
				panic(emsg)
			}
		} else {
			fmt.Printf("block %d tx %d: %s\n", block.Height, index, tx)
			for msgIndex, msg := range sdkTx.GetMsgs() {
				print(msg)
				newTransactions := EncodeMsg(logger, &encodingConfig, msg, msgIndex, block.Height, blockTime, somm_to_osmo, osmo_to_eth)
				foundTransactions = append(foundTransactions, newTransactions...)
			}
		}
	}
	if len(foundTransactions) == 0 {
		emptyTransaction := Transaction{BlockNumber: block.Height,
			Time:            blockTime,
			MessageType:     "empty_block",
			Source:          "",
			Destination:     "",
			Amount:          0,
			SommToOsmoPrice: float32(somm_to_osmo),
			OsmoToEthPrice:  float32(osmo_to_eth),
		}

		foundTransactions = append(foundTransactions, emptyTransaction)
	}
	return foundTransactions
}

func getFileWithPrefix(folderPath string, prefix string) string {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		emsg := fmt.Sprintf("Failed to read directory: %s", err)
		panic(emsg)
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

func initParquetReader(fileName string) *reader.ParquetReader {
	fr, err := local.NewLocalFileReader(fileName)
	if err != nil {
		panic(err)
	}
	pr, err := reader.NewParquetReader(fr, new(Transaction), 2)
	if err != nil {
		panic(err)
	}
	return pr
}

func initParquetWriterDeletingExistingData(fileName string) *writer.ParquetWriter {
	// Create or overwrite the file
	fw, err := local.NewLocalFileWriter(fileName)
	if err != nil {
		panic(err)
	}
	pw, err := writer.NewParquetWriter(fw, new(Transaction), 2)
	if err != nil {
		panic(err)
	}
	pw.RowGroupSize = 128 * 1024 * 1024 // 128M
	pw.CompressionType = parquet.CompressionCodec_SNAPPY
	return pw
}

func initParquetWriterKeepingExistingData(fileName string) *writer.ParquetWriter {
	var transactions []*Transaction
	// Check if file exists and has data
	if fileInfo, err := os.Stat(fileName); err == nil && fileInfo.Size() > 0 {
		fr, err := local.NewLocalFileReader(fileName)
		if err != nil {
			panic(err)
		}
		defer fr.Close()

		pr, err := reader.NewParquetReader(fr, new(Transaction), 2)
		//most likely file is opened but never shut properly.
		if err != nil {
			//likely invalid format
			// Create or overwrite the file
			fmt.Println("Error reading file, will overwrite", err)
		} else {
			defer pr.ReadStop()
			num := int(pr.GetNumRows())
			for i := 0; i < num; i++ {
				transaction := make([]*Transaction, 1)
				if err = pr.Read(&transaction); err != nil {
					panic(err)
				}
				transactions = append(transactions, transaction...)
			}
		}
	}
	pw := initParquetWriterDeletingExistingData(fileName)
	// Write existing data
	for _, transaction := range transactions {
		if err := pw.Write(transaction); err != nil {
			panic(err)
		}
	}
	return pw
}

func CloseCurrentParquetFile() {
	CloseParquetFile(curPw)
}

func CloseParquetFile(pw *writer.ParquetWriter) {
	if pw != nil {
		pw.Flush(true)
		pw.WriteStop()
		pw.PFile.Close()
	}
}

func GetParquetWriter(rootDirectory string, block *cometbft.Block, logger log.Logger, maybePw *writer.ParquetWriter, curFile string, curPrefix string) (*writer.ParquetWriter, string, string) {
	//ensure it's a UTC time
	t := time.Unix(block.Time.Unix(), 0).UTC()
	//t := time.Unix(blockTime.Unix(), 0).UTC()
	// Format the time to construct the folder and file name

	filePrefix := t.Format("02")
	if curPrefix == filePrefix && maybePw != nil && curFile != "" {
		return maybePw, curFile, curPrefix
	}
	folderPath := filepath.Join(rootDirectory, t.Format("./2006/01/"))
	// Check if the folder exists
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// Create missing directories
		if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
			logger.Error("Could not create directory: %v", err)
			panic(err)
		}
	}
	//check if the file exists
	fileName := getFileWithPrefix(folderPath, filePrefix)
	if fileName == "" {
		fileName = filePrefix + "-" + fmt.Sprint(block.Height) + ".parquet"
		// Combine folder and file names to get the full path
		fullPath := filepath.Join(folderPath, fileName)
		CloseParquetFile(maybePw)
		pw := initParquetWriterDeletingExistingData(fullPath)
		return pw, fullPath, filePrefix
	} else {
		fullPath := filepath.Join(folderPath, fileName)
		CloseParquetFile(maybePw)
		pw := initParquetWriterDeletingExistingData(fullPath)
		return pw, fullPath, filePrefix
	}

}
