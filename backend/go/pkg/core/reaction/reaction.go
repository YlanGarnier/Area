package reaction

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"google.golang.org/protobuf/proto"

	"github.com/lenismtho/area/pkg/protogen"
)

type Reaction interface {
	React(ctx context.Context, route string, rawMsg []byte, b *protogen.Base) error
}

var RespTypes = map[ReaService]map[string]protogen.Format_Kind{
	GithubService: {
		string(CreateIssue): protogen.Format_GHIncidentReportKind,
		string(CreateRepo):  protogen.Format_OnlyTitleKind,
		string(CreateRepoNamed): protogen.Format_NoParamKind,
		string(CreateGists):     protogen.Format_ManyFilesKind,
		string(AddTopics):       protogen.Format_TagsKind,
	},
	TwitterService: {
		string(PostTweet):                protogen.Format_NoParamKind,
		string(PostTweetWithContent):     protogen.Format_OnlyTitleKind,
		string(PostTweetWithPoll):        protogen.Format_NoParamKind,
		string(PostTweetWithContentPoll): protogen.Format_OnlyTitleKind,
	},
	TwitchService: {
		string(SendDefaultMessage): protogen.Format_OnlyTitleKind,
		string(SendMessage):        protogen.Format_OnlyTitleKind,
	},
	DropboxService: {
		string(CreateFolder): protogen.Format_OnlyTitleKind,
		string(CreateFile):   protogen.Format_GHIncidentReportKind,
		string(CreateTag):    protogen.Format_OnlyTitleKind,
	},
	MiroService: {
		string(CreateBoard):      protogen.Format_OnlyTitleKind,
		string(CreateTags):       protogen.Format_OnlyTitleKind,
		string(CreateStickyCard): protogen.Format_OnlyTitleKind,
		string(CreateCardItem):   protogen.Format_GHIncidentReportKind,
	},
	LinkedinService: {
		string(CreateDefaultPost): protogen.Format_OnlyTitleKind,
		string(CreatePost):        protogen.Format_OnlyTitleKind,
	},
	NotionService: {
		string(CreateDefaultComment): protogen.Format_OnlyTitleKind,
		string(CreateComment):        protogen.Format_GHIncidentReportKind,
		string(CreateDefaultPage):    protogen.Format_OnlyTitleKind,
		string(CreatePage):           protogen.Format_GHIncidentReportKind,
		string(CreateBlock):          protogen.Format_GHIncidentReportKind,
	},
}

func GetRespType(service ReaService, route string) protogen.Format_Kind {
	if len(RespTypes[service]) == 0 {
		return protogen.Format_UndefinedKind
	}
	return RespTypes[service][route]
}

type ReaService string

const (
	GithubService ReaService = "github"
)

type GHreact string

const (
	CreateIssue     GHreact = "createIssue"
	CreateRepo      GHreact = "createRepo"
	CreateRepoNamed GHreact = "createRepoNamed"
	CreateGists     GHreact = "createGists"
	AddTopics       GHreact = "addTopics"
	CreateBranch    GHreact = "createBranch"
)

type GHServiceReaction struct {
	protogen.GHServiceReactionClient
}

func (g *GHServiceReaction) React(ctx context.Context, route string, rawMsg []byte, b *protogen.Base) error {
	log.Println("reacting")
	switch GHreact(route) {
	case CreateIssue:
		fmtt := &protogen.Format_GHIncidentReport{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = g.CreateIssue(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	case CreateRepo:
		fmtt := &protogen.Format_OnlyTitle{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = g.CreateRepo(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	case CreateRepoNamed:
		fmtt := &protogen.Format_NoParam{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = g.CreateRepoNamed(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	case CreateGists:
		fmtt := &protogen.Format_ManyFiles{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = g.CreateGists(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	case AddTopics:
		fmtt := &protogen.Format_Tags{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = g.AddTags(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	default:
		return fmt.Errorf("unknown route: %s", route)
	}

	return nil
}

const (
	DropboxService ReaService = "dropbox"
)

type DBreact string

const (
	CreateFolder DBreact = "createFolder"
	CreateFile   DBreact = "createFile"
	CreateTag    DBreact = "createTag"
)

type DropboxServiceReaction struct {
	protogen.DropboxServiceReactionClient
}

func (db *DropboxServiceReaction) React(ctx context.Context, route string, rawMsg []byte, b *protogen.Base) error {
	switch DBreact(route) {
	case CreateFolder:
		fmtt := &protogen.Format_OnlyTitle{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = db.CreateFolder(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	case CreateFile:
		fmtt := &protogen.Format_GHIncidentReport{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = db.CreateFile(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	case CreateTag:
		fmtt := &protogen.Format_OnlyTitle{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = db.CreateTag(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	default:
		return fmt.Errorf("unknown route: %s", route)
	}

	return nil
}

// TWITCH SERVICE

const (
	TwitchService ReaService = "twitch"
)

type TwitchReact string

const (
	SendDefaultMessage TwitchReact = "sendDefaultMessage"
	SendMessage        TwitchReact = "sendMessage"
)

type TwitchServiceReaction struct {
	protogen.TwitchServiceReactionClient
}

func (tw *TwitchServiceReaction) React(ctx context.Context, route string, rawMsg []byte, b *protogen.Base) error {

	switch TwitchReact(route) {
	case SendDefaultMessage:
		fmtt := &protogen.Format_OnlyTitle{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = tw.SendDefaultMessage(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	case SendMessage:
		fmtt := &protogen.Format_OnlyTitle{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = tw.SendMessage(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	default:
		return fmt.Errorf("unknown route: %s", route)
	}
	return nil
}

const TwitterService ReaService = "twitter"

type TwitterReact string

const (
	PostTweet                TwitterReact = "postTweet"
	PostTweetWithContent     TwitterReact = "postTweetWithContent"
	PostTweetWithPoll        TwitterReact = "postTweetWithPoll"
	PostTweetWithContentPoll TwitterReact = "postTweetWithContentPoll"
)

type TwitterServiceReaction struct {
	protogen.TwitterServiceReactionClient
}

func (t *TwitterServiceReaction) React(ctx context.Context, route string, rawMsg []byte, b *protogen.Base) error {
	switch TwitterReact(route) {
	case PostTweet:
		log.Println("reacting")
		fmtt := &protogen.Format_NoParam{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = t.PostTweet(ctx, fmtt)
		if err != nil {
			return err
		}
	case PostTweetWithContent:
		log.Println("reacting")
		fmtt := &protogen.Format_OnlyTitle{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = t.PostTweetWithContent(ctx, fmtt)
		if err != nil {
			return err
		}
	case PostTweetWithPoll:
		log.Println("reacting")
		fmtt := &protogen.Format_NoParam{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = t.PostTweetWithPoll(ctx, fmtt)
		if err != nil {
			return err
		}
	case PostTweetWithContentPoll:
		log.Println("reacting")
		fmtt := &protogen.Format_OnlyTitle{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = t.PostTweetWithContentWithPoll(ctx, fmtt)
		if err != nil {
			return err
		}
	default:
		log.Println("not reacting")
		return fmt.Errorf("unknown route: %s", route)
	}

	return nil
}

const (
	MiroService ReaService = "miro"
)

const (
	CreateBoard      MiroReact = "createBoard"
	CreateTags       MiroReact = "createTags"
	CreateStickyCard MiroReact = "createStickyCard"
	CreateCardItem   MiroReact = "createCardItem"
)

type MiroReact string

type MiroServiceReaction struct {
	protogen.MiroServiceReactionClient
}

func (m *MiroServiceReaction) React(ctx context.Context, route string, rawMsg []byte, b *protogen.Base) error {
	switch MiroReact(route) {
	case CreateBoard:
		fmtt := &protogen.Format_OnlyTitle{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = m.CreateBoard(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	case CreateTags:
		fmtt := &protogen.Format_OnlyTitle{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = m.CreateTags(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	case CreateStickyCard:
		fmtt := &protogen.Format_OnlyTitle{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = m.CreateStickyCard(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	case CreateCardItem:
		fmtt := &protogen.Format_GHIncidentReport{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = m.CreateCardItem(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	}
	return nil
}

const (
	LinkedinService ReaService = "linkedin"
)

type LKreact string

const (
	CreateDefaultPost LKreact = "createDefaultPost"
	CreatePost        LKreact = "createPost"
)

type LKServiceReaction struct {
	protogen.LinkedinServiceReactionClient
}

func (lk *LKServiceReaction) React(ctx context.Context, route string, rawMsg []byte, b *protogen.Base) error {

	switch LKreact(route) {
	case CreateDefaultPost:
		fmtt := &protogen.Format_OnlyTitle{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = lk.CreateDefaultPost(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	case CreatePost:
		fmtt := &protogen.Format_OnlyTitle{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = lk.CreatePost(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	default:
		return fmt.Errorf("unknown route: %s", route)
	}

	return nil
}

const (
	NotionService ReaService = "notion"
)

type NotionReact string

const (
	CreateDefaultComment NotionReact = "createDefaultComment"
	CreateComment        NotionReact = "createComment"
	CreateDefaultPage    NotionReact = "createDefaultPage"
	CreatePage           NotionReact = "createPage"
	CreateBlock          NotionReact = "createBlock"
)

type NotionServiceReaction struct {
	protogen.NotionServiceReactionClient
}

func (no *NotionServiceReaction) React(ctx context.Context, route string, rawMsg []byte, b *protogen.Base) error {
	switch NotionReact(route) {
	case CreateDefaultComment:
		fmtt := &protogen.Format_OnlyTitle{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = no.CreateDefaultComment(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	case CreateComment:
		fmtt := &protogen.Format_GHIncidentReport{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = no.CreateComment(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	case CreateDefaultPage:
		fmtt := &protogen.Format_OnlyTitle{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = no.CreateDefaultPage(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	case CreatePage:
		fmtt := &protogen.Format_GHIncidentReport{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = no.CreatePage(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	case CreateBlock:
		fmtt := &protogen.Format_GHIncidentReport{}
		err := proto.Unmarshal(rawMsg, fmtt)
		if err != nil {
			return err
		}
		fmtt.Base = b
		_, err = no.CreateBlock(ctx, fmtt)
		if err != nil {
			return err
		}
		fmt.Println("received", fmtt)
	default:
		return fmt.Errorf("unknown route: %s", route)
	}
	return nil
}
