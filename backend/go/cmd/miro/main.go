package miro

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/lenismtho/area/pkg/protogen"
)

type miroServer struct {
	protogen.UnimplementedMiroServiceReactionServer
}

func Main() {
	err := godotenv.Load(".env")
	lis, err := net.Listen("tcp", os.Getenv("MIRO_SERVICE"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	log.Info("listening on", os.Getenv("MIRO_SERVICE"))
	miroServ, err := NewMiroServer()
	if err != nil {
		log.Fatal(err)
	}

	protogen.RegisterMiroServiceReactionServer(grpcServer, miroServ)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func NewMiroServer() (protogen.MiroServiceReactionServer, error) {
	return &miroServer{}, nil
}

func (m miroServer) CreateBoard(ctx context.Context, request *protogen.Format_OnlyTitle) (*protogen.Empty, error) {
	data := []byte(fmt.Sprintf(`{
		"description": "Board created by AREA API",
		"name": "%s"
	}`, request.GetTitle()))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", os.Getenv("MIRO_API_V2_URL"), "/boards"), bytes.NewBuffer(data))

	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+request.GetBase().GetToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	if resp.StatusCode == http.StatusCreated {
		fmt.Printf("New board created : '%s'!\n", request.GetTitle())
	} else {
		fmt.Println(resp)
		fmt.Printf("Error during the creation of the board. Code error : %d\n", resp.StatusCode)
	}
	return &protogen.Empty{}, nil
}

var cmpTags = 0

func (m miroServer) CreateTags(ctx context.Context, request *protogen.Format_OnlyTitle) (*protogen.Empty, error) {
	data := []byte(fmt.Sprintf(`{
		"title": "%s #%d",
		"fillColor": "red"
	}`, request.GetTitle(), cmpTags))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/boards/%s/tags", os.Getenv("MIRO_API_V2_URL"), request.GetBase().GetTarget()), bytes.NewBuffer(data))

	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+request.GetBase().GetToken())

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	if resp.StatusCode == http.StatusCreated {
		fmt.Printf("New tags created : '%s'!\n", request.GetTitle())
		cmpTags++
	} else {
		fmt.Println(resp)
		fmt.Printf("Error during the creation of the tags. Code error : %d\n", resp.StatusCode)
	}
	return &protogen.Empty{}, nil
}

func (m miroServer) CreateStickyCard(ctx context.Context, request *protogen.Format_OnlyTitle) (*protogen.Empty, error) {
	data := []byte(fmt.Sprintf(`{
		"data": {
			"content": "%s",
			"shape": "square"
		}
	}`, request.GetTitle()))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/boards/%s/sticky_notes", os.Getenv("MIRO_API_V2_URL"), request.GetBase().GetTarget()), bytes.NewBuffer(data))

	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+request.GetBase().GetToken())

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	if resp.StatusCode == http.StatusCreated {
		fmt.Printf("New Sticky card created : '%s'!\n", request.GetTitle())
	} else {
		fmt.Printf("Error during the creation of the sticky card. Code error : %d\n", resp.StatusCode)
	}
	return &protogen.Empty{}, nil
}

func (m miroServer) CreateCardItem(ctx context.Context, request *protogen.Format_GHIncidentReport) (*protogen.Empty, error) {

	data := []byte(fmt.Sprintf(`{
		"data": {
			"description": "%s",
			"title": "%s"
		}
	}`, request.GetContent(), request.GetTitle()))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/boards/%s/cards", os.Getenv("MIRO_API_V2_URL"), request.GetBase().GetTarget()), bytes.NewBuffer(data))

	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+request.GetBase().GetToken())

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	if resp.StatusCode == http.StatusCreated {
		fmt.Printf("New Card Item created : '%s'!\n", request.GetBase().GetTarget())
	} else {
		fmt.Printf("Error during the creation of the card item. Code error : %d\n", resp.StatusCode)
	}
	return &protogen.Empty{}, nil
}
