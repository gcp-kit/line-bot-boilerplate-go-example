package main

import (
	"context"
	"log"
	"os"

	"github.com/gcp-kit/line-bot-boilerplate-go-example/function"
	"github.com/gcp-kit/line-bot-boilerplate-go/core"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/linebot"
	"gopkg.in/yaml.v2"
)

type config struct {
	GinMode            string `yaml:"GIN_MODE"`
	BotName            string `yaml:"BOT_NAME"`
	Mid                string `yaml:"MID"`
	ChannelSecret      string `yaml:"CHANNEL_SECRET"`
	ChannelAccessToken string `yaml:"CHANNEL_ACCESS_TOKEN"`
}

func main() {
	fp, err := os.Open("../env/.env.yaml")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	cfg := new(config)
	if err := yaml.NewDecoder(fp).Decode(cfg); err != nil {
		panic(err)
	}

	os.Setenv(core.EnvKeyMid, cfg.Mid)
	os.Setenv(core.EnvKeyBotName, cfg.BotName)
	os.Setenv(core.EnvKeyChannelSecret, cfg.ChannelSecret)
	os.Setenv(core.EnvKeyChannelAccessToken, cfg.ChannelAccessToken)

	tracer := &core.Tracer{
		Function: map[core.TracerName]func(context.Context, *core.Operation, *linebot.Event) *core.TracerResp{
			core.TracerFollowEvent:     function.FollowEvent,
			core.TracerUnfollowEvent:   function.UnfollowEvent,
			core.TracerTextMessage:     function.TextEvent,
			core.TracerLocationMessage: function.LocationEvent,
		},
		LiffFunc: map[string]func(ctx *gin.Context){
			"liff_c": function.Liff, // Compact
		},
	}

	engine := gin.Default()
	engine.LoadHTMLGlob("../templates/*.tmpl")
	ctx := context.Background()
	if err := tracer.Execute(ctx, engine); err != nil {
		log.Fatal(err)
	}
}
