package discord

import (
	"context"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"github.com/lenismtho/area/pkg/core/utils"
	"github.com/lenismtho/area/pkg/protogen"
)

type polls struct {
	id       uint32
	token    string
	respType protogen.Format_Kind

	currentGuilds []Guild
}

type Guild struct {
	Id                       string   `json:"id"`
	Name                     string   `json:"name"`
	Icon                     string   `json:"icon"`
	Owner                    bool     `json:"owner"`
	Permissions              int      `json:"permissions"`
	Features                 []string `json:"features"`
	ApproximateMemberCount   int      `json:"approximate_member_count"`
	ApproximatePresenceCount int      `json:"approximate_presence_count"`
}

func Main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
		return
	}

	mut := &sync.Mutex{}
	arr := &[]polls{}

	lis, err := net.Listen("tcp", os.Getenv("DISCORD_SERVICE"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	discordServer, err := NewDiscordActionServer(mut, arr)
	if err != nil {
		log.Fatal(err)
	}
	protogen.RegisterDiscordServiceActionServer(grpcServer, discordServer)

	coreConn, err := grpc.Dial(os.Getenv("CORE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	coreClient := protogen.NewCoreServiceClient(coreConn)

	go func() {
		err = grpcServer.Serve(lis)
	}()

	log.Info("Starting service")
	for {
		time.Sleep(10 * time.Second)

		mut.Lock()
		for id := range *arr {
			guilds, err := utils.GetFromJsonReq[[]Guild](os.Getenv("DISCORD_API_URL")+"/users/@me/guilds", utils.GET, "", []utils.Header{
				{
					Key:   "Content-Type",
					Value: "application/x-www-form-urlencoded",
				},
				{
					Key:   "Accept",
					Value: "application/json",
				},
				{
					Key:   "Authorization",
					Value: fmt.Sprintf("Bearer %s", (*arr)[id].token),
				},
			}, "")
			if err != nil {
				log.WithField("error", err).Error("failed to get guilds")
				return
			}

			if len(guilds) > len((*arr)[id].currentGuilds) {
				log.WithField("guilds", len(guilds)).Info("new guilds detected")

				newGuilds := utils.Diff(guilds, (*arr)[id].currentGuilds, func(a Guild, b Guild) bool { return a.Id != b.Id })

				raw, err := format(newGuilds, (*arr)[id].respType)
				if err != nil {
					log.WithField("error", err).Error("failed to format")
					continue
				}

				_, err = coreClient.ForwardAction(context.Background(), &protogen.ForwardActionReq{
					Id:   (*arr)[id].id,
					Type: (*arr)[id].respType,
					Data: raw,
				})
				if err != nil {
					log.WithField("error", err).Error("failed to forward")
					continue
				}
				(*arr)[id].currentGuilds = guilds
			}
		}
		mut.Unlock()
	}
}

type DiscordActionServer struct {
	mut *sync.Mutex
	arr *[]polls

	protogen.UnimplementedDiscordServiceActionServer
}

func NewDiscordActionServer(mut *sync.Mutex, arr *[]polls) (protogen.DiscordServiceActionServer, error) {
	return &DiscordActionServer{
		mut: mut,
		arr: arr,
	}, nil
}

func (d *DiscordActionServer) RegisterJoinChannelWatcher(_ context.Context, req *protogen.InviteWatcher_Request) (*protogen.Empty, error) {
	d.mut.Lock()
	defer d.mut.Unlock()

	if utils.Contains(*d.arr, func(e polls) bool { return e.id == req.GetId() }) {
		return nil, fmt.Errorf("already watching this poll")
	}

	guilds, err := utils.GetFromJsonReq[[]Guild](os.Getenv("DISCORD_API_URL")+"/users/@me/guilds", utils.GET, "", []utils.Header{
		{
			Key:   "Content-Type",
			Value: "application/x-www-form-urlencoded",
		},
		{
			Key:   "Accept",
			Value: "application/json",
		},
		{
			Key:   "Authorization",
			Value: fmt.Sprintf("Bearer %s", req.Token),
		},
	}, "")
	if err != nil {
		//fmt.Println("")
		return nil, fmt.Errorf("failed to get guilds: %v", err)
	}

	log.WithField("guilds", len(guilds)).Info("guilds fetched")

	p := polls{
		id:            req.GetId(),
		token:         req.Token,
		respType:      req.ResponseType,
		currentGuilds: guilds,
	}

	*d.arr = append(*d.arr, p)

	return &protogen.Empty{}, nil
}

// format
func format(guild []Guild, kind protogen.Format_Kind) ([]byte, error) {
	switch kind {
	case protogen.Format_GHIncidentReportKind:
		p := protogen.Format_GHIncidentReport{
			Title:   fmt.Sprintf("%d new guilds detected", len(guild)),
			Content: fmt.Sprintf("Guild %s joined", guild[0].Name),
		}
		return proto.Marshal(&p)
	case protogen.Format_OnlyTitleKind:
		p := protogen.Format_OnlyTitle{
			Title: fmt.Sprintf("%d new guilds detected", len(guild)),
		}
		return proto.Marshal(&p)
	case protogen.Format_ManyFilesKind:
		var p []*protogen.Format_ManyFiles_File
		for _, g := range guild {
			p = append(p, &protogen.Format_ManyFiles_File{
				Name: fmt.Sprintf("%s join gist", g.Name),
				Content: fmt.Sprintf("Guild %s joined\nThe is currently %d member connected for a total of %d members.",
					g.Name,
					g.ApproximatePresenceCount,
					g.ApproximateMemberCount,
				),
			})
		}
		return proto.Marshal(&protogen.Format_ManyFiles{
			Files: p,
		})
	}

	return nil, fmt.Errorf("not implemented")
}
