package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"path"
	"reflect"
	"strings"
	"time"

	//"github.com/cosmos/cosmos-sdk/server/config"
	"github.com/peggyjv/sommelier/v6/flappy_trade"

	"github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/store"
	"github.com/tendermint/tendermint/version"
	dbm "github.com/tendermint/tm-db"

	//authzTypes "github.com/cosmos/cosmos-sdk/x/authz/types"

	//ibcTypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	bc "github.com/tendermint/tendermint/blockchain"
	blocksync "github.com/tendermint/tendermint/blockchain/v0"
	bcproto "github.com/tendermint/tendermint/proto/tendermint/blockchain"
	"github.com/tendermint/tendermint/types"

	tmlog "github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/p2p"
	sm "github.com/tendermint/tendermint/state"
	"go.uber.org/zap"
)

var (
	IndexBlockWhileSyncing = false
)

const (
	maxTotalRequesters = 10

	trySyncIntervalMS = 10

	// stop syncing when last block's time is
	// within this much of the system time.
	// stopSyncingDurationMinutes = 10

	// ask for best height every 10s
	statusUpdateIntervalSeconds = 10
	// check if we should switch to consensus reactor
	switchToConsensusIntervalSeconds = 1
)

// Reactor handles long-term catchup syncing.
type MyReactor struct {
	p2p.BaseReactor

	// immutable
	//initialState sm.State

	blockExec *sm.BlockExecutor
	store     *store.BlockStore
	pool      *blocksync.BlockPool
	fastSync  bool

	requestsCh <-chan blocksync.BlockRequest
	errorsCh   <-chan PeerError
}

type PeerError struct {
	err    error
	peerID p2p.ID
}

func NewMyReactor(state sm.State, blockExec *sm.BlockExecutor, store *store.BlockStore,
	fastSync bool, logger tmlog.Logger) *MyReactor {

	if state.LastBlockHeight != store.Height() {
		logger.Error(fmt.Sprintf("state (%v) and store (%v) height mismatch, will ignore as only care about blocks for fast sync", state.LastBlockHeight,
			store.Height()))
	}

	requestsCh := make(chan blocksync.BlockRequest, 1)

	//const capacity = 1000                                // must be bigger than peers count
	//errorsCh := make(chan blocksync.PeerError, capacity) // so we don't block in #Receive#pool.AddBlock

	startHeight := store.Height() + 1
	if startHeight == 1 {
		startHeight = state.InitialHeight
	}
	pool := blocksync.NewBlockPool(startHeight, requestsCh, nil)

	bcR := &MyReactor{
		blockExec:  blockExec,
		store:      store,
		pool:       pool,
		fastSync:   fastSync,
		requestsCh: requestsCh,
		//errorsCh:     errorsCh,
	}
	bcR.BaseReactor = *p2p.NewBaseReactor("Reactor", bcR)
	return bcR
}

// AddPeer implements Reactor by sending our state to peer.
// this doesn't immediately add the peer, this is done later
func (MyR *MyReactor) AddPeer(peer p2p.Peer) {
	MyR.Logger.Info("Add Peer", "peer", peer)
	msg, err := bc.EncodeMsg(&bcproto.StatusResponse{
		Base:   MyR.store.Base(),
		Height: MyR.store.Height(),
	})
	if err != nil {
		MyR.Logger.Error("could not convert msg to proto", "err", err)
	}
	peer.Send(blocksync.BlockchainChannel, msg)
	// it's OK if send fails. will try later in poolRoutine

	// peer is added to the pool once we receive the first
	// bcStatusResponseMessage from the peer and call pool.SetPeerRange
}

func (MyR *MyReactor) OnStart() error {
	err := MyR.pool.Start()
	if err != nil {
		return err
	}
	go MyR.poolRoutine()
	return nil
}

// BroadcastStatusRequest broadcasts `BlockStore` base and height.
func (MyR *MyReactor) BroadcastStatusRequest() error {
	bm, err := bc.EncodeMsg(&bcproto.StatusRequest{})
	if err != nil {
		MyR.Logger.Error("could not convert msg to proto", "err", err)
		return fmt.Errorf("could not convert msg to proto: %w", err)
	}

	MyR.Switch.Broadcast(blocksync.BlockchainChannel, bm)

	return nil
}

// respondToPeer loads a block and sends it to the requesting peer,
// if we have it. Otherwise, we'll respond saying we don't have it.
func (MyR *MyReactor) respondToPeer(msg *bcproto.BlockRequest,
	src p2p.Peer) (queued bool) {
	MyR.Logger.Info("Pretend we don't have the block ", "src", src, "height", msg.Height)

	msgBytes, err := bc.EncodeMsg(&bcproto.NoBlockResponse{Height: msg.Height})
	if err != nil {
		MyR.Logger.Error("could not convert msg to protobuf", "err", err)
		return false
	}

	return src.TrySend(blocksync.BlockchainChannel, msgBytes)
}

func (MyR *MyReactor) poolRoutine() {

	trySyncTicker := time.NewTicker(trySyncIntervalMS * time.Millisecond)
	defer trySyncTicker.Stop()

	statusUpdateTicker := time.NewTicker(statusUpdateIntervalSeconds * time.Second)
	defer statusUpdateTicker.Stop()

	switchToConsensusTicker := time.NewTicker(switchToConsensusIntervalSeconds * time.Second)
	defer switchToConsensusTicker.Stop()

	blocksSynced := uint64(0)

	lastHundred := time.Now()
	lastRate := 0.0

	didProcessCh := make(chan struct{}, 1)

	go func() {
		for {
			select {
			case <-MyR.Quit():
				return
			case <-MyR.pool.Quit():
				return
			case request := <-MyR.requestsCh:
				peer := MyR.Switch.Peers().Get(request.PeerID)
				if peer == nil {
					continue
				}
				msgBytes, err := bc.EncodeMsg(&bcproto.BlockRequest{Height: request.Height})
				if err != nil {
					MyR.Logger.Error("could not convert msg to proto", "err", err)
					continue
				}

				queued := peer.TrySend(blocksync.BlockchainChannel, msgBytes)
				if !queued {
					MyR.Logger.Debug("Send queue is full, drop block request", "peer", peer.ID(), "height", request.Height)
				}
			case err := <-MyR.errorsCh:
				peer := MyR.Switch.Peers().Get(err.peerID)
				if peer != nil {
					MyR.Switch.StopPeerForError(peer, err)
				}

			case <-statusUpdateTicker.C:
				// ask for status updates
				MyR.Logger.Info("Handling status request")
				go MyR.BroadcastStatusRequest() //nolint: errcheck

			}
		}
	}()

FOR_LOOP:
	for {
		select {
		case <-switchToConsensusTicker.C:
			/*height, numPending, lenRequesters := bcR.pool.GetStatus()
			outbound, inbound, _ := bcR.Switch.NumPeers()
			bcR.Logger.Debug("Consensus ticker", "numPending", numPending, "total", lenRequesters,
				"outbound", outbound, "inbound", inbound)
			if bcR.pool.IsCaughtUp() {
				bcR.Logger.Info("Time to switch to consensus reactor!", "height", height)
				if err := bcR.pool.Stop(); err != nil {
					bcR.Logger.Error("Error stopping pool", "err", err)
				}
				conR, ok := bcR.Switch.Reactor("CONSENSUS").(consensusReactor)
				if ok {
					conR.SwitchToConsensus(state, blocksSynced > 0 || stateSynced)
				}
				// else {
				// should only happen during testing
				// }

				break FOR_LOOP
			}
			*/
			if MyR.pool.IsCaughtUp() {
				MyR.Logger.Error("Caught up!")
			}

		case <-trySyncTicker.C: // chan time
			select {
			case didProcessCh <- struct{}{}:
			default:
			}

		case <-didProcessCh:
			// NOTE: It is a subtle mistake to process more than a single block
			// at a time (e.g. 10) here, because we only TrySend 1 request per
			// loop.  The ratio mismatch can result in starving of blocks, a
			// sudden burst of requests and responses, and repeat.
			// Consequently, it is better to split these routines rather than
			// coupling them as it's written here.  TODO uncouple from request
			// routine.

			// See if there are any blocks to sync.
			first, _ := MyR.pool.PeekTwoBlocks()
			//ignoring second block, don't need to validate
			if first == nil { // || second == nil {
				// We need both to sync the first block.
				continue FOR_LOOP
			} else {
				// Try again quickly next loop.
				didProcessCh <- struct{}{}
			}

			firstParts := first.MakePartSet(types.BlockPartSizeBytes)
			//firstPartSetHeader := firstParts.Header()
			//firstID := types.BlockID{Hash: first.Hash(), PartSetHeader: firstPartSetHeader}
			// Finally, verify the first block using the second's commit
			// NOTE: we can probably make this more efficient, but note that calling
			// first.Hash() doesn't verify the tx contents, so MakePartSet() is
			// currently necessary.

			MyR.pool.PopRequest()

			MyR.store.SaveBlock(first, firstParts, first.LastCommit)
			MyR.Logger.Info("Flappy Trader Storring", "BlockHeight", first.Height)
			somm_to_osmo := 0.0
			osmo_to_eth := 0.0
			if MyR.pool.IsCaughtUp() {
				//Get the osmosis prices
				//get the osmosis price
				//eth osmo
				//osmo somm
				somm_to_osmo = flappy_trade.GetPoolPrice(flappy_trade.SOMM_OSMO_POOL, "SOMM", "OSMO")
				MyR.Logger.Info("SOMM_OSMO_PRICE", somm_to_osmo)
				osmo_to_eth = flappy_trade.GetPoolPrice(flappy_trade.OSMO_ETH_POOL, "OSMO", "ETH")
				MyR.Logger.Info("ETH_OSMO_PRICE")
			}
			foundTransactions := flappy_trade.EncodeBlockTransactions(first, somm_to_osmo, osmo_to_eth, MyR.Logger)
			//if len(foundTransactions) > 0 && MyR.pool.IsCaughtUp() {
			//Now do the trading..
			if len(foundTransactions) > 0 && MyR.pool.IsCaughtUp() {
				for _, transaction := range foundTransactions {
					if flappy_trade.ShouldSell(&transaction, "http://127.0.0.1:5000/predict") {
						MyR.Logger.Error("SELL SELL SELLL SELLLLLL NOW!!!! -  Transaction", transaction)
					} else {
						MyR.Logger.Error("Hoooooooollld. ", transaction)
					}
				}
			} else {
				MyR.Logger.Info("Not caught up yet / no transactions found, not checking for transactions")
			}
			if IndexBlockWhileSyncing {
				MyR.Logger.Error("Not indexing block")
				flappy_trade.IndexTransactions(foundTransactions, first, MyR.Logger)
			} else {
				MyR.Logger.Error("Not indexing block")
			}

			blocksSynced++

			if blocksSynced%100 == 0 {
				lastRate = 0.9*lastRate + 0.1*(100/time.Since(lastHundred).Seconds())
				MyR.Logger.Info("Fast Sync Rate",
					"max_peer_height", MyR.pool.MaxPeerHeight(), "blocks/s", lastRate)
				lastHundred = time.Now()
			}

			continue FOR_LOOP

		case <-MyR.Quit():
			break FOR_LOOP
		}
	}
}

// OnStop implements service.Service.
func (MyR *MyReactor) OnStop() {
	if MyR.fastSync {
		if err := MyR.pool.Stop(); err != nil {
			MyR.Logger.Error("Error stopping pool", "err", err)
		}
	}
}

// Receive implements Reactor by handling 4 types of messages (look below).
func (MyR *MyReactor) Receive(chID byte, src p2p.Peer, msgBytes []byte) {
	msg, err := bc.DecodeMsg(msgBytes)
	if err != nil {
		MyR.Logger.Error("Error decoding message", "src", src, "chId", chID, "err", err)
		MyR.Switch.StopPeerForError(src, err)
		return
	}

	if err = bc.ValidateMsg(msg); err != nil {
		MyR.Logger.Error("Peer sent us invalid msg", "peer", src, "msg", msg, "err", err)
		MyR.Switch.StopPeerForError(src, err)
		return
	}

	MyR.Logger.Debug("Receive", "src", src, "chID", chID, "msg", msg)

	switch msg := msg.(type) {
	case *bcproto.BlockRequest:
		MyR.Logger.Info("Block request, will send don't have")
		MyR.respondToPeer(msg, src)
	case *bcproto.BlockResponse:
		bi, err := types.BlockFromProto(msg.Block)
		if err != nil {
			MyR.Logger.Error("Block content is invalid", "err", err)
			return
		}
		MyR.pool.AddBlock(src.ID(), bi, len(msgBytes))
	case *bcproto.StatusRequest:
		// Send peer our state.
		msgBytes, err := bc.EncodeMsg(&bcproto.StatusResponse{
			Height: MyR.store.Height(),
			Base:   MyR.store.Base(),
		})
		if err != nil {
			MyR.Logger.Error("could not convert msg to protobut", "err", err)
			return
		}
		src.TrySend(blocksync.BlockchainChannel, msgBytes)
	case *bcproto.StatusResponse:
		// Got a peer status. Unverified.
		MyR.pool.SetPeerRange(src.ID(), msg.Base, msg.Height)
	case *bcproto.NoBlockResponse:
		MyR.Logger.Debug("Peer does not have requested block", "peer", src, "height", msg.Height)
	default:
		MyR.Logger.Error(fmt.Sprintf("Unknown message type %v", reflect.TypeOf(msg)))
	}
}

// GetChannels implements Reactor
func (MyR *MyReactor) GetChannels() []*p2p.ChannelDescriptor {
	return []*p2p.ChannelDescriptor{
		{
			ID:                  blocksync.BlockchainChannel,
			Priority:            5,
			SendQueueCapacity:   1000,
			RecvBufferCapacity:  50 * 4096,
			RecvMessageCapacity: bc.MaxMsgSize,
		},
	}
}

func fastSync(tmlogger tmlog.Logger, homeFolder string, dataFolder string) {
	fmt.Println("Starting fast Reactor")
	var configFolder = path.Join(homeFolder, "config")
	var genesisFile = path.Join(configFolder, "genesis.json")
	var nodeKeyFile = path.Join(configFolder, "node_key.json")
	var logger, _ = zap.NewDevelopment()
	/// read the blockchain myself.
	//indexDbStore(logger)
	//readCorruptDatabase()
	//logger.Info("done")
	storeDB, err := dbm.NewGoLevelDB("blockstore", dataFolder)
	if err != nil {
		logger.Error("failed to open blockstore", zap.Error(err))
		log.Fatal(err)
	}
	myerStore := store.NewBlockStore(storeDB)
	logger.Info("Opened Blockstore", zap.Int64("height", myerStore.Height()), zap.Int64("base", myerStore.Base()))
	//TODO - maybe discard the responses
	nswStore := sm.NewStore(storeDB, sm.StoreOptions{DiscardABCIResponses: false})
	stateStore, err := nswStore.LoadFromDBOrGenesisFile(genesisFile)
	if err != nil {
		logger.Error("failed to load state store", zap.Error(err))
		panic(err)
	}
	config := config.DefaultConfig()
	mConnConfig := p2p.MConnConfig(config.P2P)
	nodeKey, err := p2p.LoadOrGenNodeKey(nodeKeyFile)
	if err != nil {
		logger.Error("failed to load node key", zap.Error(err))
		panic(err)
	}

	myreactor := NewMyReactor(stateStore, nil, myerStore, true, tmlogger)
	myreactor.SetLogger(tmlogger)

	nodeInfo := p2p.DefaultNodeInfo{
		ProtocolVersion: p2p.NewProtocolVersion(
			version.P2PProtocol, // global
			stateStore.Version.Consensus.Block,
			stateStore.Version.Consensus.App,
			//state.Version.Consensus.Block,
			//state.Version.Consensus.App,
		),
		DefaultNodeID: nodeKey.ID(),
		Network:       "sommelier-3",
		Version:       version.TMCoreSemVer,
		Channels: []byte{
			blocksync.BlockchainChannel,
			// cs.StateChannel, cs.DataChannel, cs.VoteChannel, cs.VoteSetBitsChannel,
			// mempl.MempoolChannel,
			// evidence.EvidenceChannel,
			// statesync.SnapshotChannel, statesync.ChunkChannel,
		},
		Moniker: config.Moniker,
		Other: p2p.DefaultNodeInfoOther{
			TxIndex:    "off",
			RPCAddress: config.RPC.ListenAddress,
		},
		ListenAddr: config.P2P.ListenAddress,
	}

	err = nodeInfo.Validate()
	if err != nil {
		logger.Error("failed to validate node info", zap.Error(err))
		panic(err)
	}

	transport := p2p.NewMultiplexTransport(nodeInfo, *nodeKey, mConnConfig)
	transport.AddChannel(blocksync.BlockchainChannel)
	//lets find if anyone has block 1.
	//var blocks = []int64{10760228, 10760229, 10760230, 10760231, 10760232, 10760233, 10760234, 10760235, 10760236}
	//a := int64(10760228)
	// run the indexer
	//start_index := 1936000
	//peersString := "590de0c91ce7c6ca1dbdf32c75873617b5503f9f@65.108.101.4:29886,302cb1aa66e2ce8be9a766b6e40d98747cdcf906@65.109.90.96:60656,3cfe6e50a07cce40881a5648c54a823769e9bf47@135.181.129.122:26656,ad242fe3d7e5abffcae24af91fdcbd004daa3d0f@65.21.84.223:14156,04def519ec8b578bc6d9efaa56eb957a131cf6f3@135.181.192.166:26656,7e041096585d12b99d9f575979744251409dd0a0@93.115.25.106:36656,292b84d3b7cd6e3b081f39e6721a5c3620f19300@51.195.6.227:26658,75be8f90f2aa1da26744ca52ce38aff2d806df26@213.239.207.165:20986,362d69cdada0d5926181c602c40f7a7d97186d87@157.90.95.15:26656,53f2d2169c9febf4f43d281c354fc51a83e0b294@5.9.99.172:14156,733ee0a5a74de20745c5236df015642aa53060bc@57.128.96.155:14156,281bca9514fe002d90a4b3db2a1567f66a0ec357@159.69.66.149:26646,50dd2f2f9ab1c1ddbc2d67bedeed768d7227d331@46.4.107.112:14156,d4ff71a5cca634651ff1580ea368bc8e786d3102@212.102.34.6:36656,0d67a804aae86f5bb60041abb0e7ce15644167ba@95.217.144.107:14156,8b08753a36812d7e6d8799e6afdebe9b05119ba7@95.211.66.25:13656,86cbeedea81c89192f5a57f2916c4c89e0f2686a@46.166.140.167:26656,97ede4e6d2300abe224af6d1d86e62a5f9c4d6b8@69.197.6.24:22856,15861ad4a98ee9a872f85fdace2b23422c796cdf@146.59.81.92:36656,ff826c4b26ddfb7d5d79612585624115712b0a27@80.64.208.52:26656,03d22833a8b8e414c2da861cab1cbe8f5694063b@5.78.93.29:26656,db1748aee7658ce2614e27acb38d184654027194@35.245.10.16:26656,75c7972462772278129cb6585f63b94fa4ee81a7@44.196.14.252:26656,b690d37718edd4000c11c80addf5565b72e08078@35.188.241.180:26656,de0754813904e926583afe4fc83c07173316e8b2@186.233.186.53:26656,aad751c8a0b3e887fc75ec9e11c85853c93ce8ce@65.108.68.86:11056,2c390272eb51d00f2cb312b31471adc8a27b76d4@84.244.95.233:26656,7bb0b626fe74992f07447e5f9239a188317517a7@45.83.104.218:2000,1f3736062f874d214e132084644c6a3f357936de@35.186.165.222:26656,ad7312bcf1beac7ceb817cb6a47b6697d63daf5e@34.134.46.169:26656,3fb643a2e4bcd9e79ea745c0e61a06fa3ab99ecf@15.235.87.217:26656,3cade0e5deb5bbd8adfa98f62eca48b7d8f87414@34.16.142.53:26656,6e42ef37b18bd6547c4ce7a76111d5377c25e684@34.145.250.69:26656,ed7ef7587d0d7552e973d6fa480ce0ff7f9f39ee@34.70.157.247:26656,24a9d13e9e95ecf030253b64c1ef6f283f37d155@34.125.175.228:26656,e15ca4ba5ee5dd515ba73c53737bb3f6deb051a1@65.108.121.190:2130,4b301813e7db41ed3c8109994e18c0e701b29acd@35.185.11.168:26656,be0f0de2c77ad41401711d6a395dd6c6440ef124@144.217.158.217:26656,a919153f11e86a0d80672fefea7e845b9bd1bbde@65.19.136.133:23356,c08eced51e376d9314b313cbece53b1059931ba6@65.108.238.103:14156,47f79f4296e898a40254bcbffc904d6f0a8a55cc@34.86.193.165:26656,a8fa97fbc3187b4ffbb25eaff7636a88944145de@65.109.117.74:26656,38955171f30ec1c2817673c84012a9d723623862@116.202.208.214:2620,f3251d67547f88595966101562094ca538d39f8e@65.108.106.135:26696,4d95665a3b60b478c4191c090a1e8ba5bfda448c@139.177.182.181:26656,ee74b3c6d5037e2c56bc95b989b0a26643d79aa4@81.0.220.94:24556,402902424e92850107cd0d5b35fcf40b54256b91@93.151.164.241:26656,a53131c0903105f7107b44bcee337ac1e7e4e260@148.113.6.121:34656,d366022d8505e328ee4f0b9b5b8fde210013d178@65.108.230.46:14156,ef6d3bcfee910e1743e9076848597956e98c5d9b@155.133.22.9:24556,ebc272824924ea1a27ea3183dd0b9ba713494f83@95.214.52.139:27176,c94fd60124e3656df54ff965d178e36c760c195d@65.108.57.224:26656,9ffc9ade220b0a0fc0c9abded0b032f94f3cbb38@65.108.109.103:10656,e6c0d56e8c220b835eb5f4a2786cbcd4ea56696a@107.150.119.135:26656,8663e7da214559e6dd094a1dde7b877cf9571b6f@37.59.18.132:20002,fdc824a6662e0d1a0265269b13d2bcccb663fb8f@195.20.239.232:3080,20e1000e88125698264454a884812746c2eb4807@65.108.227.217:14156"
	//peersString := "ae3df68c1ac6b7a493504e2871b0ab26e1b9c542@38.109.200.33:26656,f3251d67547f88595966101562094ca538d39f8e@65.108.106.135:26696,302cb1aa66e2ce8be9a766b6e40d98747cdcf906@65.109.90.96:60656,86cbeedea81c89192f5a57f2916c4c89e0f2686a@46.166.140.167:26656,3cfe6e50a07cce40881a5648c54a823769e9bf47@135.181.129.122:26656,f40f1b40032802bc7552f2d88b29523d0da7a948@34.79.14.20:26656,aad751c8a0b3e887fc75ec9e11c85853c93ce8ce@65.108.68.86:11056,222e3b421288e46d8ae30532aef9458741dcc531@51.91.152.102:20003,7e041096585d12b99d9f575979744251409dd0a0@93.115.25.106:36656,e8e3359cfe774d0cad774a0bf34f53cc5b872703@95.216.161.87:26656,53f2d2169c9febf4f43d281c354fc51a83e0b294@5.9.99.172:14156,0e9387a4aa548998eda8f2bb4a5cd799345d5367@65.21.198.100:11056,ba60b2331fa5f8f63f74108eb5125cd31c07cc43@208.88.251.50:26656,4d95665a3b60b478c4191c090a1e8ba5bfda448c@139.177.182.181:26656,d4ff71a5cca634651ff1580ea368bc8e786d3102@212.102.34.6:36656,5e6678fc29b62e1a81144bf010dfd1228dff1031@63.229.234.75:26656,75be8f90f2aa1da26744ca52ce38aff2d806df26@213.239.207.165:20986,50dd2f2f9ab1c1ddbc2d67bedeed768d7227d331@46.4.107.112:14156,ad7312bcf1beac7ceb817cb6a47b6697d63daf5e@34.134.46.169:26656,9ffc9ade220b0a0fc0c9abded0b032f94f3cbb38@65.108.109.103:10656,281bca9514fe002d90a4b3db2a1567f66a0ec357@159.69.66.149:26646,8b08753a36812d7e6d8799e6afdebe9b05119ba7@95.211.66.25:13656,a53131c0903105f7107b44bcee337ac1e7e4e260@148.113.6.121:34656,8663e7da214559e6dd094a1dde7b877cf9571b6f@37.59.18.132:20002,75c7972462772278129cb6585f63b94fa4ee81a7@44.196.14.252:26656,2c390272eb51d00f2cb312b31471adc8a27b76d4@84.244.95.233:26656,41caa4106f68977e3a5123e56f57934a2d34a1c1@95.214.55.227:27176,4b301813e7db41ed3c8109994e18c0e701b29acd@35.185.11.168:26656,ff826c4b26ddfb7d5d79612585624115712b0a27@80.64.208.52:26656,15861ad4a98ee9a872f85fdace2b23422c796cdf@146.59.81.92:36656,97ede4e6d2300abe224af6d1d86e62a5f9c4d6b8@69.197.6.24:22856,fde339de96b78d7d602dc653d88c8064c85cbf0f@128.140.76.182:26658,47f79f4296e898a40254bcbffc904d6f0a8a55cc@34.86.193.165:26656,b690d37718edd4000c11c80addf5565b72e08078@35.188.241.180:26656,db1748aee7658ce2614e27acb38d184654027194@35.245.10.16:26656,1f3736062f874d214e132084644c6a3f357936de@35.186.165.222:26656,6e42ef37b18bd6547c4ce7a76111d5377c25e684@34.145.250.69:26656,47ae983d3619b09a44afd5ac2d980f035849f49a@95.217.202.49:48656,ed7ef7587d0d7552e973d6fa480ce0ff7f9f39ee@34.70.157.247:26656,9ff90e8c4d8078f00f60ccc9b8cfc3636d8c498c@64.226.112.184:32357,24a9d13e9e95ecf030253b64c1ef6f283f37d155@34.125.175.228:26656,c94fd60124e3656df54ff965d178e36c760c195d@65.108.57.224:26656,d733db63985b400d5e84d11adcced8eb5a5b904c@152.67.3.239:26656,03d22833a8b8e414c2da861cab1cbe8f5694063b@5.78.93.29:26656,362d69cdada0d5926181c602c40f7a7d97186d87@157.90.95.15:26656,3cade0e5deb5bbd8adfa98f62eca48b7d8f87414@34.16.142.53:26656,c5c7a59b0a3e7f4eedba9d76cfcd222607cd0682@141.95.73.27:41656,ad242fe3d7e5abffcae24af91fdcbd004daa3d0f@65.21.84.223:14156,e15ca4ba5ee5dd515ba73c53737bb3f6deb051a1@65.108.121.190:2130,ebc272824924ea1a27ea3183dd0b9ba713494f83@95.214.52.139:27176,be0f0de2c77ad41401711d6a395dd6c6440ef124@144.217.158.217:26656,3fb643a2e4bcd9e79ea745c0e61a06fa3ab99ecf@15.235.87.217:26656,50ce1794942a3f28adef700caf51b49d3e3beacb@34.83.56.101:26656,fdc824a6662e0d1a0265269b13d2bcccb663fb8f@195.20.239.232:3080,733ee0a5a74de20745c5236df015642aa53060bc@57.128.96.155:14156,d87359b6d0f7c736282d48bd38525e2f50cb60d3@35.245.9.232:26656,d366022d8505e328ee4f0b9b5b8fde210013d178@65.108.230.46:14156,590de0c91ce7c6ca1dbdf32c75873617b5503f9f@65.108.101.4:29886,e6c0d56e8c220b835eb5f4a2786cbcd4ea56696a@107.150.119.135:26656,de0754813904e926583afe4fc83c07173316e8b2@186.233.186.53:26656,04def519ec8b578bc6d9efaa56eb957a131cf6f3@135.181.192.166:26656,163d9989fb6ff216db571b2a95ce47fb9282f93f@157.90.26.189:26658,499f8844c9349e38795bf92c166ff71c6dcb7815@51.255.66.46:3080,17555d5d9b45850eecc5fffc05858c4c5ab33303@49.13.59.226:27002,9a3795cef9443b37db9f0e769902729e9865f383@51.15.206.33:26656,c23b220b53aee5b26d1babd62f9c4819c2f47c54@65.109.117.74:26656,402902424e92850107cd0d5b35fcf40b54256b91@87.20.12.119:26656,ee74b3c6d5037e2c56bc95b989b0a26643d79aa4@81.0.220.94:24556,c08eced51e376d9314b313cbece53b1059931ba6@65.108.238.103:14156,38955171f30ec1c2817673c84012a9d723623862@116.202.208.214:2620,373d6987909193e960c0ede48f69afc850648fd9@148.251.53.110:3080,a919153f11e86a0d80672fefea7e845b9bd1bbde@65.19.136.133:23356,0d67a804aae86f5bb60041abb0e7ce15644167ba@95.217.144.107:14156,7bb0b626fe74992f07447e5f9239a188317517a7@45.83.104.218:2000"
	peersString := "3cfe6e50a07cce40881a5648c54a823769e9bf47@135.181.129.122:26656,302cb1aa66e2ce8be9a766b6e40d98747cdcf906@65.109.90.96:60656,f40f1b40032802bc7552f2d88b29523d0da7a948@34.79.14.20:26656,aad751c8a0b3e887fc75ec9e11c85853c93ce8ce@65.108.68.86:11056,9ffc9ade220b0a0fc0c9abded0b032f94f3cbb38@65.108.109.103:10656,c94fd60124e3656df54ff965d178e36c760c195d@65.108.57.224:26656,04def519ec8b578bc6d9efaa56eb957a131cf6f3@135.181.192.166:26656,163d9989fb6ff216db571b2a95ce47fb9282f93f@157.90.26.189:26658,fde339de96b78d7d602dc653d88c8064c85cbf0f@128.140.76.182:26658,4d95665a3b60b478c4191c090a1e8ba5bfda448c@139.177.182.181:26656,5d650f2357b03b932c12c983b14eb83ba6c0771b@34.82.228.43:26656,75be8f90f2aa1da26744ca52ce38aff2d806df26@213.239.207.165:20986,362d69cdada0d5926181c602c40f7a7d97186d87@157.90.95.15:26656,38955171f30ec1c2817673c84012a9d723623862@116.202.208.214:2620,53f2d2169c9febf4f43d281c354fc51a83e0b294@5.9.99.172:14156,d4ff71a5cca634651ff1580ea368bc8e786d3102@212.102.34.6:36656,8b08753a36812d7e6d8799e6afdebe9b05119ba7@95.211.66.25:13656,8663e7da214559e6dd094a1dde7b877cf9571b6f@37.59.18.132:20002,0d67a804aae86f5bb60041abb0e7ce15644167ba@95.217.144.107:14156,86cbeedea81c89192f5a57f2916c4c89e0f2686a@46.166.140.167:26656,41caa4106f68977e3a5123e56f57934a2d34a1c1@95.214.55.227:27176,75c7972462772278129cb6585f63b94fa4ee81a7@44.196.14.252:26656,97ede4e6d2300abe224af6d1d86e62a5f9c4d6b8@69.197.6.24:22856,373d6987909193e960c0ede48f69afc850648fd9@148.251.53.110:3080,6e42ef37b18bd6547c4ce7a76111d5377c25e684@34.145.250.69:26656,1f3736062f874d214e132084644c6a3f357936de@35.186.165.222:26656,db1748aee7658ce2614e27acb38d184654027194@35.245.10.16:26656,b690d37718edd4000c11c80addf5565b72e08078@35.188.241.180:26656,47f79f4296e898a40254bcbffc904d6f0a8a55cc@34.86.193.165:26656,24a9d13e9e95ecf030253b64c1ef6f283f37d155@34.125.175.228:26656,a53131c0903105f7107b44bcee337ac1e7e4e260@148.113.6.121:34656,fcbfdd15efc18f9a368fffd6422b1b33995b8533@103.219.170.31:49230,e15ca4ba5ee5dd515ba73c53737bb3f6deb051a1@65.108.121.190:2130,7bb0b626fe74992f07447e5f9239a188317517a7@45.83.104.218:2000,7e041096585d12b99d9f575979744251409dd0a0@93.115.25.106:36656,a919153f11e86a0d80672fefea7e845b9bd1bbde@65.19.136.133:23356,3fb643a2e4bcd9e79ea745c0e61a06fa3ab99ecf@15.235.87.217:26656,d733db63985b400d5e84d11adcced8eb5a5b904c@152.67.3.239:26656,ad242fe3d7e5abffcae24af91fdcbd004daa3d0f@65.21.84.223:14156,222e3b421288e46d8ae30532aef9458741dcc531@51.91.152.102:20003,ad7312bcf1beac7ceb817cb6a47b6697d63daf5e@34.134.46.169:26656,733ee0a5a74de20745c5236df015642aa53060bc@57.128.96.155:14156,c08eced51e376d9314b313cbece53b1059931ba6@65.108.238.103:14156,de0754813904e926583afe4fc83c07173316e8b2@186.233.186.53:26656,50dd2f2f9ab1c1ddbc2d67bedeed768d7227d331@46.4.107.112:14156,499f8844c9349e38795bf92c166ff71c6dcb7815@51.255.66.46:3080,ff826c4b26ddfb7d5d79612585624115712b0a27@80.64.208.52:26656,d366022d8505e328ee4f0b9b5b8fde210013d178@65.108.230.46:14156,be0f0de2c77ad41401711d6a395dd6c6440ef124@144.217.158.217:26656,f3251d67547f88595966101562094ca538d39f8e@65.108.106.135:26696,081a82d99879d57c0302535ef2dd4e39c83f5a1b@159.69.213.220:27014,d87359b6d0f7c736282d48bd38525e2f50cb60d3@35.245.9.232:26656,15861ad4a98ee9a872f85fdace2b23422c796cdf@146.59.81.92:36656,3cade0e5deb5bbd8adfa98f62eca48b7d8f87414@34.16.142.53:26656,590de0c91ce7c6ca1dbdf32c75873617b5503f9f@65.108.101.4:29886,2c390272eb51d00f2cb312b31471adc8a27b76d4@84.244.95.233:26656,03d22833a8b8e414c2da861cab1cbe8f5694063b@5.78.93.29:26656,ed7ef7587d0d7552e973d6fa480ce0ff7f9f39ee@34.70.157.247:26656,281bca9514fe002d90a4b3db2a1567f66a0ec357@159.69.66.149:26646,c23b220b53aee5b26d1babd62f9c4819c2f47c54@65.109.117.74:26656,e19df6f4008a06e393895b32efe3116c657450f9@65.21.154.144:27014,4b301813e7db41ed3c8109994e18c0e701b29acd@35.185.11.168:26656,42e5448dc33ec46ee91efbcc74d3078ea9e3ddc4@185.165.170.88:26656,e8e3359cfe774d0cad774a0bf34f53cc5b872703@95.216.161.87:26656,0e9387a4aa548998eda8f2bb4a5cd799345d5367@65.21.198.100:11056,fdc824a6662e0d1a0265269b13d2bcccb663fb8f@195.20.239.232:3080,ee74b3c6d5037e2c56bc95b989b0a26643d79aa4@81.0.220.94:24556,ebc272824924ea1a27ea3183dd0b9ba713494f83@95.214.52.139:27176"

	peersList := strings.Split(peersString, ",")
	sw := p2p.NewSwitch(config.P2P, transport)
	sw.SetLogger(tmlogger.With("module", "switch"))
	// #TODO how does the switch connect to peers and start downloading blocks?
	sw.AddReactor("BLOCKSYNC", myreactor)
	sw.SetNodeInfo(nodeInfo)
	sw.SetNodeKey(nodeKey)
	// don't add them as persistent for now, we'll be spamming them
	//during testing with the reconnects.
	//sw.AddPersistentPeers(peersList)
	//This is the PEX reactor, should only be used for peer discovery,
	//am downloading them from the polkachu site so not needed.
	//createPEXReactorAndAddToSwitch(nil, config, sw, logger)
	sw.Start()
	for peerIndex := range peersList {
		peerParts := strings.Split(peersList[peerIndex], "@")
		tcpAddr, err := net.ResolveTCPAddr("tcp", peerParts[1])
		if err != nil {
			tmlogger.Error("failed to resolve tcp address", err.Error())
		}
		peerAddress := p2p.NewNetAddress(p2p.ID(peerParts[0]), tcpAddr)
		//this will add the peer to the address if it connects
		err = sw.DialPeerWithAddress(peerAddress)
		if err != nil {
			tmlogger.Error("failed to dial peer", err.Error())
		} else {
			tmlogger.Info("Halleilujah, connected to peer", "peer", peerAddress)
		}
		//Need SetPeerRange before the peer is added to the pool.
		//SetPeerRange is received from a
		// 	err = sw.DialPeerWithAddress(peerAddress)
		// 	//peer, err := transport.Dial(*peerAddress, ^)
		// 	if err != nil {
		// 		log.Fatal("failed to dial peer", err.Error())
	}
	for {
		time.Sleep(5 * time.Second)
	}
}

func main() {
	var performFastSync = false
	var performIndexDb = true
	var usingExternalHardDrive = true
	tmlogger := tmlog.NewTMLogger(os.Stdout)
	tmlogger = tmlog.NewFilter(tmlogger, tmlog.AllowInfo()) //, log.AllowInfo(), log.AllowError())
	//On Mac
	userDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	var homeFolder = path.Join(userDir, "workspace", "sommelier-data")
	var dataFolder = path.Join(homeFolder, "data")
	if usingExternalHardDrive {
		dataFolder = "/Volumes/harry/sommelier"
	}
	if performFastSync {
		fastSync(tmlogger, homeFolder, dataFolder)
	}
	if performIndexDb {
		indexDbStore(tmlogger, dataFolder)
		//indexDbStore(tmlogger, dataFolder)
	}

}

func indexDbStore(logger tmlog.Logger, dataFolder string) {
	//storedb, err := dbm.NewGoLevelDB("blockstore", "/Users/richard/workspace/somm-bucket/.sommelier/data")
	//storedb, err := dbm.NewGoLevelDB("blockstore", "/Users/richard/workspace/cometbft")
	storedb, err := dbm.NewGoLevelDB("blockstore", dataFolder)
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
	// 2023/12/18-12234205.parquet
	start_block := int64(12234205)
	//end_block := start_block + 200000 //should be about a week
	//end_block := start_block + 30000 //should be about a day
	//end_block := start_block + 1000 //should be about an hour
	end_block := start_block + 5000000
	defer flappy_trade.CloseCurrentParquetFile()
	for i := start_block; i < end_block; i++ {
		block := store.LoadBlock(i)
		if block == nil {
			panic("block is nil")
		} else {
			transactions := flappy_trade.EncodeBlockTransactions(block, 0.0, 0.0, logger)
			flappy_trade.IndexTransactions(transactions, block, logger)
		}
		//pw.Flush(true)
	}
	//pw.WriteStop()
}
