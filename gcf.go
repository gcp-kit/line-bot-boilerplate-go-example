package gcf

// nolint
import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/gcp-kit/line-bot-boilerplate-go-example/function"
	"github.com/gcp-kit/line-bot-boilerplate-go/core"
	"github.com/gcp-kit/line-bot-boilerplate-go/util"
	"github.com/line/line-bot-sdk-go/linebot"
)

// example
var (
	count int
)

// nolint
func init() {
	if projectID != "" {
		entryPoint = os.Getenv("FUNCTION_NAME")
		setting("parent-test", "child-test")
	}
}

// setFunction add a function to use in `ChildFunctions`
// editing required
func setFunction() {
	tracer.Function = map[core.TracerName]func(context.Context, *core.Operation, *linebot.Event) *core.TracerResp{
		core.TracerFollowEvent:     function.FollowEvent,
		core.TracerUnfollowEvent:   function.UnfollowEvent,
		core.TracerTextMessage:     function.TextEvent,
		core.TracerLocationMessage: function.LocationEvent,
	}
	/*
	** the processing to put in the Global variable is here
	 */
	count++ // example
}

// WebHook CloudFunctions(Trigger: HTTP)
// no edit
// nolint
func WebHook(w http.ResponseWriter, r *http.Request) {
	ctx := util.SetGinContext(w, r)
	core.WebHook(ctx, secret, parentTopic)
}

// ParentFunctions CloudFunctions(Trigger: Pub/Sub)
// no edit
// nolint
func Forking(ctx context.Context, m *pubsub.Message) error {
	log.Println("EntryPoint:", entryPoint)
	switch entryPoint {
	case RouteParentFunctions:
		return core.ParentFunctions(ctx, m, tracer, childTopic)
	case RouteChildFunctions:
		return core.ChildFunctions(ctx, m, operation)
	default:
		return fmt.Errorf("invalid function name")
	}
}

// LiffFull CloudFunctions(Trigger: HTTP)
// no edit
// nolint
func LiffFull(w http.ResponseWriter, r *http.Request) {
	ctx := util.SetGinContext(w, r)
	function.Liff(ctx)
}

// LiffTall CloudFunctions(Trigger: HTTP)
// no edit
// nolint
func LiffTall(w http.ResponseWriter, r *http.Request) {
	ctx := util.SetGinContext(w, r)
	function.Liff(ctx)
}

// LiffCompact CloudFunctions(Trigger: HTTP)
// no edit
// nolint
func LiffCompact(w http.ResponseWriter, r *http.Request) {
	ctx := util.SetGinContext(w, r)
	function.Liff(ctx)
}
