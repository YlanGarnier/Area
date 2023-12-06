package action

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/log"
	"strconv"
	"strings"

	"github.com/lenismtho/area/pkg/core/utils"
	"github.com/lenismtho/area/pkg/protogen"
)

type HTTPServiceAction struct {
	protogen.HTTPServiceActionClient
}

type HTTPAct string

const (
	watcher HTTPAct = "watcher"
)

type WatcherReq struct {
	Url      string    `json:"url"`
	Expected []float64 `json:"expected"`
}

func (s *HTTPServiceAction) ActionHTTPRouter(ctx context.Context, route HTTPAct, param map[string]any, respType protogen.Format_Kind, areaID uint32) error {
	expected, err := utils.MapErr(strings.Split(param["expected"].(string), ","), func(s string) (float64, error) {
		i, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		return float64(i), nil
	})
	if err != nil {
		return fmt.Errorf("failed to parse expected: %v", err)
	}
	switch route {
	case watcher:
		req := WatcherReq{
			Url:      param["url"].(string),
			Expected: expected,
		}
		_, err := s.RegisterWatcher(ctx, &protogen.RegisterWatcher_Request{
			Id:           areaID,
			ResponseType: respType,
			Url:          req.Url,
			Expected: utils.Map(req.Expected, func(i float64) uint32 {
				return uint32(i)
			}),
		})
		if err != nil {
			return fmt.Errorf("failed to register watcher: %v", err)
		}
	default:
		return fmt.Errorf("unknown route: %s", route)
	}
	return nil
}

/* ---- DiscordServiceAction ---- */

type DiscordServiceAction struct {
	protogen.DiscordServiceActionClient
}

type DiscordAct string

const (
	newGuild DiscordAct = "newGuild"
)

type NewGuildReq struct {
	Token string `json:"token"`
}

func (s *DiscordServiceAction) ActionDiscordRouter(
	ctx context.Context,
	route DiscordAct,
	param map[string]any,
	respType protogen.Format_Kind,
	areaID uint32,
	token string,
) error {
	switch route {
	case newGuild:
		_, err := s.RegisterJoinChannelWatcher(ctx, &protogen.InviteWatcher_Request{
			Id:           areaID,
			ResponseType: respType,
			Token:        token,
		})
		if err != nil {
			return fmt.Errorf("failed to register join channel watcher: %v", err)
		}
	default:
		return fmt.Errorf("unknown route: %s", route)
	}
	return nil
}

/* ---- EthereumServiceAction ---- */

type EthereumServiceAction struct {
	protogen.EthereumServiceActionClient
}

type EthereumAct string

const (
	watchTransaction EthereumAct = "watchTransaction"
	eventWatcher     EthereumAct = "eventWatcher"
)

func (s *EthereumServiceAction) ActionEthereumRouter(
	ctx context.Context,
	route EthereumAct,
	param map[string]any,
	respType protogen.Format_Kind,
	areaID uint32,
) error {
	switch route {
	case watchTransaction:
		_, err := s.RegisterAddresseWatcher(ctx, &protogen.AddressWatcher_Request{
			Id:           areaID,
			ResponseType: respType,
			Address:      param["address"].(string),
		})
		if err != nil {
			return fmt.Errorf("failed to register address watcher: %v", err)
		}
	case eventWatcher:
		log.Info("type ", respType)
		_, err := s.RegisterEventWatcher(ctx, &protogen.EventWatcher_Request{
			Id:           areaID,
			ResponseType: respType,
			Event:        param["event"].(string),
		})
		if err != nil {
			return fmt.Errorf("failed to register event watcher: %v", err)
		}
	default:
		return fmt.Errorf("unknown route: %s", route)
	}
	return nil
}

/* ---- SpotifyServiceAction ---- */

type SpotifyServiceAction struct {
	protogen.SpotifyServiceActionClient
}

type SpotifyAct string

const (
	watchArtist SpotifyAct = "watchArtist"
	watchSong   SpotifyAct = "watchSong"
)

func (s *SpotifyServiceAction) ActionSpotifyRouter(
	ctx context.Context,
	route SpotifyAct,
	param map[string]any,
	respType protogen.Format_Kind,
	areaID uint32,
	token string,
) error {
	switch route {
	case watchArtist:
		_, err := s.WatchArtist(ctx, &protogen.Watcher_Request{
			Id:           areaID,
			ResponseType: respType,
			Token:        token,
			Name:         param["artist"].(string),
		})
		if err != nil {
			return fmt.Errorf("failed to register artist watcher: %v", err)
		}
	case watchSong:
		_, err := s.WatchSong(ctx, &protogen.Watcher_Request{
			Id:           areaID,
			ResponseType: respType,
			Token:        token,
			Name:         param["song"].(string),
		})
		if err != nil {
			return fmt.Errorf("failed to register song watcher: %v", err)
		}
	default:
		return fmt.Errorf("unknown route: %s", route)
	}
	return nil
}

/* ---- GmailServiceAaction ---- */
type GmailServiceAction struct {
	protogen.GmailServiceActionClient
}

type GmailAct string

const (
	newEmail                   GmailAct = "NewEmail"
	newEmailWithSender         GmailAct = "NewEmailWithSender"
	newEmailAtDate             GmailAct = "NewEmailAtDate"
	newEmailAtDateWithSender   GmailAct = "NewEmailAtDateWithSender"
	newDraft                   GmailAct = "NewDraft"
	newDraftAtDate             GmailAct = "NewDraftAtDate"
	newDraftWithReceiver       GmailAct = "NewDraftWithReceiver"
	newDraftAtDateWithReceiver GmailAct = "NewDraftAtDateWithReceiver"
	newLabel                   GmailAct = "NewLabel"
	newLabelWithName           GmailAct = "NewLabelWithName"
	newEmailInLabel            GmailAct = "NewEmailInLabel"
)

func (s *GmailServiceAction) ActionGmailRouter(
	ctx context.Context,
	route GmailAct,
	param map[string]any,
	respType protogen.Format_Kind,
	areaID uint32,
	token string,
) error {
	switch route {
	case newEmail:
		_, err := s.RegisterNewEmail(ctx, &protogen.NewEmail_Request{
			Id:           areaID,
			ResponseType: respType,
			Token:        token,
		})
		if err != nil {
			return fmt.Errorf("failed to register new email %v", err)
		}
	case newEmailWithSender:
		_, err := s.RegisterNewEmailWithSender(ctx, &protogen.NewEmailWithSender_Request{
			Id:           areaID,
			ResponseType: respType,
			Token:        token,
			Sender:       param["sender"].(string),
		})
		if err != nil {
			return fmt.Errorf("failed to register new email with sender %v", err)
		}
	case newEmailAtDate:
		_, err := s.RegisterNewEmailAtDate(ctx, &protogen.NewEmailAtDate_Request{
			Id:           areaID,
			ResponseType: respType,
			Token:        token,
			Date:         param["date"].(string),
		})
		if err != nil {
			return fmt.Errorf("failed to register new email at date %v", err)
		}
	case newEmailAtDateWithSender:
		_, err := s.RegisterNewEmailAtDateWithSender(ctx, &protogen.NewEmailAtDateWithSender_Request{
			Id:           areaID,
			ResponseType: respType,
			Token:        token,
			Date:         param["date"].(string),
			Sender:       param["sender"].(string),
		})
		if err != nil {
			return fmt.Errorf("failed to register new email at date with sender %v", err)
		}
	case newDraft:
		_, err := s.RegisterNewDraft(ctx, &protogen.NewDraft_Request{
			Id:           areaID,
			ResponseType: respType,
			Token:        token,
		})
		if err != nil {
			return fmt.Errorf("failed to register new draft %v", err)
		}
	case newDraftAtDate:
		_, err := s.RegisterNewDraftAtDate(ctx, &protogen.NewDraftAtDate_Request{
			Id:           areaID,
			ResponseType: respType,
			Token:        token,
			Date:         param["date"].(string),
		})
		if err != nil {
			return fmt.Errorf("failed to register new draft at date %v", err)
		}
	case newDraftWithReceiver:
		_, err := s.RegisterNewDraftWithReceiver(ctx, &protogen.NewDraftWithReceiver_Request{
			Id:           areaID,
			ResponseType: respType,
			Token:        token,
			Receiver:     param["receiver"].(string),
		})
		if err != nil {
			return fmt.Errorf("failed to register new draft with sender %v", err)
		}
	case newDraftAtDateWithReceiver:
		_, err := s.RegisterNewDraftAtDateWithReceiver(ctx, &protogen.NewDraftAtDateWithReceiver_Request{
			Id:           areaID,
			ResponseType: respType,
			Token:        token,
			Date:         param["date"].(string),
			Receiver:     param["receiver"].(string),
		})
		if err != nil {
			return fmt.Errorf("failed to register new draft at date with sender %v", err)
		}
	case newLabel:
		_, err := s.RegisterNewLabel(ctx, &protogen.NewLabel_Request{
			Id:           areaID,
			ResponseType: respType,
			Token:        token,
		})
		if err != nil {
			return fmt.Errorf("failed to register new label %v", err)
		}
	case newLabelWithName:
		_, err := s.RegisterNewLabelWithName(ctx, &protogen.NewLabelWithName_Request{
			Id:           areaID,
			ResponseType: respType,
			Token:        token,
			Name:         param["name"].(string),
		})
		if err != nil {
			return fmt.Errorf("failed to register new label with name %v", err)
		}
	case newEmailInLabel:
		_, err := s.RegisterNewEmailInLabel(ctx, &protogen.NewEmailInLabel_Request{
			Id:           areaID,
			ResponseType: respType,
			Token:        token,
			Label:        param["label"].(string),
		})
		if err != nil {
			return fmt.Errorf("failed to register new email in label %v", err)
		}
	default:
		return fmt.Errorf("unknown route: %s", route)
	}
	return nil
}
