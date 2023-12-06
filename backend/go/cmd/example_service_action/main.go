package main

import (
	"context"
	"fmt"
	"math/big"
	"net"
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"github.com/lenismtho/area/pkg/protogen"
)

type polls struct {
	id       uint32
	respType protogen.Format_Kind
	token string
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
		return
	}

	lis, err := net.Listen("tcp", os.Getenv($SERVICE_NAME))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	watchers := &[]polls{}
	mut := &sync.Mutex{}

	$SERVICEServer, err := New$SERVICEActionServer(watchers, mut)
	if err != nil {
		log.Fatal(err)
	}
	protogen.Register$SERVICEServiceActionServer(grpcServer, $SERVICEServer)

	coreConn, err := grpc.Dial(os.Getenv("CORE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	coreClient := protogen.NewCoreServiceClient(coreConn)

	go func() {
		err = grpcServer.Serve(lis)
	}()

	log.Info("Starting service")

	func() {
		err := Routine(coreClient)
		if err != nil {
			log.WithField("error", err).Error("failed to start event watcher routine")
		}
	}()
}

type $SERVICEActionServer struct {
watchers *[]polls
mut      *sync.Mutex
}

func New$SERVICEActionServer(watchers *[]polls, mut *sync.Mutex) (*$SERVICEActionServer, error) {
	return &$SERVICEActionServer{
		watchers: watchers,
		mut:      mut,
	}, nil
}

func Routine(coreClient protogen.CoreServiceClient) error {
	for {
		time.Sleep(5 * time.Second)
	}
}

func format(event string, kind protogen.Format_Kind) ([]byte, error) {
	switch kind {
	case protogen.Format_UndefinedKind:
		return nil, fmt.Errorf("undefined format")
	}

	return nil, fmt.Errorf("not implemented")
}
