package core

import (
	"flag"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/peterbourgon/ff/v4"
	log "github.com/sirupsen/logrus"

	_ "github.com/lenismtho/area/docs"
	"github.com/lenismtho/area/pkg/core/cache"
	"github.com/lenismtho/area/pkg/core/database"
	"github.com/lenismtho/area/pkg/core/router"
)

// TODO: refacto with ffcli or cobra
// flags
var (
	fs       = flag.NewFlagSet("core", flag.ContinueOnError)
	grpcPort = fs.Uint("grpcPort", uint(9090), "grpc port")
	httpPort = fs.Uint("httpPort", uint(8080), "http port")
)

func init() {
	err := ff.Parse(fs, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

// @title Area API
// @version 1.0
// @description This is the Area API server.
// @host localhost:8080
// @BasePath /v1
func Main() {
	//Load the .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
	}

	if os.Getenv("ENV") == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	cacheDb := cache.NewRedisCache()
	if err != nil {
		log.Fatal(err)
	}
	// Start gRPC srv

	//gmailConn, err := grpc.Dial(os.Getenv("GMAIL_SERVICE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatalf("did not connect: %v", err)
	//}
	//gmailClient := protogen.NewGmailServiceReactionClient(gmailConn)

	// db

	//

	// Setup router
	routes, err := router.NewRouter(&db, &cacheDb)
	if err != nil {
		log.Fatal(err)
	}
	// Start HTTP srv
	if err := routes.Run(":" + strconv.Itoa(int(*httpPort))); err != nil {
		log.Fatal(err)
	}
}
