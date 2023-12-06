package dropbox

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/lenismtho/area/pkg/protogen"
)

type dropboxServer struct {
	protogen.UnimplementedDropboxServiceReactionServer
}

func Main() {
	err := godotenv.Load(".env")
	lis, err := net.Listen("tcp", os.Getenv("DROPBOX_SERVICE"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	log.Info("listening on", os.Getenv("DROPBOX_SERVICE"))
	dbServer, err := NewDropboxServer()
	if err != nil {
		log.Fatal(err)
	}

	protogen.RegisterDropboxServiceReactionServer(grpcServer, dbServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

var (
	nbFolder = 0
)

func (db dropboxServer) CreateFolder(ctx context.Context, request *protogen.Format_OnlyTitle) (*protogen.Empty, error) {

	data := []byte(fmt.Sprintf(`{
		"autorename": false,
		"path": "/%s #%d"
	}`, request.GetTitle(), nbFolder))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/files/create_folder_v2", os.Getenv("DROPBOX_API_URL")), bytes.NewBuffer(data))

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

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("New folder created : %s\n", request.GetTitle())
		nbFolder++
	} else {
		fmt.Printf("Error during the creation of the folder. Code error : %d\n", resp.StatusCode)
	}
	return &protogen.Empty{}, nil
}

func (db dropboxServer) CreateFile(ctx context.Context, request *protogen.Format_GHIncidentReport) (*protogen.Empty, error) {
	fileContents := request.GetContent()

	path := fmt.Sprintf("/%s.paper", request.GetTitle())

	requestBody := bytes.NewBuffer([]byte(fileContents))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/files/paper/create", os.Getenv("DROPBOX_API_URL")), requestBody)
	if err != nil {
		fmt.Println("Erreur lors de la création de la requête:", err)
		return &protogen.Empty{}, nil
	}

	req.Header.Set("Authorization", "Bearer "+request.GetBase().GetToken())
	req.Header.Set("Dropbox-API-Arg", `{"import_format":"html","path":"`+path+`"}`)
	req.Header.Set("Content-Type", "application/octet-stream")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Erreur lors de la requête:", err)
		return &protogen.Empty{}, nil
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("New file created : %s\n", request.GetTitle())
	} else {
		fmt.Printf("Error during the creation of the file. Code error : %d\n", resp.StatusCode)
	}

	return &protogen.Empty{}, nil
}

var cmpTag = 0

func (db dropboxServer) CreateTag(ctx context.Context, request *protogen.Format_OnlyTitle) (*protogen.Empty, error) {
	data := []byte(fmt.Sprintf(`{
		"path": "/%s",
		"tag_text": "%s_%d"
	}`, request.GetBase().GetTarget(), strings.ReplaceAll(strings.ReplaceAll(request.GetTitle(), " ", "_"), ":", "_"), cmpTag))

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/files/tags/add", os.Getenv("DROPBOX_API_URL")), bytes.NewBuffer(data))

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

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("New Tag created : %s\n", request.GetTitle())
		cmpTag++
	} else {
		fmt.Printf("Error during the creation of the tags. Code error : %d\n", resp.StatusCode)
	}
	return &protogen.Empty{}, nil
}

func NewDropboxServer() (protogen.DropboxServiceReactionServer, error) {
	return &dropboxServer{}, nil
}
