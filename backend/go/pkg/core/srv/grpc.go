package srv

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/lenismtho/area/pkg/core/models"
	"github.com/lenismtho/area/pkg/core/reaction"
	"github.com/lenismtho/area/pkg/protogen"
)

type coreDatabase interface {
	GetAreaByID(id uint) (*models.Area, error)
}

type CoreServer struct {
	protogen.UnimplementedCoreServiceServer

	cGH reaction.GHServiceReaction
	cTW reaction.TwitchServiceReaction
	cDB reaction.DropboxServiceReaction
	cLK reaction.LKServiceReaction
	cTwitter reaction.TwitterServiceReaction
	cMiro reaction.MiroServiceReaction
	cNO reaction.NotionServiceReaction

	db  coreDatabase
}

var _ protogen.CoreServiceServer = (*CoreServer)(nil)

func NewCoreServer(db coreDatabase, cGH protogen.GHServiceReactionClient, cTW protogen.TwitchServiceReactionClient, cDB protogen.DropboxServiceReactionClient, cTwitter protogen.TwitterServiceReactionClient, cMiro protogen.MiroServiceReactionClient, cLK protogen.LinkedinServiceReactionClient, cNO protogen.NotionServiceReactionClient) (*CoreServer, error) {
	return &CoreServer{
		db:    db,
		cGH:   reaction.GHServiceReaction{GHServiceReactionClient: cGH},
		cTW:   reaction.TwitchServiceReaction{TwitchServiceReactionClient: cTW},
		cDB:   reaction.DropboxServiceReaction{DropboxServiceReactionClient: cDB},
		cTwitter: reaction.TwitterServiceReaction{TwitterServiceReactionClient: cTwitter},
		cMiro: reaction.MiroServiceReaction{MiroServiceReactionClient: cMiro},
		cLK: reaction.LKServiceReaction{LinkedinServiceReactionClient: cLK},
		cNO: reaction.NotionServiceReaction{NotionServiceReactionClient: cNO},
	}, nil
}

func (s *CoreServer) ForwardAction(ctx context.Context, req *protogen.ForwardActionReq) (*protogen.Empty, error) {
	area, err := s.db.GetAreaByID(uint(req.GetId()))
	if err != nil {
		return nil, err
	}

	switch reaction.ReaService(area.ReaServices) {
	case reaction.GithubService:
		err = s.cGH.React(ctx, area.Route, req.GetData(), area.Base)
		if err != nil {
			log.WithField("err", err).Error("failed to react")
			return nil, fmt.Errorf("failed to react: %w", err)
		}
	case reaction.TwitterService:
		fmt.Println("twitter")
		err = s.cTwitter.React(ctx, area.Route, req.GetData(), area.Base)
		if err != nil {
			log.WithField("err", err).Error("failed to react")
			return nil, fmt.Errorf("failed to react: %w", err)
		}
	case reaction.TwitchService:
		err = s.cTW.React(ctx, area.Route, req.GetData(), area.Base)
		if err != nil {
			return nil, fmt.Errorf("failed to react: %w", err)
		}
	case reaction.DropboxService:
		err = s.cDB.React(ctx, area.Route, req.GetData(), area.Base)
		if err != nil {
			return nil, fmt.Errorf("failed to react: %w", err)
		}
	case reaction.MiroService:
		err = s.cMiro.React(ctx, area.Route, req.GetData(), area.Base)
		if err != nil {
			return nil, fmt.Errorf("failed to react: %w", err)
		}
	case reaction.LinkedinService:
		err = s.cLK.React(ctx, area.Route, req.GetData(), area.Base)
		if err != nil {
			return nil, fmt.Errorf("failed to react: %w", err)
		}
	case reaction.NotionService:
		err = s.cNO.React(ctx, area.Route, req.GetData(), area.Base)
		if err != nil {
			return nil, fmt.Errorf("failed to react: %w", err)
		}
	default:
		return nil, fmt.Errorf("unknown route: %s", area.Route)
	}

	return &protogen.Empty{}, nil
}
