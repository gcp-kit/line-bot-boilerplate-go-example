package function

import (
	"context"
	"log"

	"github.com/gcp-kit/line-bot-boilerplate-go/core"
	"github.com/gcp-kit/line-bot-boilerplate-go/util"
	"github.com/line/line-bot-sdk-go/linebot"
)

// UnfollowEvent - handle unfollow events
func UnfollowEvent(_ context.Context, _ *core.Operation, event *linebot.Event) *core.TracerResp {
	resp := new(core.TracerResp)
	name := util.GetCallFuncName()
	log.Println("Call:", name)

	uid := event.Source.UserID
	log.Println("UID:", uid)

	return resp
}
