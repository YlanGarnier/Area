package linkedin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/lenismtho/area/pkg/protogen"
)

func Main() {
	err := godotenv.Load(".env")
	lis, err := net.Listen("tcp", os.Getenv("LINKEDIN_SERVICE"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	log.Info("listening on", os.Getenv("LINKEDIN_SERVICE"))
	lkServer, err := NewLinkedinServer()
	if err != nil {
		log.Fatal(err)
	}

	protogen.RegisterLinkedinServiceReactionServer(grpcServer, lkServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

var (
	nbDefaultPost = 0
	nbPost        = 0
)

func getUserInfo(accessToken string) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/userinfo", os.Getenv("LINKEDIN_API_URL")), nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return parseUserInfo(body)
}

func parseUserInfo(response []byte) (string, error) {
	var data map[string]interface{}

	err := json.Unmarshal(response, &data)
	if err != nil {
		return "", err
	}

	sub, ok := data["sub"].(string)
	if !ok {
		return "", fmt.Errorf("no sub key found")
	}

	return sub, nil
}

type linkedinServer struct {
	protogen.UnimplementedLinkedinServiceReactionServer
}

func (lk linkedinServer) CreateDefaultPost(ctx context.Context, request *protogen.Format_OnlyTitle) (*protogen.Empty, error) {
	userId, err := getUserInfo(request.GetBase().GetToken())

	if err != nil {
		return &protogen.Empty{}, nil
	}

	postData := map[string]interface{}{
		"author":         "urn:li:person:" + userId,
		"lifecycleState": "PUBLISHED",
		"specificContent": map[string]interface{}{
			"com.linkedin.ugc.ShareContent": map[string]interface{}{
				"shareCommentary": map[string]interface{}{
					"text": "This message was created by AREA API" + " #" + strconv.Itoa(nbDefaultPost),
				},
				"shareMediaCategory": "NONE",
			},
		},
		"visibility": map[string]interface{}{
			"com.linkedin.ugc.MemberNetworkVisibility": "PUBLIC",
		},
	}

	postDataJSON, err := json.Marshal(postData)
	if err != nil {
		log.Fatalf("error: %v", err)
		return &protogen.Empty{}, nil
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/ugcPosts", os.Getenv("LINKEDIN_API_URL")), bytes.NewBuffer(postDataJSON))
	if err != nil {
		log.Fatalf("error: %v", err)
		return &protogen.Empty{}, nil
	}

	req.Header.Set("Authorization", "Bearer "+request.GetBase().GetToken())
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &protogen.Empty{}, nil
	}

	if resp.StatusCode == http.StatusCreated {
		fmt.Printf("New default post created !\n")
		nbDefaultPost++
	} else {
		fmt.Printf("Error during post creation. Response code : %d\n", resp.StatusCode)
	}

	return &protogen.Empty{}, nil
}

func (lk linkedinServer) CreatePost(ctx context.Context, request *protogen.Format_OnlyTitle) (*protogen.Empty, error) {
	userId, err := getUserInfo(request.GetBase().GetToken())

	if err != nil {
		return &protogen.Empty{}, nil
	}

	postData := map[string]interface{}{
		"author":         "urn:li:person:" + userId,
		"lifecycleState": "PUBLISHED",
		"specificContent": map[string]interface{}{
			"com.linkedin.ugc.ShareContent": map[string]interface{}{
				"shareCommentary": map[string]interface{}{
					"text": request.GetTitle() + " #" + strconv.Itoa(nbPost),
				},
				"shareMediaCategory": "NONE",
			},
		},
		"visibility": map[string]interface{}{
			"com.linkedin.ugc.MemberNetworkVisibility": "PUBLIC",
		},
	}

	postDataJSON, err := json.Marshal(postData)
	if err != nil {
		log.Fatalf("error: %v", err)
		return &protogen.Empty{}, nil
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/ugcPosts", os.Getenv("LINKEDIN_API_URL")), bytes.NewBuffer(postDataJSON))
	if err != nil {
		log.Fatalf("error: %v", err)
		return &protogen.Empty{}, nil
	}

	req.Header.Set("Authorization", "Bearer "+request.GetBase().GetToken())
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &protogen.Empty{}, nil
	}

	if resp.StatusCode == http.StatusCreated {
		fmt.Printf("New post created : %s\n", request.GetTitle())
		nbPost++
	} else {
		fmt.Printf("Error during post creation. Response code : %d\n", resp.StatusCode)
	}

	return &protogen.Empty{}, nil
}

func NewLinkedinServer() (protogen.LinkedinServiceReactionServer, error) {
	return &linkedinServer{}, nil
}
