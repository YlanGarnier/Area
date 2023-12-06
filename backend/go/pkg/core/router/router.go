package router

import (
	"fmt"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/lenismtho/area/docs"
	"github.com/lenismtho/area/pkg/core/cache"
	"github.com/lenismtho/area/pkg/core/controllers/area"
	"github.com/lenismtho/area/pkg/core/controllers/user"
	"github.com/lenismtho/area/pkg/core/database"
	"github.com/lenismtho/area/pkg/core/middleware"
	"github.com/lenismtho/area/pkg/core/srv"
	"github.com/lenismtho/area/pkg/protogen"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		urls := []string{
			"http://localhost:5173",
			"http://localhost:8081",
			"http://localhost",
		}

		origin := c.Request.Header.Get("Origin")
		for _, url := range urls {
			if url == origin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", url)
				c.Writer.Header().Set("Access-Control-Max-Age", "86400")
				c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
				c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
				c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
				c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
				break
			}
		}

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func NewRouter(db *database.DB, cacheDb *cache.Cache) (*gin.Engine, error) {
	r := gin.Default()
	r.Use(CORSMiddleware())

	lis, err := net.Listen("tcp", os.Getenv("CORE"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	httpConn, err := grpc.Dial(os.Getenv("HTTP_SERVICE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	httpClient := protogen.NewHTTPServiceActionClient(httpConn)

	ghConn, err := grpc.Dial(os.Getenv("GITHUB_SERVICE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ghClient := protogen.NewGHServiceReactionClient(ghConn)

	discordConn, err := grpc.Dial(os.Getenv("DISCORD_SERVICE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	discordClient := protogen.NewDiscordServiceActionClient(discordConn)

	ethereumConn, err := grpc.Dial(os.Getenv("ETHEREUM_SERVICE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ethereumClient := protogen.NewEthereumServiceActionClient(ethereumConn)

	twitterConn, err := grpc.Dial(os.Getenv("TWITTER_SERVICE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	twitterClient := protogen.NewTwitterServiceReactionClient(twitterConn)

	twitchConn, err := grpc.Dial(os.Getenv("TWITCH_SERVICE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	twitchClient := protogen.NewTwitchServiceReactionClient(twitchConn)

	dropboxConn, err := grpc.Dial(os.Getenv("DROPBOX_SERVICE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	dropboxClient := protogen.NewDropboxServiceReactionClient(dropboxConn)

	spotifyConn, err := grpc.Dial(os.Getenv("SPOTIFY_SERVICE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	spotifyClient := protogen.NewSpotifyServiceActionClient(spotifyConn)
	gmailConn, err := grpc.Dial(os.Getenv("GMAIL_ACTION_SERVICE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	gmailClient := protogen.NewGmailServiceActionClient(gmailConn)

	miroConn, err := grpc.Dial(os.Getenv("MIRO_SERVICE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	miroClient := protogen.NewMiroServiceReactionClient(miroConn)

	linkedinConn, err := grpc.Dial(os.Getenv("LINKEDIN_SERVICE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	linkedinClient := protogen.NewLinkedinServiceReactionClient(linkedinConn)

	notionConn, err := grpc.Dial(os.Getenv("NOTION_SERVICE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	notionClient := protogen.NewNotionServiceReactionClient(notionConn)

	coreSrv, err := srv.NewCoreServer(*db, ghClient, twitchClient, dropboxClient, twitterClient, miroClient, linkedinClient, notionClient)
	if err != nil {
		log.Fatal(err)
	}
	protogen.RegisterCoreServiceServer(grpcServer, coreSrv)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	r.GET("/about.json", docs.About)

	r.GET("/area/list", func(ctx *gin.Context) {
		ctx.File("area-list.json")
	})

	v1 := r.Group("/v1")
	{

		/*** START SWAGGER ***/
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
		/**** END SWAGGER ****/
		/*** START MIDDLEWARE ***/
		middlewareCtrl, err := middleware.NewMiddleware(*db)
		if err != nil {
			return nil, fmt.Errorf("failed to create middleware controller: %v", err)
		}
		/*** END MIDDLEWARE ***/
		/*** START USER ***/
		userCtrl, err := user.NewController(*db, *cacheDb)
		if err != nil {
			return nil, fmt.Errorf("failed to create user controller: %v", err)
		}
		/*** PUBLIC USERS ***/
		v1.POST("/users/login", userCtrl.Login)
		v1.POST("/users/signup", userCtrl.SignUp)
		v1.POST("/users/signup/oauth", userCtrl.SignUpOauth)
		v1.PUT("/users/reset_password", userCtrl.ResetUserPassword)
		/*** PROTECTED USERS ***/
		v1.GET("/users/me", middlewareCtrl.IsValidToken, userCtrl.Me)
		v1.PUT("/users/me", middlewareCtrl.IsValidToken, userCtrl.UpdateMe)
		v1.PUT("/users/me/email", middlewareCtrl.IsValidToken, userCtrl.UpdateMeEmail)
		v1.PUT("/users/me/password", middlewareCtrl.IsValidToken, userCtrl.UpdateMePassword)
		v1.DELETE("/users/me", middlewareCtrl.IsValidToken, userCtrl.DeleteMe)
		v1.POST("/users/me/services", middlewareCtrl.IsValidToken, userCtrl.UserConnectService)
		v1.GET("/users/me/services", middlewareCtrl.IsValidToken, userCtrl.GetUserServices)
		v1.GET("/users/me/areas", middlewareCtrl.IsValidToken, userCtrl.GetUserAreas)
		v1.DELETE("/users/me/services/:serviceName", middlewareCtrl.IsValidToken, userCtrl.UserDisconnectService)
		v1.GET("/mobile/callback", userCtrl.MobileCallback)
		/*** END USERS ***/

		/*** START SERVICES ***/
		//serviceCtrl, err := services.NewService(*db)
		//if err != nil {
		//	return nil, fmt.Errorf("failed to create service controller: %v", err)
		//}
		/*** PUBLIC SERVICES ***/
		/*** PROTECTED SERVICES ***/
		/*** END SERVICES ***/

		areaCtrl, err := area.NewController(*db, httpClient, discordClient, ethereumClient, spotifyClient, gmailClient)
		if err != nil {
			return nil, fmt.Errorf("failed to create area controller: %v", err)
		}

		v1.POST("/area/new", areaCtrl.CreateArea)
		v1.DELETE("/area/:id", areaCtrl.DeleteArea)
		/*** AUTH ***/
		//v1.GET("/auth/auth/redirect", authCtrl.GetUserToken)
		/*** ENDPOINT AUTH ***/

		/*protected*/
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{})
	})

	return r, nil
}
