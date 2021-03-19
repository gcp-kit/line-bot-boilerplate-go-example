package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime/debug"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"github.com/gcp-kit/line-bot-boilerplate-go-example/function"
	"github.com/gcp-kit/line-bot-boilerplate-go-example/pkg/utils"
	"github.com/gcp-kit/line-bot-boilerplate-go/core"
	"github.com/gcp-kit/line-bot-boilerplate-go/gae"
	"github.com/gcp-kit/stalog"
	"github.com/labstack/echo/v4"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	defer func() {
		if rec := recover(); rec != nil {
			debug.PrintStack()
		}
	}()

	e := echo.New()
	e.HideBanner = true

	ctx := context.Background()

	projectID := utils.GetProjectID()

	cloudtasksClient, err := cloudtasks.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to initialize cloudtasks client: %+v", err)
	}

	lineSecret := os.Getenv(core.EnvKeyChannelSecret)
	token := os.Getenv(core.EnvKeyChannelAccessToken)

	lineClient, err := linebot.New(lineSecret, token)
	if err != nil {
		log.Fatalf("failed to initialize linebot client: %+v", err)
	}

	const (
		locationID      = "asia-northeast1"
		queuePathFormat = "projects/%s/locations/%s/queues/%s"
	)

	var (
		parentQueue = fmt.Sprintf("%s-parent", projectID)
		childQueue  = fmt.Sprintf("%s-child", projectID)
		serviceName = utils.GetServiceName()
	)

	g := e.Group("/line/")
	{
		props := &gae.Props{
			QueuePath:   fmt.Sprintf(queuePathFormat, projectID, locationID, parentQueue),
			RelativeURI: "/line/tq/parent",
			Service:     serviceName,
		}
		props.SetTQClient(cloudtasksClient)
		props.SetSecret(lineSecret)
		g.POST("webhook", func(c echo.Context) error {
			props.LineWebHook(c.Response().Writer, c.Request())
			return nil
		})
	}

	tq := g.Group("tq/")
	{
		tq.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				t, ok := c.Request().Header["X-Appengine-Taskname"]
				if !ok || len(t[0]) == 0 {
					log.Println("Invalid Task: No X-Appengine-Taskname request header found")
					return c.String(http.StatusBadRequest, "Bad Request - Invalid Task\n")
				}

				var (
					taskName  = t[0]
					queueName string
				)
				q, ok := c.Request().Header["X-Appengine-Queuename"]
				if ok {
					queueName = q[0]
				}

				fmt.Printf("Completed task: task queue(%s), task name(%s)\n", queueName, taskName)
				return next(c)
			}
		})

		props := &gae.Props{
			QueuePath:   fmt.Sprintf(queuePathFormat, projectID, locationID, childQueue),
			RelativeURI: "/line/tq/child",
			Service:     serviceName,
		}
		props.SetTQClient(cloudtasksClient)
		tq.POST("parent", func(c echo.Context) error {
			body, err := ioutil.ReadAll(c.Request().Body)
			if err != nil {
				return err
			}
			return props.ParentFunctions(ctx, body)
		})

		op := operation(lineClient)
		tq.POST("child", func(c echo.Context) error {
			body, err := ioutil.ReadAll(c.Request().Body)
			if err != nil {
				return err
			}
			return props.ChildFunctions(ctx, op, body)
		})
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	cfg := utils.NewStalogConfig()
	e.Use(stalog.RequestLoggingWithEcho(cfg))

	e.Logger.Fatal(e.Start(":" + port))
}

func operation(client *linebot.Client) *core.Operation {
	tracer := new(core.Tracer)
	tracer.Function = map[core.TracerName]func(context.Context, *core.Operation, *linebot.Event) *core.TracerResp{
		core.TracerFollowEvent:     function.FollowEvent,
		core.TracerUnfollowEvent:   function.UnfollowEvent,
		core.TracerTextMessage:     function.TextEvent,
		core.TracerLocationMessage: function.LocationEvent,
	}
	systemError := linebot.NewTextMessage("システムエラーです。")
	return &core.Operation{
		ErrMessage: []linebot.SendingMessage{systemError},
		Client:     client,
		Tracer:     tracer,
	}
}
