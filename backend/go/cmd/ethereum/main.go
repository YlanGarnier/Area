package ethereum

import (
	"context"
	"fmt"
	"math/big"
	"net"
	"os"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"github.com/lenismtho/area/pkg/protogen"
)

type pollsTransactionsWatcher struct {
	id       uint32
	respType protogen.Format_Kind
	address  common.Address
}

type pollsEventsWatcher struct {
	id       uint32
	respType protogen.Format_Kind
	event    string
}

func Main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
		return
	}

	eventsWatcherA := eventWatcherAction{
		watchersEvents: &[]pollsEventsWatcher{},
		mut:            &sync.Mutex{},
	}
	transactionsWatcherA := transactionsWatcherAction{
		watchersTransactions: &[]pollsTransactionsWatcher{},
		mut:                  &sync.Mutex{},
	}
	lis, err := net.Listen("tcp", os.Getenv("ETHEREUM_SERVICE"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	ethereumServer, err := NewEthereumActionServer(&transactionsWatcherA, &eventsWatcherA)
	if err != nil {
		log.Fatal(err)
	}
	protogen.RegisterEthereumServiceActionServer(grpcServer, ethereumServer)

	coreConn, err := grpc.Dial(os.Getenv("CORE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	coreClient := protogen.NewCoreServiceClient(coreConn)

	go func() {
		err = grpcServer.Serve(lis)
	}()

	log.Info("Starting service")
	client, err := ethclient.Dial(os.Getenv("INFURA_SECRET_API"))
	go func() {
		err := transactionsWatcherRoutine(&transactionsWatcherA, client, coreClient)
		if err != nil {
			log.WithField("error", err).Error("failed to start transaction watcher routine")
		}
	}()
	func() {
		err := eventsWatcherRoutine(eventsWatcherA, client, coreClient)
		if err != nil {
			log.WithField("error", err).Error("failed to start event watcher routine")
		}
	}()
}

func transactionsWatcherRoutine(transactionsWatcher *transactionsWatcherAction, client *ethclient.Client, coreClient protogen.CoreServiceClient) error {
	ctx := context.Background()
	current, err := client.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("failed to get block number: %v", err)
	}
	for {
		time.Sleep(10 * time.Second)
		transactionsWatcher.mut.Lock()
		for _, poll := range *transactionsWatcher.watchersTransactions {
			new, err := client.BlockNumber(ctx)
			if err != nil {
				log.WithField("error", err).Error("failed to get block number")
				continue
			}
			if new == current {
				continue
			}
			log.Info("new block")
			for i := current + 1; i <= new; i++ {
				var block *types.Block
				block, err := client.BlockByNumber(ctx, big.NewInt(int64(i)))
				if err != nil {
					log.WithField("error", err).Error("failed to get block")
					continue
				}
				log.WithField("i", i).Info("new block")
				for _, tx := range block.Transactions() {
					from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
					if err != nil {
						log.WithField("error", err).Error("failed to get sender")
						continue
					}
					if from == poll.address || (tx.To() != nil && *tx.To() == poll.address) {
						log.Info("new transaction")
						raw, err := formatTransactionsWatcherAction(*tx, poll.address, poll.respType)
						if err != nil {
							log.WithField("error", err).Error("failed to formatTransactionsWatcherAction")
							continue
						}
						_, err = coreClient.ForwardAction(context.Background(), &protogen.ForwardActionReq{
							Id:   poll.id,
							Type: poll.respType,
							Data: raw,
						})
						if err != nil {
							log.WithField("error", err).Error("failed to forward")
							continue
						}
					}
				}
			}
			current = new
		}
		transactionsWatcher.mut.Unlock()
	}
}

func eventsWatcherRoutine(eventsWatcher eventWatcherAction, client *ethclient.Client, coreClient protogen.CoreServiceClient) error {
	ctx := context.Background()
	current, err := client.BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("failed to get block number: %v", err)
	}
	for {
		time.Sleep(10 * time.Second)
		eventsWatcher.mut.Lock()
		for _, poll := range *eventsWatcher.watchersEvents {
			fmt.Println("new event")
			new, err := client.BlockNumber(ctx)
			fmt.Println("Starting from block", current)
			fmt.Println("Ending to block", new)
			if err != nil {
				fmt.Println("Failed to retrieve block number:", err)
				continue
			}
			if new == current {
				continue
			}
			eventSignatureBytes := []byte(poll.event)
			eventSignaturehash := crypto.Keccak256Hash(eventSignatureBytes)
			q := ethereum.FilterQuery{
				FromBlock: big.NewInt(int64(current + 1)),
				ToBlock:   big.NewInt(int64(new)),
				Topics: [][]common.Hash{
					{eventSignaturehash},
				},
			}
			logs, err := client.FilterLogs(context.Background(), q)
			if err != nil {
				log.Fatal(err)
			}
			if len(logs) > 0 {
				fmt.Println("Event detected")
				raw, err := formatEventsWatcherAction(logs, poll.event, poll.respType)
				if err != nil {
					log.WithField("error", err).Error("failed to formatTransactionsWatcherAction")
					continue
				}
				_, err = coreClient.ForwardAction(context.Background(), &protogen.ForwardActionReq{
					Id:   poll.id,
					Type: poll.respType,
					Data: raw,
				})
				if err != nil {
					log.WithField("error", err).Error("failed to forward")
					continue
				}
			}
			current = new
		}
		eventsWatcher.mut.Unlock()
	}
}

type eventWatcherAction struct {
	watchersEvents *[]pollsEventsWatcher
	mut            *sync.Mutex
}

type transactionsWatcherAction struct {
	watchersTransactions *[]pollsTransactionsWatcher
	mut                  *sync.Mutex
}

type EthereumActionServer struct {
	transactionsWatchers *transactionsWatcherAction
	eventsWatchers       *eventWatcherAction
	protogen.UnimplementedEthereumServiceActionServer
}

func (e EthereumActionServer) RegisterAddresseWatcher(ctx context.Context, request *protogen.AddressWatcher_Request) (*protogen.Empty, error) {
	*e.transactionsWatchers.watchersTransactions = append(*e.transactionsWatchers.watchersTransactions, pollsTransactionsWatcher{
		id:       request.GetId(),
		respType: request.GetResponseType(),
		address:  common.HexToAddress(request.GetAddress()),
	})
	return &protogen.Empty{}, nil
}

func (e EthereumActionServer) RegisterEventWatcher(ctx context.Context, request *protogen.EventWatcher_Request) (*protogen.Empty, error) {
	*e.eventsWatchers.watchersEvents = append(*e.eventsWatchers.watchersEvents, pollsEventsWatcher{
		id:       request.GetId(),
		respType: request.GetResponseType(),
		event:    request.GetEvent(),
	})
	return &protogen.Empty{}, nil
}
func NewEthereumActionServer(transactionsWatchers *transactionsWatcherAction, eventsWatchers *eventWatcherAction) (protogen.EthereumServiceActionServer, error) {
	return &EthereumActionServer{
		transactionsWatchers: transactionsWatchers,
		eventsWatchers:       eventsWatchers,
	}, nil
}

// formatTransactionsWatcherAction
func formatTransactionsWatcherAction(tx types.Transaction, target common.Address, kind protogen.Format_Kind) ([]byte, error) {
	switch kind {
	case protogen.Format_GHIncidentReportKind:
		p := protogen.Format_GHIncidentReport{
			Title:   fmt.Sprintf("new transactions detected for %s", target.String()),
			Content: fmt.Sprintf("%x", tx),
		}
		return proto.Marshal(&p)
	}
	return nil, fmt.Errorf("not implemented")
}

func formatEventsWatcherAction(logs []types.Log, event string, kind protogen.Format_Kind) ([]byte, error) {
	switch kind {
	case protogen.Format_GHIncidentReportKind:
		p := protogen.Format_GHIncidentReport{
			Title:   fmt.Sprintf("new event detected: %s", event),
			Content: fmt.Sprintf("%x", logs[0].TxHash),
		}
		return proto.Marshal(&p)
	case protogen.Format_OnlyTitleKind:
		p := protogen.Format_OnlyTitle{
			Title: fmt.Sprintf("new event detected: %s", event),
		}
		return proto.Marshal(&p)
	}

	return nil, fmt.Errorf("not implemented")
}
