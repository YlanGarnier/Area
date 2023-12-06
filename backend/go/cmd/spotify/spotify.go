package spotify

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/lenismtho/area/pkg/core/utils"
	"github.com/lenismtho/area/pkg/protogen"
)

type watchersType uint

const (
	artistWatcher watchersType = iota
	songWatcher
)

type polls struct {
	id       uint32
	respType protogen.Format_Kind
	token    string

	kind watchersType
	name string
}

func Main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
		return
	}

	lis, err := net.Listen("tcp", os.Getenv("SPOTIFY_SERVICE"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	watchers := &[]polls{}
	mut := &sync.Mutex{}

	SpotifyServer, err := NewSpotifyActionServer(watchers, mut)
	if err != nil {
		log.Fatal(err)
	}
	protogen.RegisterSpotifyServiceActionServer(grpcServer, SpotifyServer)

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
		err := Routine(coreClient, watchers)
		if err != nil {
			log.WithField("error", err).Error("failed to start event watcher routine")
		}
	}()
}

type SpotifyActionServer struct {
	watchers *[]polls
	mut      *sync.Mutex

	protogen.UnimplementedSpotifyServiceActionServer
}

func NewSpotifyActionServer(watchers *[]polls, mut *sync.Mutex) (*SpotifyActionServer, error) {
	return &SpotifyActionServer{
		watchers: watchers,
		mut:      mut,
	}, nil
}

type GetTrackRes struct {
	Item struct {
		Album struct {
			Name string `json:"name"`
		} `json:"album"`
		Artists []struct {
			Name string `json:"name"`
		} `json:"artists"`
		Name string `json:"name"`
	} `json:"item"`
	IsPlaying bool `json:"is_playing"`
}

func (s SpotifyActionServer) WatchArtist(ctx context.Context, request *protogen.Watcher_Request) (*protogen.Empty, error) {
	s.mut.Lock()
	defer s.mut.Unlock()

	*s.watchers = append(*s.watchers, polls{
		id:       request.Id,
		respType: request.ResponseType,
		token:    request.Token,

		kind: artistWatcher,
		name: request.Name,
	})

	return &protogen.Empty{}, nil
}

func (s SpotifyActionServer) WatchSong(ctx context.Context, request *protogen.Watcher_Request) (*protogen.Empty, error) {
	s.mut.Lock()
	defer s.mut.Unlock()

	*s.watchers = append(*s.watchers, polls{
		id:       request.Id,
		respType: request.ResponseType,
		token:    request.Token,

		kind: songWatcher,
		name: request.Name,
	})

	return &protogen.Empty{}, nil
}

func Routine(coreClient protogen.CoreServiceClient, watchers *[]polls) error {
	for {
		for _, w := range *watchers {
			track, err := utils.GetFromJsonReq[GetTrackRes]("https://api.spotify.com/v1/me/player/currently-playing", utils.GET, "", []utils.Header{
				{
					"Authorization",
					fmt.Sprintf("Bearer %s", w.token),
				},
			}, "")
			if err != nil {
				log.Error(err)
				continue
			}

			if w.kind == artistWatcher {
				log.Println("artists", track.Item.Artists)
				for _, artist := range track.Item.Artists {
					if artist.Name == w.name {
						raw, err := format(track, w.respType)
						_, err = coreClient.ForwardAction(context.Background(), &protogen.ForwardActionReq{
							Id:   w.id,
							Data: raw,
						})
						if err != nil {
							return err
						}
					}
				}
			} else if w.kind == songWatcher {
				log.Println("song", track.Item.Name)
				if track.Item.Name == w.name {
					raw, err := format(track, w.respType)
					_, err = coreClient.ForwardAction(context.Background(), &protogen.ForwardActionReq{
						Id:   w.id,
						Data: raw,
					})
					if err != nil {
						return err
					}
				}
			}
		}
		time.Sleep(5 * time.Second)
	}
}

func format(track GetTrackRes, kind protogen.Format_Kind) ([]byte, error) {
	switch kind {
	case protogen.Format_GHIncidentReportKind:
		p := protogen.Format_GHIncidentReport{
			Title:   fmt.Sprintf("New listening to %s detected !", track.Item.Artists[0].Name),
			Content: fmt.Sprintf("You are currently listening to %s from %s - %s", track.Item.Name, track.Item.Artists[0].Name, track.Item.Album.Name),
		}
		return proto.Marshal(&p)
	case protogen.Format_OnlyTitleKind:
		p := protogen.Format_OnlyTitle{
			Title: fmt.Sprintf("New listening to %s from %s - %s detected !", track.Item.Name, track.Item.Artists[0].Name, track.Item.Album.Name),
		}
		return proto.Marshal(&p)
	case protogen.Format_NoParamKind:
		p := protogen.Format_NoParam{}
		return proto.Marshal(&p)
	case protogen.Format_TagsKind:
		p := protogen.Format_Tags{
			Tags: []string{
				track.Item.Name,
				track.Item.Artists[0].Name,
				track.Item.Album.Name,
			},
		}
		return proto.Marshal(&p)
	}

	return nil, fmt.Errorf("not implemented")
}
