PROTO_PATH=./proto
PROTO_GEN_PATH=./go/pkg/

go_generate:
	protoc -I=$(PROTO_PATH) --go_out=$(PROTO_GEN_PATH)  --go-grpc_out=$(PROTO_GEN_PATH) ./proto/http_action_service.proto
	protoc -I=$(PROTO_PATH) --go_out=$(PROTO_GEN_PATH)  --go-grpc_out=$(PROTO_GEN_PATH) ./proto/pmodel.proto
	protoc -I=$(PROTO_PATH) --go_out=$(PROTO_GEN_PATH)  --go-grpc_out=$(PROTO_GEN_PATH) ./proto/core.proto
	protoc -I=$(PROTO_PATH) --go_out=$(PROTO_GEN_PATH)  --go-grpc_out=$(PROTO_GEN_PATH) ./proto/github_reaction_service.proto
	protoc -I=$(PROTO_PATH) --go_out=$(PROTO_GEN_PATH)  --go-grpc_out=$(PROTO_GEN_PATH) ./proto/discord_action_service.proto
	protoc -I=$(PROTO_PATH) --go_out=$(PROTO_GEN_PATH)  --go-grpc_out=$(PROTO_GEN_PATH) ./proto/ethereum_action_service.proto
	protoc -I=$(PROTO_PATH) --go_out=$(PROTO_GEN_PATH)  --go-grpc_out=$(PROTO_GEN_PATH) ./proto/spotify_action_service.proto
	protoc -I=$(PROTO_PATH) --go_out=$(PROTO_GEN_PATH)  --go-grpc_out=$(PROTO_GEN_PATH) ./proto/twitter_reaction_service.proto
	protoc -I=$(PROTO_PATH) --go_out=$(PROTO_GEN_PATH)  --go-grpc_out=$(PROTO_GEN_PATH) ./proto/twitch_reaction_service.proto
	protoc -I=$(PROTO_PATH) --go_out=$(PROTO_GEN_PATH)  --go-grpc_out=$(PROTO_GEN_PATH) ./proto/dropbox_reaction_service.proto
	protoc -I=$(PROTO_PATH) --go_out=$(PROTO_GEN_PATH)  --go-grpc_out=$(PROTO_GEN_PATH) ./proto/gmail_action_service.proto
	protoc -I=$(PROTO_PATH) --go_out=$(PROTO_GEN_PATH)  --go-grpc_out=$(PROTO_GEN_PATH) ./proto/miro_reaction_service.proto
	protoc -I=$(PROTO_PATH) --go_out=$(PROTO_GEN_PATH)  --go-grpc_out=$(PROTO_GEN_PATH) ./proto/linkedin_reaction_service.proto
	protoc -I=$(PROTO_PATH) --go_out=$(PROTO_GEN_PATH)  --go-grpc_out=$(PROTO_GEN_PATH) ./proto/notion_reaction_service.proto

