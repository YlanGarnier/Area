package github

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/google/go-github/v55/github"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/lenismtho/area/pkg/protogen"
)

func Main() {
	err := godotenv.Load(".env")
	lis, err := net.Listen("tcp", os.Getenv("GITHUB_SERVICE"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	log.Info("listening on", os.Getenv("GITHUB_SERVICE"))
	ghServer, err := NewGithubServer()
	if err != nil {
		log.Fatal(err)
	}

	protogen.RegisterGHServiceReactionServer(grpcServer, ghServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type githubServer struct {
	protogen.UnimplementedGHServiceReactionServer
}

func (g githubServer) CreateRepo(ctx context.Context, req *protogen.Format_OnlyTitle) (*protogen.Empty, error) {
	client := github.NewClient(nil).WithAuthToken(req.GetBase().GetToken())

	_, res, err := client.Repositories.Create(ctx, "", &github.Repository{
		Name: &req.Title,
	})
	if err != nil {
		log.WithField("error", err).Error()
		return nil, err
	}
	log.Println(res.Status)

	return &protogen.Empty{}, nil
}

func (g githubServer) CreateRepoNamed(ctx context.Context, req *protogen.Format_NoParam) (*protogen.Empty, error) {
	client := github.NewClient(nil).WithAuthToken(req.GetBase().GetToken())

	name := req.GetBase().GetTarget()
	_, res, err := client.Repositories.Create(ctx, "", &github.Repository{
		Name: &name,
	})
	if err != nil {
		log.WithField("error", err).Error()
		return nil, err
	}
	log.Println(res.Status)

	return &protogen.Empty{}, nil
}

func (g githubServer) CreateIssue(ctx context.Context, report *protogen.Format_GHIncidentReport) (*protogen.Empty, error) {
	log.Println("received", report)

	path := strings.Split(report.Base.Target, "/")

	client := github.NewClient(nil).WithAuthToken(report.GetBase().GetToken())
	_, res, err := client.Issues.Create(ctx, path[0], path[1], &github.IssueRequest{
		Title: &report.Title,
		Body:  &report.Content,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	log.Println(res.Status)

	return &protogen.Empty{}, nil
}

func (g githubServer) CreateGists(ctx context.Context, req *protogen.Format_ManyFiles) (*protogen.Empty, error) {
	client := github.NewClient(nil).WithAuthToken(req.GetBase().GetToken())

	files := map[github.GistFilename]github.GistFile{}
	for _, f := range req.Files {
		c := f.Content
		files[github.GistFilename(f.Name)] = github.GistFile{
			Content: &c,
		}
	}

	_, res, err := client.Gists.Create(ctx, &github.Gist{
		Files: files,
	})
	if err != nil {
		log.WithField("error", err).Error()
		return nil, err
	}
	log.Println(res.Status)

	return &protogen.Empty{}, nil
}

func (g githubServer) AddTags(ctx context.Context, req *protogen.Format_Tags) (*protogen.Empty, error) {
	client := github.NewClient(nil).WithAuthToken(req.GetBase().GetToken())
	path := strings.Split(req.GetBase().GetTarget(), "/")

	topics, res, err := client.Repositories.ListAllTopics(ctx, path[0], path[1])
	if err != nil {
		log.WithField("error", err).Error()
		return nil, err
	}
	log.Info(res.Status)

	_, res, err = client.Repositories.ReplaceAllTopics(ctx, path[0], path[1], append(topics, req.GetTags()...))
	if err != nil {
		log.WithField("error", err).Error()
		return nil, err
	}
	log.Info(res.Status)

	return &protogen.Empty{}, nil
}

func getLastHashCommit(token string, owner string, nameRepo string) string {
	url := fmt.Sprintf("https://api.github.com/repos/%s/commits/main", owner+"/"+nameRepo)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.WithField("error", err).Error()
		return ""
	}

	req.SetBasicAuth(owner, token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithField("error", err).Error()
		return ""
	}

	defer resp.Body.Close()

	// Lire et décoder la réponse JSON
	var commit map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&commit)
	if err != nil {
		log.WithField("error", err).Error()
		return ""
	}

	sha := commit["sha"].(string)
	fmt.Println("Hash du dernier commit:", sha)

	return sha
}

func (g githubServer) CreateBranch(ctx context.Context, req *protogen.Format_GHNewBranch) (*protogen.Empty, error) {
	path := strings.Split(req.GetBase().GetTarget(), "/")

	// Get last commit hash to create our branch
	lastHash := getLastHashCommit(req.GetBase().GetToken(), path[0], path[1])

	data := map[string]interface{}{
		"ref": fmt.Sprintf("refs/heads/%s", req.GetName()),
		"sha": lastHash,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.WithField("error", err).Error()

		return &protogen.Empty{}, nil
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/git/refs", path[0], path[1])

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+req.GetBase().GetToken())

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	if resp.StatusCode == http.StatusCreated {
		fmt.Printf("Branche '%s' créée avec succès!\n", req.GetName())
	} else {
		fmt.Printf("Erreur lors de la création de la branche. Code de réponse : %d\n", resp.StatusCode)
	}

	return &protogen.Empty{}, nil
}

func NewGithubServer() (protogen.GHServiceReactionServer, error) {
	return &githubServer{}, nil
}
