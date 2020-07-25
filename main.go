package main

import (
	"context"
	"log"
	"os"

	"github.com/gcp-kit/line-bot-boilerplate-go-example/functions"
	"github.com/gcp-kit/line-bot-boilerplate-go/cmd"
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
	fp, err := os.Open("functions/.env.yaml")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	cfg := new(config)
	if err := yaml.NewDecoder(fp).Decode(cfg); err != nil {
		panic(err)
	}

	os.Setenv(cmd.EnvKeyMid, cfg.Mid)
	os.Setenv(cmd.EnvKeyBotName, cfg.BotName)
	os.Setenv(cmd.EnvKeyChannelSecret, cfg.ChannelSecret)
	os.Setenv(cmd.EnvKeyChannelAccessToken, cfg.ChannelAccessToken)

	tracer := &cmd.Tracer{
		Function: map[cmd.TracerName]func(context.Context, *cmd.Operation, *linebot.Event) *cmd.TracerResp{
			cmd.TracerFollowEvent:     functions.FollowEvent,
			cmd.TracerUnfollowEvent:   functions.UnfollowEvent,
			cmd.TracerTextMessage:     functions.TextEvent,
			cmd.TracerLocationMessage: functions.LocationEvent,
		},
		LiffFunc: map[string]func(ctx *gin.Context){
			"liff_c": functions.Liff, // Compact
		},
	}

	engine := gin.Default()
	engine.LoadHTMLGlob("functions/templates/*.tmpl")
	ctx := context.Background()
	if err := tracer.Execute(ctx, engine); err != nil {
		log.Fatal(err)
	}
}
