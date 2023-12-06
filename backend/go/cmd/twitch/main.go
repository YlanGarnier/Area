package twitch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/lenismtho/area/pkg/protogen"
)

type twitchServer struct {
	protogen.UnimplementedTwitchServiceReactionServer
}

func Main() {
	err := godotenv.Load(".env")
	lis, err := net.Listen("tcp", os.Getenv("TWITCH_SERVICE"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	log.Info("listening on", os.Getenv("TWITCH_SERVICE"))
	twServer, err := NewTwitchServer()
	if err != nil {
		log.Fatal(err)
	}

	protogen.RegisterTwitchServiceReactionServer(grpcServer, twServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type TwitchUser struct {
	ID string `json:"id"`
}

func getUserID(token string) (string, error) {
	req, err := http.NewRequest("GET", "https://api.twitch.tv/helix/users", nil)
	if err != nil {
		return "", fmt.Errorf("error during query creation : %s", err)
	}

	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Client-ID", os.Getenv("TWITCH_CLIENT_ID"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request : %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("the query returned a status code : %d", resp.StatusCode)
	}

	var twitchResponse map[string][]TwitchUser
	err = json.NewDecoder(resp.Body).Decode(&twitchResponse)
	if err != nil {
		return "", fmt.Errorf("error reading JSON response : %s", err)
	}

	if len(twitchResponse["data"]) > 0 {
		user := twitchResponse["data"][0]
		return user.ID, nil
	}

	return "", fmt.Errorf("no user data found in the response")
}

func (t twitchServer) SendDefaultMessage(ctx context.Context, request *protogen.Format_OnlyTitle) (*protogen.Empty, error) {
	data := []byte(`{
		"message": "This message has been sent by Area API",
		"color": "blue"
	}`)

	userId, err := getUserID(request.GetBase().GetToken())

	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/helix/chat/announcements?broadcaster_id=%s&moderator_id=%s", os.Getenv("TWITCH_API_URL"), userId, userId), bytes.NewBuffer(data))

	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Client-ID", os.Getenv("TWITCH_CLIENT_ID"))
	req.Header.Add("Authorization", "Bearer "+request.GetBase().GetToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	if resp.StatusCode == http.StatusNoContent {
		fmt.Printf("New message created !\n")
	} else {
		fmt.Printf("Error during the creation of the message. Code error : %d\n", resp.StatusCode)
	}

	return &protogen.Empty{}, nil
}

func (t twitchServer) SendMessage(ctx context.Context, request *protogen.Format_OnlyTitle) (*protogen.Empty, error) {
	messageData := map[string]string{
		"message": request.GetTitle(),
		"color":   "blue",
	}

	jsonData, err := json.Marshal(messageData)
	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, err
	}

	userId, err := getUserID(request.GetBase().GetToken())

	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/helix/chat/announcements?broadcaster_id=%s&moderator_id=%s", os.Getenv("TWITCH_API_URL"), userId, userId), bytes.NewBuffer(jsonData))

	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Client-ID", os.Getenv("TWITCH_CLIENT_ID"))
	req.Header.Add("Authorization", "Bearer "+request.GetBase().GetToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	if resp.StatusCode == http.StatusNoContent {
		fmt.Printf("New message created : %s\n", request.GetTitle())
	} else {
		fmt.Printf("Error during the creation of the message. Code error : %d\n", resp.StatusCode)
	}

	return &protogen.Empty{}, nil
}

func NewTwitchServer() (protogen.TwitchServiceReactionServer, error) {
	return &twitchServer{}, nil
}
