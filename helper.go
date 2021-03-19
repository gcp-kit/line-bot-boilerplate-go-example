// Package function no edit file
package gcf

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/gcp-kit/line-bot-boilerplate-go-example/pkg/utils"
	"github.com/gcp-kit/line-bot-boilerplate-go/core"
	"github.com/line/line-bot-sdk-go/linebot"
)

const (
	RouteWebHook         = "WebHook"
	RouteParentFunctions = "ParentFunctions"
	RouteChildFunctions  = "ChildFunctions"
	RouteLiffFull        = "LiffFull"
	RouteLiffTall        = "LiffTall"
	RouteLiffCompact     = "LiffCompact"
)

// no edit
var (
	secret       string
	entryPoint   string
	operation    *core.Operation
	tracer       *core.Tracer
	pubSubClient *pubsub.Client
	parentTopic  *pubsub.Topic
	childTopic   *pubsub.Topic
)

// Probably no edit
func setting(parentTopicName, childTopicName string) {
	ctx := context.Background()
	tracer = new(core.Tracer)

	var err error
	pubSubClient, err = pubsub.NewClient(ctx, utils.GetProjectID())
	if err != nil {
		log.Fatal(err)
	}

	secret = os.Getenv(core.EnvKeyChannelSecret)
	switch entryPoint {
	case RouteWebHook:
		parentTopic = pubSubClient.Topic(parentTopicName)
	case RouteParentFunctions:
		childTopic = pubSubClient.Topic(childTopicName)
	case RouteChildFunctions:
		setFunction()

		token := os.Getenv(core.EnvKeyChannelAccessToken)

		client, err := linebot.New(secret, token)
		if err != nil {
			log.Fatal(err)
		}

		operation = &core.Operation{
			Client: client,
			Tracer: tracer,
		}
	case RouteLiffFull,
		RouteLiffTall,
		RouteLiffCompact:
		// nop
	default:
		// nop
	}
}
