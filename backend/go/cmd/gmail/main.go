package gmail

import (
	"context"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"

	"github.com/lenismtho/area/pkg/core/utils"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"github.com/lenismtho/area/pkg/protogen"
)

type pollsNewEmail struct {
	id       uint32
	respType protogen.Format_Kind
	token    string

	LastEmail Email
}

type pollsNewDraft struct {
	id       uint32
	respType protogen.Format_Kind
	token    string

	LastDraft Draft
}

type pollsNewEmailWithSender struct {
	id       uint32
	respType protogen.Format_Kind
	token    string
	sender   string

	LastEmail Email
}

type pollsNewEmailAtDate struct {
	id       uint32
	respType protogen.Format_Kind
	token    string
	date     time.Time

	LastEmail Email
}

type pollsNewEmailAtDateWithSender struct {
	id       uint32
	respType protogen.Format_Kind
	token    string
	sender   string
	date     time.Time

	LastEmail Email
}

type pollsNewDraftAtDate struct {
	id       uint32
	respType protogen.Format_Kind
	token    string
	date     time.Time

	LastDraft Draft
}

type pollsNewDraftWithReceiver struct {
	id       uint32
	respType protogen.Format_Kind
	token    string
	receiver string

	LastDraft Draft
}

type pollsNewDraftAtDateWithReceiver struct {
	id       uint32
	respType protogen.Format_Kind
	token    string
	date     time.Time
	receiver string

	LastDraft Draft
}

type pollsNewLabel struct {
	id       uint32
	respType protogen.Format_Kind
	token    string

	Labels Labels
}

type pollsNewLabelWithName struct {
	id       uint32
	respType protogen.Format_Kind
	token    string
	name     string

	Labels Labels
}

type pollsNewEmailInLabel struct {
	id       uint32
	respType protogen.Format_Kind
	token    string
	label    string

	MessagesTotal int
}

type Labels struct {
	Labels []Label `json:"labels"`
}

type Label struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type LabelContent struct {
	ID            string `json:"id"`
	MessagesTotal int    `json:"messagesTotal"`
}

type Draft struct {
	ID      string `json:"id"`
	Message Email  `json:"message"`
}

type DraftContent struct {
	ID      string       `json:"id"`
	Snippet string       `json:"snippet"`
	Message EmailContent `json:"message"`
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Payload struct {
	Headers []Header `json:"headers"`
}

type EmailContent struct {
	ID      string  `json:"id"`
	Snippet string  `json:"snippet"`
	Payload Payload `json:"payload"`
}

type Email struct {
	ID       string `json:"id"`
	ThreadID string `json:"threadId"`
}

type Emails struct {
	Messages []Email `json:"messages"`
}

type Drafts struct {
	Drafts []Draft `json:"drafts"`
}

func Main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error: failed to load the env file")
		return
	}
	newEmailA := NewEmailAction{
		watchersNewEmail: &[]pollsNewEmail{},
		mut:              &sync.Mutex{},
	}
	newEmailWithSenderA := NewEmailWithSenderAction{
		watchers: &[]pollsNewEmailWithSender{},
		mut:      &sync.Mutex{},
	}
	newEmailAtDateA := NewEmailAtDateAction{
		watchers: &[]pollsNewEmailAtDate{},
		mut:      &sync.Mutex{},
	}
	newEmailAtDateWithSenderA := NewEmailAtDateWithSenderAction{
		watchers: &[]pollsNewEmailAtDateWithSender{},
		mut:      &sync.Mutex{},
	}
	newDraftA := NewDraftAction{
		watchers: &[]pollsNewDraft{},
		mut:      &sync.Mutex{},
	}
	newDraftAtDateA := NewDraftAtDateAction{
		watchers: &[]pollsNewDraftAtDate{},
		mut:      &sync.Mutex{},
	}
	newDraftWithReceiverA := NewDraftWithReceiverAction{
		watchers: &[]pollsNewDraftWithReceiver{},
		mut:      &sync.Mutex{},
	}
	newDraftAtDateWithReceiverA := NewDraftAtDateWithReceiverAction{
		watchers: &[]pollsNewDraftAtDateWithReceiver{},
		mut:      &sync.Mutex{},
	}
	newLabelA := NewLabelAction{
		watchers: &[]pollsNewLabel{},
		mut:      &sync.Mutex{},
	}
	newLabelWithNameA := NewLabelWithNameAction{
		watchers: &[]pollsNewLabelWithName{},
		mut:      &sync.Mutex{},
	}
	newEmailInLabelA := NewEmailInLabelAction{
		watchers: &[]pollsNewEmailInLabel{},
		mut:      &sync.Mutex{},
	}
	lis, err := net.Listen("tcp", os.Getenv("GMAIL_ACTION_SERVICE"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	gmailServer, err := NewGmailActionServer(&newEmailA, &newEmailWithSenderA, &newEmailAtDateA, &newEmailAtDateWithSenderA, &newDraftA,
		&newDraftAtDateA, &newDraftWithReceiverA, &newDraftAtDateWithReceiverA, &newLabelA, &newLabelWithNameA, &newEmailInLabelA)
	if err != nil {
		log.Fatal(err)
	}
	protogen.RegisterGmailServiceActionServer(grpcServer, gmailServer)

	coreConn, err := grpc.Dial(os.Getenv("CORE"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	coreClient := protogen.NewCoreServiceClient(coreConn)

	go func() {
		err = grpcServer.Serve(lis)
	}()

	log.Info("Starting service")
	go func() {
		fmt.Println("----")
		err := eventNewEmail(newEmailA, coreClient)
		if err != nil {
			log.WithField("error", err).Error("failed to start event watcher routine")
		}
	}()
	go func() {
		fmt.Println("====")
		err := eventNewEmailWithSender(newEmailWithSenderA, coreClient)
		if err != nil {
			log.WithField("error", err).Error("failed to start event watcher routine")
		}
	}()
	go func() {
		err := eventNewEmailAtDate(newEmailAtDateA, coreClient)
		if err != nil {
			log.WithField("error", err).Error("failed to start event watcher routine")
		}
	}()
	go func() {
		err := eventNewEmailAtDateWithSender(newEmailAtDateWithSenderA, coreClient)
		if err != nil {
			log.WithField("error", err).Error("failed to start event watcher routine")
		}
	}()
	go func() {
		err := eventNewDraft(newDraftA, coreClient)
		if err != nil {
			log.WithField("error", err).Error("failed to start event watcher routine")
		}
	}()
	go func() {
		err := eventNewDraftAtDate(newDraftAtDateA, coreClient)
		if err != nil {
			log.WithField("error", err).Error("failed to start event watcher routine")
		}
	}()
	go func() {
		err := eventNewDraftWithReceiver(newDraftWithReceiverA, coreClient)
		if err != nil {
			log.WithField("error", err).Error("failed to start event watcher routine")
		}
	}()
	go func() {
		err := eventNewDraftAtDateWithReceiver(newDraftAtDateWithReceiverA, coreClient)
		if err != nil {
			log.WithField("error", err).Error("failed to start event watcher routine")
		}
	}()
	go func() {
		err := eventNewLabel(newLabelA, coreClient)
		if err != nil {
			log.WithField("error", err).Error("failed to start event watcher routine")
		}
	}()
	go func() {
		err := eventNewLabelWithName(newLabelWithNameA, coreClient)
		if err != nil {
			log.WithField("error", err).Error("failed to start event watcher routine")
		}
	}()
	func() {
		err := eventNewEmailInLabel(newEmailInLabelA, coreClient)
		if err != nil {
			log.WithField("error", err).Error("failed to start event watcher routine")
		}
	}()
}

func eventNewEmailInLabel(newEmailInLabel NewEmailInLabelAction, coreClient protogen.CoreServiceClient) error {
	for {
		time.Sleep(10 * time.Second)
		newEmailInLabel.mut.Lock()
		for id, poll := range *newEmailInLabel.watchers {
			labelContent, err := utils.GetFromJsonReq[LabelContent](os.Getenv("GMAIL_API_URL")+"/users/me/labels/"+poll.label, utils.GET, "", []utils.Header{
				{
					Key:   "Authorization",
					Value: fmt.Sprintf("Bearer %s", poll.token),
				},
			}, "")
			if err != nil {
				log.WithField("error", err).Error("failed to get label content")
				return err
			}
			fmt.Println("labelContent = ", labelContent)
			if labelContent.MessagesTotal != poll.MessagesTotal {
				raw, err := formatNewEmailInLabel(labelContent, poll.respType)
				if err != nil {
					log.WithField("error", err).Error("failed to format new email in label action")
					return err
				}
				_, err = coreClient.ForwardAction(context.Background(), &protogen.ForwardActionReq{
					Id:   poll.id,
					Type: poll.respType,
					Data: raw,
				})
				if err != nil {
					log.WithField("error", err).Error("failed to forward")
					continue
				}
				(*newEmailInLabel.watchers)[id].MessagesTotal = labelContent.MessagesTotal
			}

		}
		newEmailInLabel.mut.Unlock()
	}
}

func eventNewLabelWithName(newLabelWithName NewLabelWithNameAction, coreClient protogen.CoreServiceClient) error {
	for {
		time.Sleep(10 * time.Second)
		newLabelWithName.mut.Lock()
		for id, poll := range *newLabelWithName.watchers {
			labels, err := utils.GetFromJsonReq[Labels](os.Getenv("GMAIL_API_URL")+"/users/me/labels", utils.GET, "", []utils.Header{
				{
					Key:   "Authorization",
					Value: fmt.Sprintf("Bearer %s", poll.token),
				},
			}, "")
			if err != nil || len(labels.Labels) == 0 {
				log.WithField("error", err).Error("failed to get labels")
				return err
			}
			if len(labels.Labels) != len(poll.Labels.Labels) {
				for _, label := range labels.Labels {
					if label.Name == poll.name {
						raw, err := formatNewLabelWithNameAction(labels.Labels[len(labels.Labels)-1], poll.respType)
						if err != nil {
							log.WithField("error", err).Error("failed to format new label action")
							return err
						}
						_, err = coreClient.ForwardAction(context.Background(), &protogen.ForwardActionReq{
							Id:   poll.id,
							Type: poll.respType,
							Data: raw,
						})
						if err != nil {
							log.WithField("error", err).Error("failed to forward")
							continue
						}
					}
				}
				(*newLabelWithName.watchers)[id].Labels = labels
			}
		}
		newLabelWithName.mut.Unlock()
	}
}

func eventNewLabel(newLabel NewLabelAction, coreClient protogen.CoreServiceClient) error {
	for {
		time.Sleep(10 * time.Second)
		newLabel.mut.Lock()
		for id, poll := range *newLabel.watchers {
			labelFound := false
			labels, err := utils.GetFromJsonReq[Labels](os.Getenv("GMAIL_API_URL")+"/users/me/labels", utils.GET, "", []utils.Header{
				{
					Key:   "Authorization",
					Value: fmt.Sprintf("Bearer %s", poll.token),
				},
			}, "")
			if err != nil || len(labels.Labels) == 0 {
				log.WithField("error", err).Error("failed to get labels")
				return err
			}
			if len(labels.Labels) != len(poll.Labels.Labels) {
				for _, label := range labels.Labels {
					for _, oldLabel := range poll.Labels.Labels {
						if label.Name == oldLabel.Name {
							labelFound = true
						}
					}
					if !labelFound {
						raw, err := formatNewLabelAction(labels, poll.respType)
						if err != nil {
							log.WithField("error", err).Error("failed to format new label action")
							return err
						}
						_, err = coreClient.ForwardAction(context.Background(), &protogen.ForwardActionReq{
							Id:   poll.id,
							Type: poll.respType,
							Data: raw,
						})
						if err != nil {
							log.WithField("error", err).Error("failed to forward")
							continue
						}
					}
					labelFound = false
				}
				(*newLabel.watchers)[id].Labels = labels
			}
		}
		newLabel.mut.Unlock()
	}
}

func eventNewDraftAtDateWithReceiver(newDraftAtDateWithReceiver NewDraftAtDateWithReceiverAction, coreClient protogen.CoreServiceClient) error {
	for {
		time.Sleep(10 * time.Second)
		newDraftAtDateWithReceiver.mut.Lock()
		for id, poll := range *newDraftAtDateWithReceiver.watchers {
			dateOk := false
			drafts, err := utils.GetFromJsonReq[Drafts](os.Getenv("GMAIL_API_URL")+"/users/me/drafts", utils.GET, "", []utils.Header{
				{
					Key:   "Authorization",
					Value: fmt.Sprintf("Bearer %s", poll.token),
				},
			}, "")
			if err != nil || len(drafts.Drafts) == 0 {
				log.WithField("error", err).Error("failed to get drafts")
				return err
			}
			if drafts.Drafts[0].ID != poll.LastDraft.ID {
				draftContent, err := utils.GetFromJsonReq[DraftContent](os.Getenv("GMAIL_API_URL")+"/users/me/drafts/"+drafts.Drafts[0].ID, utils.GET, "", []utils.Header{
					{
						Key:   "Authorization",
						Value: fmt.Sprintf("Bearer %s", poll.token),
					},
				}, "")
				fmt.Println("draftContent = ", draftContent)
				if err != nil {
					log.WithField("error", err).Error("failed to get draft content")
					return err
				}
				for _, header := range draftContent.Message.Payload.Headers {
					if header.Name == "Date" {
						dateHeaderValue := checkFormatDate(header.Value)
						date, err := time.Parse(time.RFC1123Z, dateHeaderValue)
						if err != nil {
							log.WithField("error", err).Error("failed to parse date")
							return err
						}
						dateFormat, err := time.Parse(time.DateOnly, date.Format(time.DateOnly))
						if err != nil {
							log.WithField("error", err).Error("failed to parse date")
							return err
						}
						if dateFormat.Equal(poll.date) {
							dateOk = true
						}
					}
					if header.Name == "To" && strings.Contains(header.Value, poll.receiver) {
						if !dateOk {
							break
						}
						raw, err := formatNewDraftAtDateWithReceiverAction(poll.date, poll.receiver, drafts.Drafts[0], poll.respType)
						if err != nil {
							log.WithField("error", err).Error("failed to format new draft action")
							return err
						}
						_, err = coreClient.ForwardAction(context.Background(), &protogen.ForwardActionReq{
							Id:   poll.id,
							Type: poll.respType,
							Data: raw,
						})
						if err != nil {
							log.WithField("error", err).Error("failed to forward")
							continue
						}
					}
				}
				(*newDraftAtDateWithReceiver.watchers)[id].LastDraft = drafts.Drafts[0]
			}
		}
		newDraftAtDateWithReceiver.mut.Unlock()
	}
}

func eventNewDraftWithReceiver(newDraftWithReceiver NewDraftWithReceiverAction, coreClient protogen.CoreServiceClient) error {
	for {
		time.Sleep(10 * time.Second)
		newDraftWithReceiver.mut.Lock()
		for id, poll := range *newDraftWithReceiver.watchers {
			drafts, err := utils.GetFromJsonReq[Drafts](os.Getenv("GMAIL_API_URL")+"/users/me/drafts", utils.GET, "", []utils.Header{
				{
					Key:   "Authorization",
					Value: fmt.Sprintf("Bearer %s", poll.token),
				},
			}, "")
			if err != nil || len(drafts.Drafts) == 0 {
				log.WithField("error", err).Error("failed to get drafts")
				return err
			}
			fmt.Println("drafts.Drafts[0].ID = ", drafts.Drafts[0].ID, " poll.LastDraft.ID = ", poll.LastDraft.ID)
			if drafts.Drafts[0].ID != poll.LastDraft.ID {
				draftContent, err := utils.GetFromJsonReq[DraftContent](os.Getenv("GMAIL_API_URL")+"/users/me/drafts/"+drafts.Drafts[0].ID, utils.GET, "", []utils.Header{
					{
						Key:   "Authorization",
						Value: fmt.Sprintf("Bearer %s", poll.token),
					},
				}, "")
				fmt.Println("draftContent = ", draftContent)
				if err != nil {
					log.WithField("error", err).Error("failed to get draft content")
					return err
				}
				for _, header := range draftContent.Message.Payload.Headers {
					if header.Name == "To" && strings.Contains(header.Value, poll.receiver) {
						raw, err := formatNewDraftWithReceiverAction(poll.receiver, drafts.Drafts[0], poll.respType)
						if err != nil {
							log.WithField("error", err).Error("failed to format new draft action")
							return err
						}
						_, err = coreClient.ForwardAction(context.Background(), &protogen.ForwardActionReq{
							Id:   poll.id,
							Type: poll.respType,
							Data: raw,
						})
						if err != nil {
							log.WithField("error", err).Error("failed to forward")
							continue
						}
					}
				}
				(*newDraftWithReceiver.watchers)[id].LastDraft = drafts.Drafts[0]
			}
		}
		newDraftWithReceiver.mut.Unlock()
	}
}

func eventNewDraftAtDate(newDraftAtDate NewDraftAtDateAction, coreClient protogen.CoreServiceClient) error {
	for {
		time.Sleep(10 * time.Second)
		newDraftAtDate.mut.Lock()
		for id, poll := range *newDraftAtDate.watchers {
			drafts, err := utils.GetFromJsonReq[Drafts](os.Getenv("GMAIL_API_URL")+"/users/me/drafts", utils.GET, "", []utils.Header{
				{
					Key:   "Authorization",
					Value: fmt.Sprintf("Bearer %s", poll.token),
				},
			}, "")
			if err != nil || len(drafts.Drafts) == 0 {
				log.WithField("error", err).Error("failed to get drafts")
				return err
			}
			if drafts.Drafts[0].ID != poll.LastDraft.ID {
				draftContent, err := utils.GetFromJsonReq[DraftContent](os.Getenv("GMAIL_API_URL")+"/users/me/drafts/"+drafts.Drafts[0].ID, utils.GET, "", []utils.Header{
					{
						Key:   "Authorization",
						Value: fmt.Sprintf("Bearer %s", poll.token),
					},
				}, "")
				fmt.Println("draftContent = ", draftContent)
				if err != nil || len(draftContent.Message.Payload.Headers) == 0 {
					log.WithField("error", err).Error("failed to get draft content")
					return err
				}
				for _, header := range draftContent.Message.Payload.Headers {
					if header.Name == "Date" {
						dateHeaderValue := checkFormatDate(header.Value)
						date, err := time.Parse(time.RFC1123Z, dateHeaderValue)
						if err != nil {
							log.WithField("error", err).Error("failed to parse date")
							return err
						}
						dateFormat, err := time.Parse(time.DateOnly, date.Format(time.DateOnly))
						if err != nil {
							log.WithField("error", err).Error("failed to parse date")
							return err
						}
						if dateFormat.Equal(poll.date) {
							raw, err := formatNewDraftAtDateAction(poll.date, drafts.Drafts[0], poll.respType)
							if err != nil {
								log.WithField("error", err).Error("failed to format new draft action")
								return err
							}
							_, err = coreClient.ForwardAction(context.Background(), &protogen.ForwardActionReq{
								Id:   poll.id,
								Type: poll.respType,
								Data: raw,
							})
							if err != nil {
								log.WithField("error", err).Error("failed to forward")
								continue
							}
						}
					}
				}
				(*newDraftAtDate.watchers)[id].LastDraft = drafts.Drafts[0]
			}
		}
		newDraftAtDate.mut.Unlock()
	}
}

func eventNewDraft(newDraft NewDraftAction, coreClient protogen.CoreServiceClient) error {
	for {
		time.Sleep(10 * time.Second)
		newDraft.mut.Lock()
		for id, poll := range *newDraft.watchers {
			drafts, err := utils.GetFromJsonReq[Drafts](os.Getenv("GMAIL_API_URL")+"/users/me/drafts", utils.GET, "", []utils.Header{
				{
					Key:   "Authorization",
					Value: fmt.Sprintf("Bearer %s", poll.token),
				},
			}, "")
			if err != nil || len(drafts.Drafts) == 0 {
				log.WithField("error", err).Error("failed to get drafts")
				return err
			}
			if drafts.Drafts[0].ID != poll.LastDraft.ID {
				raw, err := formatNewDraftAction(drafts.Drafts[0], poll.respType)
				if err != nil {
					log.WithField("error", err).Error("failed to format new draft action")
					return err
				}
				_, err = coreClient.ForwardAction(context.Background(), &protogen.ForwardActionReq{
					Id:   poll.id,
					Type: poll.respType,
					Data: raw,
				})
				if err != nil {
					log.WithField("error", err).Error("failed to forward")
					continue
				}
				(*newDraft.watchers)[id].LastDraft = drafts.Drafts[0]
			}
		}
		newDraft.mut.Unlock()
	}
}

func eventNewEmailAtDateWithSender(newEmailAtDateWithSender NewEmailAtDateWithSenderAction, coreClient protogen.CoreServiceClient) error {
	for {
		time.Sleep(10 * time.Second)
		newEmailAtDateWithSender.mut.Lock()
		for id, poll := range *newEmailAtDateWithSender.watchers {
			senderOk := false
			emails, err := utils.GetFromJsonReq[Emails](os.Getenv("GMAIL_API_URL")+"/users/me/messages", utils.GET, "", []utils.Header{
				{
					Key:   "Authorization",
					Value: fmt.Sprintf("Bearer %s", poll.token),
				},
			}, "")
			if err != nil {
				log.WithField("error", err).Error("failed to get emails")
				return err
			}
			if emails.Messages[0].ID != poll.LastEmail.ID {
				emailContent, err := utils.GetFromJsonReq[EmailContent](os.Getenv("GMAIL_API_URL")+"/users/me/messages/"+emails.Messages[0].ID, utils.GET, "", []utils.Header{
					{
						Key:   "Authorization",
						Value: fmt.Sprintf("Bearer %s", poll.token),
					},
				}, "")
				if err != nil || len(emailContent.Payload.Headers) == 0 {
					log.WithField("error", err).Error("failed to get email content")
					return err
				}
				for _, header := range emailContent.Payload.Headers {
					if header.Name == "Date" {
						fmt.Println("date ok")
						dateHeaderValue := checkFormatDate(header.Value)
						date, err := time.Parse(time.RFC1123Z, dateHeaderValue)
						if err != nil {
							log.WithField("error", err).Error("failed to parse date")
							return err
						}
						dateFormat, err := time.Parse(time.DateOnly, date.Format(time.DateOnly))
						if err != nil {
							log.WithField("error", err).Error("failed to parse date")
							return err
						}
						if dateFormat.Equal(poll.date) {
							senderOk = true
							fmt.Println("after sender ok = ", senderOk)
						}
					}
					if header.Name == "From" && strings.Contains(header.Value, poll.sender) {
						fmt.Println("sender ok = ", senderOk)
						if !senderOk {
							break
						}
						raw, err := formatNewEmailAtDateWithSenderAction(poll.date, poll.sender, emailContent, poll.respType)
						if err != nil {
							log.WithField("error", err).Error("failed to format new email action")
							return err
						}
						_, err = coreClient.ForwardAction(context.Background(), &protogen.ForwardActionReq{
							Id:   poll.id,
							Type: poll.respType,
							Data: raw,
						})
						if err != nil {
							log.WithField("error", err).Error("failed to forward")
							continue
						}
					}
				}
				(*newEmailAtDateWithSender.watchers)[id].LastEmail = emails.Messages[0]
			}
		}
		newEmailAtDateWithSender.mut.Unlock()
	}
}

func insertCharAtIndex(inputString string, charToInsert string, index int) string {
	if index < 0 {
		index = 0
	} else if index > len(inputString) {
		index = len(inputString)
	}

	return inputString[:index] + charToInsert + inputString[index:]
}

func checkFormatDate(date string) string {
	if date[6] == ' ' {
		date = insertCharAtIndex(date, "0", 5)
		return date
	}
	return date
}

func eventNewEmailAtDate(newEmailAtDate NewEmailAtDateAction, coreClient protogen.CoreServiceClient) error {
	for {
		time.Sleep(10 * time.Second)
		newEmailAtDate.mut.Lock()
		for id, poll := range *newEmailAtDate.watchers {
			emails, err := utils.GetFromJsonReq[Emails](os.Getenv("GMAIL_API_URL")+"/users/me/messages", utils.GET, "", []utils.Header{
				{
					Key:   "Authorization",
					Value: fmt.Sprintf("Bearer %s", poll.token),
				},
			}, "")
			if err != nil {
				log.WithField("error", err).Error("failed to get emails")
				return err
			}
			if emails.Messages[0].ID != poll.LastEmail.ID {
				emailContent, err := utils.GetFromJsonReq[EmailContent](os.Getenv("GMAIL_API_URL")+"/users/me/messages/"+emails.Messages[0].ID, utils.GET, "", []utils.Header{
					{
						Key:   "Authorization",
						Value: fmt.Sprintf("Bearer %s", poll.token),
					},
				}, "")
				if err != nil || len(emailContent.Payload.Headers) == 0 {
					log.WithField("error", err).Error("failed to get email content")
					return err
				}
				for _, header := range emailContent.Payload.Headers {
					if header.Name == "Date" {
						dateHeaderValue := checkFormatDate(header.Value)
						date, err := time.Parse(time.RFC1123Z, dateHeaderValue)
						if err != nil {
							log.WithField("error", err).Error("failed to parse date")
							return err
						}
						dateFormat, err := time.Parse(time.DateOnly, date.Format(time.DateOnly))
						if err != nil {
							log.WithField("error", err).Error("failed to parse date")
							return err
						}
						if dateFormat.Equal(poll.date) {
							raw, err := formatNewEmailAtDateAction(poll.date, emailContent, poll.respType)
							if err != nil {
								log.WithField("error", err).Error("failed to format new email action")
								return err
							}
							_, err = coreClient.ForwardAction(context.Background(), &protogen.ForwardActionReq{
								Id:   poll.id,
								Type: poll.respType,
								Data: raw,
							})
							if err != nil {
								log.WithField("error", err).Error("failed to forward")
								continue
							}
						}
					}
				}
				(*newEmailAtDate.watchers)[id].LastEmail = emails.Messages[0]
			}
		}
		newEmailAtDate.mut.Unlock()
	}
}

func eventNewEmailWithSender(newEmailWithSender NewEmailWithSenderAction, coreClient protogen.CoreServiceClient) error {
	for {
		time.Sleep(10 * time.Second)
		newEmailWithSender.mut.Lock()
		for id, poll := range *newEmailWithSender.watchers {
			emails, err := utils.GetFromJsonReq[Emails](os.Getenv("GMAIL_API_URL")+"/users/me/messages", utils.GET, "", []utils.Header{
				{
					Key:   "Authorization",
					Value: fmt.Sprintf("Bearer %s", poll.token),
				},
			}, "")
			if err != nil {
				log.WithField("error", err).Error("failed to get emails")
				return err
			}
			if emails.Messages[0].ID != poll.LastEmail.ID {
				emailContent, err := utils.GetFromJsonReq[EmailContent](os.Getenv("GMAIL_API_URL")+"/users/me/messages/"+emails.Messages[0].ID, utils.GET, "", []utils.Header{
					{
						Key:   "Authorization",
						Value: fmt.Sprintf("Bearer %s", poll.token),
					},
				}, "")
				if err != nil || len(emailContent.Payload.Headers) == 0 {
					log.WithField("error", err).Error("failed to get email content")
					return err
				}
				for _, header := range emailContent.Payload.Headers {
					if header.Name == "From" && strings.Contains(header.Value, poll.sender) {
						raw, err := formatNewEmailWithSenderAction(poll.sender, emailContent, poll.respType)
						if err != nil {
							log.WithField("error", err).Error("failed to format new email action")
							return err
						}
						_, err = coreClient.ForwardAction(context.Background(), &protogen.ForwardActionReq{
							Id:   poll.id,
							Type: poll.respType,
							Data: raw,
						})
						if err != nil {
							log.WithField("error", err).Error("failed to forward")
							continue
						}
					}
				}
				(*newEmailWithSender.watchers)[id].LastEmail = emails.Messages[0]
			}
		}
		newEmailWithSender.mut.Unlock()
	}
}

func eventNewEmail(newEmail NewEmailAction, coreClient protogen.CoreServiceClient) error {
	for {
		time.Sleep(10 * time.Second)
		newEmail.mut.Lock()
		for id, poll := range *newEmail.watchersNewEmail {
			emails, err := utils.GetFromJsonReq[Emails](os.Getenv("GMAIL_API_URL")+"/users/me/messages", utils.GET, "", []utils.Header{
				{
					Key:   "Authorization",
					Value: fmt.Sprintf("Bearer %s", poll.token),
				},
			}, "")
			if err != nil {
				log.WithField("error", err).Error("failed to get emails")
				return err
			}
			fmt.Println("last email: ", poll.LastEmail.ID, "new email: ", emails.Messages[0].ID)
			if emails.Messages[0].ID != poll.LastEmail.ID {
				(*newEmail.watchersNewEmail)[id].LastEmail = emails.Messages[0]
				raw, err := formatNewEmailAction(emails.Messages[0].ID, poll.respType)
				if err != nil {
					log.WithField("error", err).Error("failed to format new email action")
					return err
				}
				_, err = coreClient.ForwardAction(context.Background(), &protogen.ForwardActionReq{
					Id:   poll.id,
					Type: poll.respType,
					Data: raw,
				})
				if err != nil {
					log.WithField("error", err).Error("failed to forward")
					continue
				}
			}
		}
		newEmail.mut.Unlock()
	}
}

type NewEmailAction struct {
	watchersNewEmail *[]pollsNewEmail
	mut              *sync.Mutex
}

type NewEmailWithSenderAction struct {
	watchers *[]pollsNewEmailWithSender
	mut      *sync.Mutex
}

type NewEmailAtDateAction struct {
	watchers *[]pollsNewEmailAtDate
	mut      *sync.Mutex
}

type NewEmailAtDateWithSenderAction struct {
	watchers *[]pollsNewEmailAtDateWithSender
	mut      *sync.Mutex
}

type NewDraftAction struct {
	watchers *[]pollsNewDraft
	mut      *sync.Mutex
}

type NewDraftAtDateAction struct {
	watchers *[]pollsNewDraftAtDate
	mut      *sync.Mutex
}

type NewDraftWithReceiverAction struct {
	watchers *[]pollsNewDraftWithReceiver
	mut      *sync.Mutex
}

type NewDraftAtDateWithReceiverAction struct {
	watchers *[]pollsNewDraftAtDateWithReceiver
	mut      *sync.Mutex
}

type NewLabelAction struct {
	watchers *[]pollsNewLabel
	mut      *sync.Mutex
}

type NewLabelWithNameAction struct {
	watchers *[]pollsNewLabelWithName
	mut      *sync.Mutex
}

type NewEmailInLabelAction struct {
	watchers *[]pollsNewEmailInLabel
	mut      *sync.Mutex
}

type GmailActionServer struct {
	NewEmail                   *NewEmailAction
	NewEmailWithSender         *NewEmailWithSenderAction
	NewEmailAtDate             *NewEmailAtDateAction
	NewEmailAtDateWithSender   *NewEmailAtDateWithSenderAction
	NewDraft                   *NewDraftAction
	NewDraftAtDate             *NewDraftAtDateAction
	NewDraftWithReceiver       *NewDraftWithReceiverAction
	NewDraftAtDateWithReceiver *NewDraftAtDateWithReceiverAction
	NewLabel                   *NewLabelAction
	NewLabelWithName           *NewLabelWithNameAction
	NewEmailInLabel            *NewEmailInLabelAction
	protogen.UnimplementedGmailServiceActionServer
}

func (e GmailActionServer) RegisterNewEmailInLabel(ctx context.Context, request *protogen.NewEmailInLabel_Request) (*protogen.Empty, error) {
	labels, err := utils.GetFromJsonReq[Labels](os.Getenv("GMAIL_API_URL")+"/users/me/labels", utils.GET, "", []utils.Header{
		{
			Key:   "Authorization",
			Value: fmt.Sprintf("Bearer %s", request.GetToken()),
		},
	}, "")
	if err != nil || len(labels.Labels) == 0 {
		return nil, fmt.Errorf("failed to get labels: %v", err)
	}
	for _, label := range labels.Labels {
		if label.Name == request.GetLabel() {
			labelContent, err := utils.GetFromJsonReq[LabelContent](os.Getenv("GMAIL_API_URL")+"/users/me/labels/"+label.ID, utils.GET, "", []utils.Header{
				{
					Key:   "Authorization",
					Value: fmt.Sprintf("Bearer %s", request.GetToken()),
				},
			}, "")
			if err != nil {
				return nil, fmt.Errorf("failed to get label content: %v", err)
			}
			*e.NewEmailInLabel.watchers = append(*e.NewEmailInLabel.watchers, pollsNewEmailInLabel{
				id:            request.GetId(),
				respType:      request.GetResponseType(),
				token:         request.GetToken(),
				label:         label.ID,
				MessagesTotal: labelContent.MessagesTotal,
			})
			return &protogen.Empty{}, nil
		}
	}
	return nil, fmt.Errorf("label not found")
}

func (e GmailActionServer) RegisterNewLabelWithName(ctx context.Context, request *protogen.NewLabelWithName_Request) (*protogen.Empty, error) {
	labels, err := utils.GetFromJsonReq[Labels](os.Getenv("GMAIL_API_URL")+"/users/me/labels", utils.GET, "", []utils.Header{
		{
			Key:   "Authorization",
			Value: fmt.Sprintf("Bearer %s", request.GetToken()),
		},
	}, "")
	if err != nil || len(labels.Labels) == 0 {
		return nil, fmt.Errorf("failed to get labels: %v", err)
	}
	*e.NewLabelWithName.watchers = append(*e.NewLabelWithName.watchers, pollsNewLabelWithName{
		id:       request.GetId(),
		respType: request.GetResponseType(),
		token:    request.GetToken(),
		name:     request.GetName(),
		Labels:   labels,
	})
	return &protogen.Empty{}, nil
}

func (e GmailActionServer) RegisterNewLabel(ctx context.Context, request *protogen.NewLabel_Request) (*protogen.Empty, error) {
	labels, err := utils.GetFromJsonReq[Labels](os.Getenv("GMAIL_API_URL")+"/users/me/labels", utils.GET, "", []utils.Header{
		{
			Key:   "Authorization",
			Value: fmt.Sprintf("Bearer %s", request.GetToken()),
		},
	}, "")
	if err != nil || len(labels.Labels) == 0 {
		return nil, fmt.Errorf("failed to get labels: %v", err)
	}
	*e.NewLabel.watchers = append(*e.NewLabel.watchers, pollsNewLabel{
		id:       request.GetId(),
		respType: request.GetResponseType(),
		token:    request.GetToken(),
		Labels:   labels,
	})
	return &protogen.Empty{}, nil
}

func (e GmailActionServer) RegisterNewDraftAtDateWithReceiver(ctx context.Context, request *protogen.NewDraftAtDateWithReceiver_Request) (*protogen.Empty, error) {
	drafts, err := utils.GetFromJsonReq[Drafts](os.Getenv("GMAIL_API_URL")+"/users/me/drafts", utils.GET, "", []utils.Header{
		{
			Key:   "Authorization",
			Value: fmt.Sprintf("Bearer %s", request.GetToken()),
		},
	}, "")
	if err != nil || len(drafts.Drafts) == 0 {
		return nil, fmt.Errorf("failed to get drafts: %v", err)
	}
	dateRequest, err := time.Parse(time.DateOnly, request.GetDate())
	if err != nil {
		return nil, fmt.Errorf("failed to parse date: %v", err)
	}
	*e.NewDraftAtDateWithReceiver.watchers = append(*e.NewDraftAtDateWithReceiver.watchers, pollsNewDraftAtDateWithReceiver{
		id:        request.GetId(),
		respType:  request.GetResponseType(),
		token:     request.GetToken(),
		receiver:  request.GetReceiver(),
		date:      dateRequest,
		LastDraft: drafts.Drafts[0],
	})
	return &protogen.Empty{}, nil
}

func (e GmailActionServer) RegisterNewDraftWithReceiver(ctx context.Context, request *protogen.NewDraftWithReceiver_Request) (*protogen.Empty, error) {
	drafts, err := utils.GetFromJsonReq[Drafts](os.Getenv("GMAIL_API_URL")+"/users/me/drafts", utils.GET, "", []utils.Header{
		{
			Key:   "Authorization",
			Value: fmt.Sprintf("Bearer %s", request.GetToken()),
		},
	}, "")
	if err != nil || len(drafts.Drafts) == 0 {
		return nil, fmt.Errorf("failed to get drafts: %v", err)
	}
	*e.NewDraftWithReceiver.watchers = append(*e.NewDraftWithReceiver.watchers, pollsNewDraftWithReceiver{
		id:        request.GetId(),
		respType:  request.GetResponseType(),
		token:     request.GetToken(),
		receiver:  request.GetReceiver(),
		LastDraft: drafts.Drafts[0],
	})
	return &protogen.Empty{}, nil
}

func (e GmailActionServer) RegisterNewDraftAtDate(ctx context.Context, request *protogen.NewDraftAtDate_Request) (*protogen.Empty, error) {
	drafts, err := utils.GetFromJsonReq[Drafts](os.Getenv("GMAIL_API_URL")+"/users/me/drafts", utils.GET, "", []utils.Header{
		{
			Key:   "Authorization",
			Value: fmt.Sprintf("Bearer %s", request.GetToken()),
		},
	}, "")
	if err != nil || len(drafts.Drafts) == 0 {
		return nil, fmt.Errorf("failed to get drafts: %v", err)
	}
	dateRequest, err := time.Parse(time.DateOnly, request.GetDate())
	if err != nil {
		return nil, fmt.Errorf("failed to parse date: %v", err)
	}
	*e.NewDraftAtDate.watchers = append(*e.NewDraftAtDate.watchers, pollsNewDraftAtDate{
		id:        request.GetId(),
		respType:  request.GetResponseType(),
		token:     request.GetToken(),
		date:      dateRequest,
		LastDraft: drafts.Drafts[0],
	})
	return &protogen.Empty{}, nil
}

func (e GmailActionServer) RegisterNewDraft(ctx context.Context, request *protogen.NewDraft_Request) (*protogen.Empty, error) {
	drafts, err := utils.GetFromJsonReq[Drafts](os.Getenv("GMAIL_API_URL")+"/users/me/drafts", utils.GET, "", []utils.Header{
		{
			Key:   "Authorization",
			Value: fmt.Sprintf("Bearer %s", request.GetToken()),
		},
	}, "")
	if err != nil || len(drafts.Drafts) == 0 {
		return nil, fmt.Errorf("failed to get drafts: %v", err)
	}
	*e.NewDraft.watchers = append(*e.NewDraft.watchers, pollsNewDraft{
		id:        request.GetId(),
		respType:  request.GetResponseType(),
		token:     request.GetToken(),
		LastDraft: drafts.Drafts[0],
	})
	return &protogen.Empty{}, nil
}

func (e GmailActionServer) RegisterNewEmailAtDateWithSender(ctx context.Context, request *protogen.NewEmailAtDateWithSender_Request) (*protogen.Empty, error) {
	emails, err := utils.GetFromJsonReq[Emails](os.Getenv("GMAIL_API_URL")+"/users/me/messages", utils.GET, "", []utils.Header{
		{
			Key:   "Authorization",
			Value: fmt.Sprintf("Bearer %s", request.GetToken()),
		},
	}, "")
	if err != nil || len(emails.Messages) == 0 {
		return nil, fmt.Errorf("failed to get emails: %v", err)
	}
	fmt.Println("getDate() = ", request.GetDate())
	dateRequest, err := time.Parse(time.DateOnly, request.GetDate())
	if err != nil {
		return nil, fmt.Errorf("failed to parse date: %v", err)
	}
	*e.NewEmailAtDateWithSender.watchers = append(*e.NewEmailAtDateWithSender.watchers, pollsNewEmailAtDateWithSender{
		id:        request.GetId(),
		respType:  request.GetResponseType(),
		token:     request.GetToken(),
		sender:    request.GetSender(),
		date:      dateRequest,
		LastEmail: emails.Messages[0],
	})
	return &protogen.Empty{}, nil
}

func (e GmailActionServer) RegisterNewEmailAtDate(ctx context.Context, request *protogen.NewEmailAtDate_Request) (*protogen.Empty, error) {
	emails, err := utils.GetFromJsonReq[Emails](os.Getenv("GMAIL_API_URL")+"/users/me/messages", utils.GET, "", []utils.Header{
		{
			Key:   "Authorization",
			Value: fmt.Sprintf("Bearer %s", request.GetToken()),
		},
	}, "")
	if err != nil || len(emails.Messages) == 0 {
		return nil, fmt.Errorf("failed to get emails: %v", err)
	}
	dateRequest, err := time.Parse(time.DateOnly, request.GetDate())
	if err != nil {
		return nil, fmt.Errorf("failed to parse date: %v", err)
	}
	*e.NewEmailAtDate.watchers = append(*e.NewEmailAtDate.watchers, pollsNewEmailAtDate{
		id:        request.GetId(),
		respType:  request.GetResponseType(),
		token:     request.GetToken(),
		date:      dateRequest,
		LastEmail: emails.Messages[0],
	})
	return &protogen.Empty{}, nil
}

func (e GmailActionServer) RegisterNewEmail(ctx context.Context, request *protogen.NewEmail_Request) (*protogen.Empty, error) {
	emails, err := utils.GetFromJsonReq[Emails](os.Getenv("GMAIL_API_URL")+"/users/me/messages", utils.GET, "", []utils.Header{
		{
			Key:   "Authorization",
			Value: fmt.Sprintf("Bearer %s", request.GetToken()),
		},
	}, "")
	if err != nil || len(emails.Messages) == 0 {
		return nil, fmt.Errorf("failed to get emails: %v", err)
	}
	*e.NewEmail.watchersNewEmail = append(*e.NewEmail.watchersNewEmail, pollsNewEmail{
		id:        request.GetId(),
		respType:  request.GetResponseType(),
		token:     request.GetToken(),
		LastEmail: emails.Messages[0],
	})
	return &protogen.Empty{}, nil
}

func (e GmailActionServer) RegisterNewEmailWithSender(ctx context.Context, request *protogen.NewEmailWithSender_Request) (*protogen.Empty, error) {
	emails, err := utils.GetFromJsonReq[Emails](os.Getenv("GMAIL_API_URL")+"/users/me/messages", utils.GET, "", []utils.Header{
		{
			Key:   "Authorization",
			Value: fmt.Sprintf("Bearer %s", request.GetToken()),
		},
	}, "")
	if err != nil || len(emails.Messages) == 0 {
		return nil, fmt.Errorf("failed to get emails: %v", err)
	}
	*e.NewEmailWithSender.watchers = append(*e.NewEmailWithSender.watchers, pollsNewEmailWithSender{
		id:        request.GetId(),
		respType:  request.GetResponseType(),
		token:     request.GetToken(),
		sender:    request.GetSender(),
		LastEmail: emails.Messages[0],
	})
	return &protogen.Empty{}, nil
}

func NewGmailActionServer(newEmail *NewEmailAction, newEmailWithSender *NewEmailWithSenderAction, newEmailAtDate *NewEmailAtDateAction, newEmailAtDateWithSender *NewEmailAtDateWithSenderAction, newDraft *NewDraftAction,
	newDraftAtDate *NewDraftAtDateAction, newDraftWithReceiver *NewDraftWithReceiverAction, newDraftAtDateWithReceiver *NewDraftAtDateWithReceiverAction,
	newLabel *NewLabelAction, newLabelWithName *NewLabelWithNameAction, newEmailInLabel *NewEmailInLabelAction) (protogen.GmailServiceActionServer, error) {
	return &GmailActionServer{
		NewEmail:                   newEmail,
		NewEmailWithSender:         newEmailWithSender,
		NewEmailAtDate:             newEmailAtDate,
		NewEmailAtDateWithSender:   newEmailAtDateWithSender,
		NewDraft:                   newDraft,
		NewDraftAtDate:             newDraftAtDate,
		NewDraftWithReceiver:       newDraftWithReceiver,
		NewDraftAtDateWithReceiver: newDraftAtDateWithReceiver,
		NewLabel:                   newLabel,
		NewLabelWithName:           newLabelWithName,
		NewEmailInLabel:            newEmailInLabel,
	}, nil
}

func formatNewEmailInLabel(labelContent LabelContent, kind protogen.Format_Kind) ([]byte, error) {
	switch kind {
	case protogen.Format_GHIncidentReportKind:
		p := protogen.Format_GHIncidentReport{
			Title:   fmt.Sprintf("New email detected in label %s", labelContent.ID),
			Content: fmt.Sprintf("Now the number of email in label is %d", labelContent.MessagesTotal),
		}
		return proto.Marshal(&p)
	case protogen.Format_OnlyTitleKind:
		p := protogen.Format_OnlyTitle{
			Title: fmt.Sprintf("New email detected in label %s", labelContent.ID),
		}
		return proto.Marshal(&p)
	case protogen.Format_ManyFilesKind:
		var files []*protogen.Format_ManyFiles_File
		files = append(files, &protogen.Format_ManyFiles_File{
			Name:    "email",
			Content: fmt.Sprintf("%x", labelContent),
		})
		return proto.Marshal(&protogen.Format_ManyFiles{
			Files: files,
		})
	case protogen.Format_MessageAndDestinaryKind:
		p := protogen.Format_MessageAndDestinary{
			Destinary: "email",
			Message:   fmt.Sprintf("%x", labelContent),
		}
		return proto.Marshal(&p)
	case protogen.Format_TagsKind:
		p := protogen.Format_Tags{
			Tags: []string{
				"email",
			},
		}
		return proto.Marshal(&p)
	case protogen.Format_GHNewBranchKing:
		p := protogen.Format_GHNewBranch{
			Name: labelContent.ID,
		}
		return proto.Marshal(&p)
	case protogen.Format_NoParamKind:
		p := protogen.Format_NoParam{}
		return proto.Marshal(&p)
	}

	return nil, fmt.Errorf("not implemented")
}

func formatNewLabelWithNameAction(label Label, kind protogen.Format_Kind) ([]byte, error) {
	switch kind {
	case protogen.Format_GHIncidentReportKind:
		p := protogen.Format_GHIncidentReport{
			Title:   fmt.Sprintf("New label detected with name %s", label.Name),
			Content: fmt.Sprintf("%x", label),
		}
		return proto.Marshal(&p)
	case protogen.Format_TagsKind:
		p := protogen.Format_Tags{
			Tags: []string{
				label.Name,
			},
		}
		return proto.Marshal(&p)
	case protogen.Format_NoParamKind:
		p := protogen.Format_NoParam{}
		return proto.Marshal(&p)
	case protogen.Format_ManyFilesKind:
		var files []*protogen.Format_ManyFiles_File
		files = append(files, &protogen.Format_ManyFiles_File{
			Name: "label",
		})
		return proto.Marshal(&protogen.Format_ManyFiles{
			Files: files,
		})
	case protogen.Format_OnlyTitleKind:
		p := protogen.Format_OnlyTitle{
			Title: fmt.Sprintf("New label detected with name %s", label.Name),
		}
		return proto.Marshal(&p)
	case protogen.Format_GHNewBranchKing:
		p := protogen.Format_GHNewBranch{
			Name: label.Name,
		}
		return proto.Marshal(&p)
	case protogen.Format_MessageAndDestinaryKind:
		p := protogen.Format_MessageAndDestinary{
			Destinary: "label",
			Message:   fmt.Sprintf("%x", label),
		}
		return proto.Marshal(&p)
	}

	return nil, fmt.Errorf("not implemented")
}

func formatNewLabelAction(labels Labels, kind protogen.Format_Kind) ([]byte, error) {
	switch kind {
	case protogen.Format_MessageAndDestinaryKind:
		p := protogen.Format_MessageAndDestinary{
			Destinary: "label",
			Message:   fmt.Sprintf("%x", labels),
		}
		return proto.Marshal(&p)
	case protogen.Format_GHIncidentReportKind:
		p := protogen.Format_GHIncidentReport{
			Title:   "New label detected",
			Content: fmt.Sprintf("Now the number of label is %d", len(labels.Labels)),
		}
		return proto.Marshal(&p)
	case protogen.Format_OnlyTitleKind:
		p := protogen.Format_OnlyTitle{
			Title: "New label detected",
		}
		return proto.Marshal(&p)
	case protogen.Format_ManyFilesKind:
		var files []*protogen.Format_ManyFiles_File
		files = append(files, &protogen.Format_ManyFiles_File{
			Name:    "labels",
			Content: fmt.Sprintf("%x", labels),
		})
		return proto.Marshal(&protogen.Format_ManyFiles{
			Files: files,
		})
	case protogen.Format_TagsKind:
		var tags []string
		for _, label := range labels.Labels {
			tags = append(tags, label.Name)
		}
		p := protogen.Format_Tags{
			Tags: tags,
		}
		return proto.Marshal(&p)
	case protogen.Format_GHNewBranchKing:
		var names []string
		for _, label := range labels.Labels {
			names = append(names, label.Name)
		}
		p := protogen.Format_GHNewBranch{
			Name: strings.Join(names, ","),
		}
		return proto.Marshal(&p)
	case protogen.Format_NoParamKind:
		p := protogen.Format_NoParam{}
		return proto.Marshal(&p)
	}

	return nil, fmt.Errorf("not implemented")
}

func formatNewDraftAtDateWithReceiverAction(date time.Time, receiver string, draft Draft, kind protogen.Format_Kind) ([]byte, error) {
	switch kind {
	case protogen.Format_MessageAndDestinaryKind:
		p := protogen.Format_MessageAndDestinary{
			Destinary: "draft",
			Message:   fmt.Sprintf("%x", draft),
		}
		return proto.Marshal(&p)
	case protogen.Format_GHIncidentReportKind:
		p := protogen.Format_GHIncidentReport{
			Title:   fmt.Sprintf("New draft detected at date %s for %s ID: %s", date, receiver, draft.ID),
			Content: fmt.Sprintf("%x", draft),
		}
		return proto.Marshal(&p)
	case protogen.Format_OnlyTitleKind:
		p := protogen.Format_OnlyTitle{
			Title: fmt.Sprintf("New draft detected at date %s for %s ID: %s", date, receiver, draft.ID),
		}
		return proto.Marshal(&p)
	case protogen.Format_ManyFilesKind:
		var files []*protogen.Format_ManyFiles_File
		files = append(files, &protogen.Format_ManyFiles_File{
			Name:    "draft",
			Content: fmt.Sprintf("%x", draft),
		})
		return proto.Marshal(&protogen.Format_ManyFiles{
			Files: files,
		})
	case protogen.Format_TagsKind:
		p := protogen.Format_Tags{
			Tags: []string{
				"draft",
			},
		}
		return proto.Marshal(&p)
	case protogen.Format_GHNewBranchKing:
		p := protogen.Format_GHNewBranch{
			Name: draft.ID,
		}
		return proto.Marshal(&p)
	case protogen.Format_NoParamKind:
		p := protogen.Format_NoParam{}
		return proto.Marshal(&p)
	}
	return nil, fmt.Errorf("not implemented")
}

func formatNewDraftWithReceiverAction(receiver string, draft Draft, kind protogen.Format_Kind) ([]byte, error) {
	switch kind {
	case protogen.Format_MessageAndDestinaryKind:
		p := protogen.Format_MessageAndDestinary{
			Destinary: "draft",
			Message:   fmt.Sprintf("%x", draft),
		}
		return proto.Marshal(&p)
	case protogen.Format_GHIncidentReportKind:
		p := protogen.Format_GHIncidentReport{
			Title:   fmt.Sprintf("New draft detected for %s ID: %s", receiver, draft.ID),
			Content: fmt.Sprintf("%x", draft),
		}
		return proto.Marshal(&p)
	case protogen.Format_OnlyTitleKind:
		p := protogen.Format_OnlyTitle{
			Title: fmt.Sprintf("New draft detected for %s ID: %s", receiver, draft.ID),
		}
		return proto.Marshal(&p)
	case protogen.Format_ManyFilesKind:
		var files []*protogen.Format_ManyFiles_File
		files = append(files, &protogen.Format_ManyFiles_File{
			Name: "draft",
		})
		return proto.Marshal(&protogen.Format_ManyFiles{
			Files: files,
		})
	case protogen.Format_TagsKind:
		p := protogen.Format_Tags{
			Tags: []string{
				"draft",
			},
		}
		return proto.Marshal(&p)
	case protogen.Format_GHNewBranchKing:
		p := protogen.Format_GHNewBranch{
			Name: draft.ID,
		}
		return proto.Marshal(&p)
	case protogen.Format_NoParamKind:
		p := protogen.Format_NoParam{}
		return proto.Marshal(&p)
	}
	return nil, fmt.Errorf("not implemented")
}

func formatNewDraftAtDateAction(date time.Time, draft Draft, kind protogen.Format_Kind) ([]byte, error) {
	switch kind {
	case protogen.Format_MessageAndDestinaryKind:
		p := protogen.Format_MessageAndDestinary{
			Destinary: "draft",
		}
		return proto.Marshal(&p)
	case protogen.Format_GHIncidentReportKind:
		p := protogen.Format_GHIncidentReport{
			Title:   fmt.Sprintf("New draft detected at date %s ID: %s", date, draft.ID),
			Content: fmt.Sprintf("%x", draft),
		}
		return proto.Marshal(&p)
	case protogen.Format_OnlyTitleKind:
		p := protogen.Format_OnlyTitle{
			Title: fmt.Sprintf("New draft detected at date %s ID: %s", date, draft.ID),
		}
		return proto.Marshal(&p)
	case protogen.Format_ManyFilesKind:
		var files []*protogen.Format_ManyFiles_File
		files = append(files, &protogen.Format_ManyFiles_File{
			Name: "draft",
		})
		return proto.Marshal(&protogen.Format_ManyFiles{
			Files: files,
		})
	case protogen.Format_TagsKind:
		p := protogen.Format_Tags{
			Tags: []string{
				"draft",
			},
		}
		return proto.Marshal(&p)
	case protogen.Format_GHNewBranchKing:
		p := protogen.Format_GHNewBranch{
			Name: draft.ID,
		}
		return proto.Marshal(&p)
	case protogen.Format_NoParamKind:
		p := protogen.Format_NoParam{}
		return proto.Marshal(&p)
	}
	return nil, fmt.Errorf("not implemented")
}

func formatNewDraftAction(draft Draft, kind protogen.Format_Kind) ([]byte, error) {
	switch kind {
	case protogen.Format_MessageAndDestinaryKind:
		p := protogen.Format_MessageAndDestinary{
			Destinary: "draft",
			Message:   fmt.Sprintf("%x", draft),
		}
		return proto.Marshal(&p)
	case protogen.Format_GHIncidentReportKind:
		p := protogen.Format_GHIncidentReport{
			Title:   fmt.Sprintf("New draft detected ID: %s", draft.ID),
			Content: fmt.Sprintf("%x", draft),
		}
		return proto.Marshal(&p)
	case protogen.Format_OnlyTitleKind:
		p := protogen.Format_OnlyTitle{
			Title: fmt.Sprintf("New draft detected ID: %s", draft.ID),
		}
		return proto.Marshal(&p)
	case protogen.Format_ManyFilesKind:
		var files []*protogen.Format_ManyFiles_File
		files = append(files, &protogen.Format_ManyFiles_File{
			Name:    "draft",
			Content: fmt.Sprintf("%x", draft),
		})
		return proto.Marshal(&protogen.Format_ManyFiles{
			Files: files,
		})
	case protogen.Format_TagsKind:
		p := protogen.Format_Tags{
			Tags: []string{
				"draft",
			},
		}
		return proto.Marshal(&p)
	case protogen.Format_GHNewBranchKing:
		p := protogen.Format_GHNewBranch{
			Name: draft.ID,
		}
		return proto.Marshal(&p)
	case protogen.Format_NoParamKind:
		p := protogen.Format_NoParam{}
		return proto.Marshal(&p)
	}
	return nil, fmt.Errorf("not implemented")
}

func formatNewEmailAtDateWithSenderAction(date time.Time, sender string, content EmailContent, kind protogen.Format_Kind) ([]byte, error) {
	switch kind {
	case protogen.Format_MessageAndDestinaryKind:
		p := protogen.Format_MessageAndDestinary{
			Destinary: "email",
		}
		return proto.Marshal(&p)
	case protogen.Format_GHIncidentReportKind:
		p := protogen.Format_GHIncidentReport{
			Title:   fmt.Sprintf("New email detected at date %s from %s", date, sender),
			Content: fmt.Sprintf("%s", content.Snippet),
		}
		return proto.Marshal(&p)
	case protogen.Format_OnlyTitleKind:
		p := protogen.Format_OnlyTitle{
			Title: fmt.Sprintf("New email detected at date %s from %s", date, sender),
		}
		return proto.Marshal(&p)
	case protogen.Format_ManyFilesKind:
		var files []*protogen.Format_ManyFiles_File
		files = append(files, &protogen.Format_ManyFiles_File{
			Name:    "email",
			Content: fmt.Sprintf("%s", content.Snippet),
		})
		return proto.Marshal(&protogen.Format_ManyFiles{
			Files: files,
		})
	case protogen.Format_TagsKind:
		p := protogen.Format_Tags{
			Tags: []string{
				"email",
			},
		}
		return proto.Marshal(&p)
	case protogen.Format_GHNewBranchKing:
		p := protogen.Format_GHNewBranch{
			Name: content.ID,
		}
		return proto.Marshal(&p)
	case protogen.Format_NoParamKind:
		p := protogen.Format_NoParam{}
		return proto.Marshal(&p)
	}

	return nil, fmt.Errorf("not implemented")
}

func formatNewEmailAtDateAction(date time.Time, content EmailContent, kind protogen.Format_Kind) ([]byte, error) {
	switch kind {
	case protogen.Format_MessageAndDestinaryKind:
		p := protogen.Format_MessageAndDestinary{
			Destinary: "email",
		}
		return proto.Marshal(&p)
	case protogen.Format_GHIncidentReportKind:
		p := protogen.Format_GHIncidentReport{
			Title:   fmt.Sprintf("New email detected at date %s", date),
			Content: fmt.Sprintf("%s", content.Snippet),
		}
		return proto.Marshal(&p)
	case protogen.Format_OnlyTitleKind:
		p := protogen.Format_OnlyTitle{
			Title: fmt.Sprintf("New email detected at date %s", date),
		}
		return proto.Marshal(&p)
	case protogen.Format_ManyFilesKind:
		var files []*protogen.Format_ManyFiles_File
		files = append(files, &protogen.Format_ManyFiles_File{
			Name:    "email",
			Content: fmt.Sprintf("%s", content.Snippet),
		})
		return proto.Marshal(&protogen.Format_ManyFiles{
			Files: files,
		})
	case protogen.Format_TagsKind:
		p := protogen.Format_Tags{
			Tags: []string{
				"email",
			},
		}
		return proto.Marshal(&p)
	case protogen.Format_GHNewBranchKing:
		p := protogen.Format_GHNewBranch{
			Name: content.ID,
		}
		return proto.Marshal(&p)
	case protogen.Format_NoParamKind:
		p := protogen.Format_NoParam{}
		return proto.Marshal(&p)
	}
	return nil, fmt.Errorf("not implemented")
}

func formatNewEmailWithSenderAction(sender string, content EmailContent, kind protogen.Format_Kind) ([]byte, error) {
	switch kind {
	case protogen.Format_MessageAndDestinaryKind:
		p := protogen.Format_MessageAndDestinary{
			Destinary: "email",
		}
		return proto.Marshal(&p)
	case protogen.Format_GHIncidentReportKind:
		p := protogen.Format_GHIncidentReport{
			Title:   fmt.Sprintf("New email detected from %s", sender),
			Content: fmt.Sprintf("%s", content.Snippet),
		}
		return proto.Marshal(&p)
	case protogen.Format_OnlyTitleKind:
		p := protogen.Format_OnlyTitle{
			Title: fmt.Sprintf("New email detected from %s", sender),
		}
		return proto.Marshal(&p)
	case protogen.Format_ManyFilesKind:
		var files []*protogen.Format_ManyFiles_File
		files = append(files, &protogen.Format_ManyFiles_File{
			Name:    "email",
			Content: fmt.Sprintf("%s", content.Snippet),
		})
		return proto.Marshal(&protogen.Format_ManyFiles{
			Files: files,
		})
	case protogen.Format_TagsKind:
		p := protogen.Format_Tags{
			Tags: []string{
				"email",
			},
		}
		return proto.Marshal(&p)
	case protogen.Format_GHNewBranchKing:
		p := protogen.Format_GHNewBranch{
			Name: content.ID,
		}
		return proto.Marshal(&p)
	case protogen.Format_NoParamKind:
		p := protogen.Format_NoParam{}
		return proto.Marshal(&p)
	}
	return nil, fmt.Errorf("not implemented")
}

func formatNewEmailAction(email string, kind protogen.Format_Kind) ([]byte, error) {
	switch kind {
	case protogen.Format_MessageAndDestinaryKind:
		p := protogen.Format_MessageAndDestinary{
			Destinary: "email",
		}
		return proto.Marshal(&p)
	case protogen.Format_GHIncidentReportKind:
		p := protogen.Format_GHIncidentReport{
			Title:   fmt.Sprintf("New email detected ID: %s", email),
			Content: fmt.Sprintf("%x", email),
		}
		return proto.Marshal(&p)
	case protogen.Format_OnlyTitleKind:
		p := protogen.Format_OnlyTitle{
			Title: fmt.Sprintf("New email detected ID: %s", email),
		}
		return proto.Marshal(&p)
	case protogen.Format_ManyFilesKind:
		var files []*protogen.Format_ManyFiles_File
		files = append(files, &protogen.Format_ManyFiles_File{
			Name:    "email",
			Content: fmt.Sprintf("%x", email),
		})
		return proto.Marshal(&protogen.Format_ManyFiles{
			Files: files,
		})
	case protogen.Format_TagsKind:
		p := protogen.Format_Tags{
			Tags: []string{
				"email",
			},
		}
		return proto.Marshal(&p)
	case protogen.Format_GHNewBranchKing:
		p := protogen.Format_GHNewBranch{
			Name: email,
		}
		return proto.Marshal(&p)
	case protogen.Format_NoParamKind:
		p := protogen.Format_NoParam{}
		return proto.Marshal(&p)
	}

	return nil, fmt.Errorf("not implemented")
}
