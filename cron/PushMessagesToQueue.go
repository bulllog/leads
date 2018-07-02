package cron

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"

	"context"
	"google.golang.org/appengine/taskqueue"
	"hubble.lead/model"
)

func UpdateQueueWithNewMessages(c *gin.Context) {
	ctx := appengine.NewContext(c.Request)
	log.Infof(ctx, "Cron for publish message")

	hosts, err := model.GetAllHosts(ctx,true)

	processNotificationsForHosts(ctx, *hosts)
	if err != nil {
		return
	}


}

func processNotificationsForHosts(ctx context.Context,
	hostNames []string)  {
	for _, host := range hostNames {
		t := taskqueue.NewPOSTTask("v1/lead/tasks/"+host, nil)
		if _, err := taskqueue.Add(ctx, t, ""); err != nil {
			log.Errorf(ctx, "error while running worker for host - " + err.Error() + host)
		}
	}
}
