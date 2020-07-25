package function

import (
	"context"

	"github.com/gcp-kit/line-bot-boilerplate-go/core"
	"github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/xerrors"
)

// TextEvent - handle text message events
func TextEvent(_ context.Context, op *core.Operation, event *linebot.Event) *core.TracerResp {
	resp := new(core.TracerResp)
	message, ok := event.Message.(*linebot.TextMessage)
	if !ok {
		resp.Error = xerrors.New("couldn't cast")
		return resp
	}
	switch message.Text {
	case "ping":
		items := &linebot.QuickReplyItems{
			Items: []*linebot.QuickReplyButton{
				{
					Action: linebot.QuickReplyAction(&linebot.MessageAction{
						Label: "ping",
						Text:  "ping",
					}),
				},
			},
		}
		prof, err := op.GetProfile(event.Source.UserID).Do()
		if err != nil {
			resp.Error = err
			return resp
		}
		sender := &linebot.Sender{
			Name:    prof.DisplayName,
			IconURL: prof.PictureURL,
		}
		msg := linebot.NewTextMessage("pong").WithQuickReplies(items).WithSender(sender)
		resp.Stack = append(resp.Stack, msg)
	default:
		msg := linebot.NewTextMessage(message.Text)
		resp.Stack = append(resp.Stack, msg)
	}
	return resp
}
