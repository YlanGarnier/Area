package notion

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

type notionServer struct {
	protogen.UnimplementedNotionServiceReactionServer
}

func Main() {
	err := godotenv.Load(".env")
	lis, err := net.Listen("tcp", os.Getenv("NOTION_SERVICE"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	log.Info("listening on", os.Getenv("NOTION_SERVICE"))
	nServer, err := NewNotionServer()
	if err != nil {
		log.Fatal(err)
	}

	protogen.RegisterNotionServiceReactionServer(grpcServer, nServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (n notionServer) CreateDefaultComment(ctx context.Context, request *protogen.Format_OnlyTitle) (*protogen.Empty, error) {
	body := map[string]interface{}{
		"parent": map[string]interface{}{
			"page_id": request.GetBase().GetTarget(),
		},
		"rich_text": []map[string]interface{}{
			{
				"text": map[string]interface{}{
					"content": "This messsage has been created by Area API",
				},
			},
		},
	}

	data, err := json.Marshal(body)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/comments", os.Getenv("NOTION_API_URL")), bytes.NewBuffer(data))

	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("Authorization", "Bearer "+request.GetBase().GetToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("New default comment created !\n")
	} else {
		fmt.Printf("Error during the creation of the comment. Code error : %d\n", resp.StatusCode)
	}

	return &protogen.Empty{}, nil
}

func (n notionServer) CreateComment(ctx context.Context, request *protogen.Format_GHIncidentReport) (*protogen.Empty, error) {
	body := map[string]interface{}{
		"parent": map[string]interface{}{
			"page_id": request.GetBase().GetTarget(),
		},
		"rich_text": []map[string]interface{}{
			{
				"text": map[string]interface{}{
					"content": fmt.Sprintf("%s\n\n%s", request.GetTitle(), request.GetContent()),
				},
			},
		},
	}

	data, err := json.Marshal(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/comments", os.Getenv("NOTION_API_URL")), bytes.NewBuffer(data))

	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("Authorization", "Bearer "+request.GetBase().GetToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("New comment created !\n")
	} else {
		fmt.Printf("Error during the creation of the comment. Code error : %d\n", resp.StatusCode)
	}

	return &protogen.Empty{}, nil
}

func (n notionServer) CreateDefaultPage(ctx context.Context, request *protogen.Format_OnlyTitle) (*protogen.Empty, error) {
	body := map[string]interface{}{
		"parent": map[string]interface{}{
			"database_id": request.GetBase().GetTarget(),
		},
		"properties": map[string]interface{}{
			"Name": map[string]interface{}{
				"title": []map[string]interface{}{
					{
						"type": "text",
						"text": map[string]interface{}{
							"content": "Area API Page",
						},
					},
				},
			},
		},
		"children": []interface{}{
			map[string]interface{}{
				"object": "block",
				"type":   "heading_2",
				"heading_2": map[string]interface{}{
					"rich_text": []map[string]interface{}{
						{
							"type": "text",
							"text": map[string]interface{}{
								"content": "Area API",
							},
						},
					},
				},
			},
			map[string]interface{}{
				"object": "block",
				"type":   "paragraph",
				"paragraph": map[string]interface{}{
					"rich_text": []map[string]interface{}{
						{
							"type": "text",
							"text": map[string]interface{}{
								"content": "This page has been created by Area API",
							},
						},
					},
				},
			},
		},
	}

	data, err := json.Marshal(body)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/pages", os.Getenv("NOTION_API_URL")), bytes.NewBuffer(data))

	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("Authorization", "Bearer "+request.GetBase().GetToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("New page created !\n")
	} else {
		fmt.Printf("Error during the creation of the page. Code error : %d\n", resp.StatusCode)
	}

	return &protogen.Empty{}, nil
}

func (n notionServer) CreatePage(ctx context.Context, request *protogen.Format_GHIncidentReport) (*protogen.Empty, error) {
	body := map[string]interface{}{
		"parent": map[string]interface{}{
			"database_id": request.GetBase().GetTarget(),
		},
		"properties": map[string]interface{}{
			"Name": map[string]interface{}{
				"title": []map[string]interface{}{
					{
						"type": "text",
						"text": map[string]interface{}{
							"content": request.GetTitle(),
						},
					},
				},
			},
		},
		"children": []interface{}{
			map[string]interface{}{
				"object": "block",
				"type":   "heading_2",
				"heading_2": map[string]interface{}{
					"rich_text": []map[string]interface{}{
						{
							"type": "text",
							"text": map[string]interface{}{
								"content": request.GetTitle(),
							},
						},
					},
				},
			},
			map[string]interface{}{
				"object": "block",
				"type":   "paragraph",
				"paragraph": map[string]interface{}{
					"rich_text": []map[string]interface{}{
						{
							"type": "text",
							"text": map[string]interface{}{
								"content": request.GetContent(),
							},
						},
					},
				},
			},
		},
	}

	data, err := json.Marshal(body)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/pages", os.Getenv("NOTION_API_URL")), bytes.NewBuffer(data))

	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("Authorization", "Bearer "+request.GetBase().GetToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("New page created !\n")
	} else {
		fmt.Printf("Error during the creation of the page. Code error : %d\n", resp.StatusCode)
	}

	return &protogen.Empty{}, nil
}

func (n notionServer) CreateBlock(ctx context.Context, request *protogen.Format_GHIncidentReport) (*protogen.Empty, error) {
	body := map[string]interface{}{
		"children": []interface{}{
			map[string]interface{}{
				"object": "block",
				"type":   "heading_2",
				"heading_2": map[string]interface{}{
					"rich_text": []map[string]interface{}{
						{
							"type": "text",
							"text": map[string]interface{}{
								"content": request.GetTitle(),
							},
						},
					},
				},
			},
			map[string]interface{}{
				"object": "block",
				"type":   "paragraph",
				"paragraph": map[string]interface{}{
					"rich_text": []map[string]interface{}{
						{
							"type": "text",
							"text": map[string]interface{}{
								"content": request.GetContent(),
							},
						},
					},
				},
			},
		},
	}
	data, err := json.Marshal(body)

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/blocks/%s/children", os.Getenv("NOTION_API_URL"), request.GetBase().GetTarget()), bytes.NewBuffer(data))

	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Notion-Version", "2022-06-28")
	req.Header.Add("Authorization", "Bearer "+request.GetBase().GetToken())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.WithField("error", err).Error()
		return &protogen.Empty{}, nil
	}

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("New block created !\n")
	} else {
		fmt.Printf("Error during the creation of the block. Code error : %d\n", resp.StatusCode)
	}
	return &protogen.Empty{}, nil
}

func NewNotionServer() (protogen.NotionServiceReactionServer, error) {
	return &notionServer{}, nil
}
