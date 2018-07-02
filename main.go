package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"fmt"
	"hubble.lead/services"
	"hubble.lead/cron"
	"hubble.lead/tasks"
)

const (
	VERSION = "/v1/lead"
	HOST = "host"
	WEBHOOK = "webhook"
	CRON = "cron"
	TASK_QUEUE = "tasks"
)

func init() {
	router := gin.New()
	v1 := router.Group(VERSION)

	host := v1.Group(HOST)
	host.POST("/:hostName", services.RegisterHostHandler)
	host.PUT("/:hostName", services.UpdateHostHandler)

	webhook := v1.Group(WEBHOOK)
	webhook.POST("/:hostName", services.WebhookPostHandler)
	webhook.GET("/:hostName", services.SendConfirmation)

	cronGroup := v1.Group(CRON)
	cronGroup.GET("/publisher", cron.UpdateQueueWithNewMessages)

	taskGroup:= v1.Group(TASK_QUEUE)
	taskGroup.GET("/:host", tasks.ProcessNotificationsForHostWorker)

	http.Handle("/", router)
	http.HandleFunc("/healthCheck", healthCheck)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintln(w, "you are lucky to have this")
	}
}



