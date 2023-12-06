package twitter

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/lenismtho/area/pkg/core/utils"
	"github.com/lenismtho/area/pkg/protogen"
)

func Main() {
	err := godotenv.Load(".env")
	lis, err := net.Listen("tcp", os.Getenv("TWITTER_SERVICE"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	log.Info("listening on", os.Getenv("TWITTER_SERVICE"))
	twitterSrv, err := newTwitterServer()
	if err != nil {
		log.Fatal(err)
	}

	protogen.RegisterTwitterServiceReactionServer(grpcServer, twitterSrv)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type twitterServer struct {
	protogen.UnimplementedTwitterServiceReactionServer
}

func craftTweet(text string, poll bool) string {
	return `{
		"text": "` + text + `"
		` + func(poll bool) string {
		if poll {
			return `,"poll": {"options": ["yes", "no"], "duration_minutes": 120}`
		}
		return ""
	}(poll) + `
	}`
}

func (t twitterServer) PostTweet(_ context.Context, req *protogen.Format_NoParam) (*protogen.Empty, error) {
	log.Info("posting tweet")
	_, err := utils.HttpReq("https://api.twitter.com/2/tweets", utils.POST, craftTweet(req.Base.Target, false), []utils.Header{
		{
			"Authorization", "Bearer " + req.Base.Token,
		},
		{
			"Content-Type", "application/json",
		},
	})
	if err != nil {
		return nil, err
	}
	return &protogen.Empty{}, nil
}

func (t twitterServer) PostTweetWithContent(ctx context.Context, req *protogen.Format_OnlyTitle) (*protogen.Empty, error) {
	log.Info("posting tweet with content")
	res, err := utils.HttpReq("https://api.twitter.com/2/tweets", utils.POST, craftTweet(req.Title, false), []utils.Header{
		{
			"Authorization", "Bearer " + req.Base.Token,
		},
		{
			"Content-Type", "application/json",
		},
	})
	fmt.Println(string(res))
	if err != nil {
		return nil, err
	}
	return &protogen.Empty{}, nil
}

func (t twitterServer) PostTweetWithPoll(ctx context.Context, req *protogen.Format_NoParam) (*protogen.Empty, error) {
	log.Info("posting tweet with poll")
	res, err := utils.HttpReq("https://api.twitter.com/2/tweets", utils.POST, craftTweet(req.Base.Target, true), []utils.Header{
		{
			"Authorization", "Bearer " + req.Base.Token,
		},
		{
			"Content-Type", "application/json",
		},
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(string(res))
	return &protogen.Empty{}, nil
}

func (t twitterServer) PostTweetWithContentWithPoll(ctx context.Context, req *protogen.Format_OnlyTitle) (*protogen.Empty, error) {
	log.Info("posting tweet with content and poll")
	res, err := utils.HttpReq("https://api.twitter.com/2/tweets", utils.POST, craftTweet(req.Title, true), []utils.Header{
		{
			"Authorization", "Bearer " + req.Base.Token,
		},
		{
			"Content-Type", "application/json",
		},
	})
	if err != nil {
		return nil, err
	}
	fmt.Println(string(res))
	return &protogen.Empty{}, nil
}

func newTwitterServer() (protogen.TwitterServiceReactionServer, error) {
	return &twitterServer{}, nil
}
