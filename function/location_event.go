package function

import (
	"context"
	"fmt"

	"github.com/gcp-kit/line-bot-boilerplate-go/core"
	"github.com/line/line-bot-sdk-go/linebot"
	"golang.org/x/xerrors"
)

// LocationEvent - handle location message events
func LocationEvent(_ context.Context, _ *core.Operation, event *linebot.Event) *core.TracerResp {
	resp := new(core.TracerResp)
	message, ok := event.Message.(*linebot.LocationMessage)
	if !ok {
		resp.Error = xerrors.New("couldn't cast")
		return resp
	}

	text := fmt.Sprintf("Latitude: %f\nLongitude: %f", message.Latitude, message.Longitude)
	resp.Stack = append(resp.Stack, linebot.NewTextMessage(text))
	return resp
}
