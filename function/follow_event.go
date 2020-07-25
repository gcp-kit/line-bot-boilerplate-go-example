package function

import (
	"context"
	"fmt"
	"log"

	"github.com/gcp-kit/line-bot-boilerplate-go/core"
	"github.com/gcp-kit/line-bot-boilerplate-go/util"
	"github.com/line/line-bot-sdk-go/linebot"
)

// FollowEvent - handle follow events
func FollowEvent(_ context.Context, op *core.Operation, event *linebot.Event) *core.TracerResp {
	resp := new(core.TracerResp)
	name := util.GetCallFuncName()
	log.Println("Call:", name)

	uid := event.Source.UserID
	log.Println("UID:", uid)

	prof, err := op.GetProfile(uid).Do()
	if err == nil {
		log.Println("Name:", prof.DisplayName)
		log.Println("Picture:", prof.PictureURL)
		text := fmt.Sprintf("%s, thanks for following!", prof.DisplayName)
		resp.Stack = append(resp.Stack, linebot.NewTextMessage(text))
	} else {
		log.Println("Error:", err.Error())
	}
	return resp
}
