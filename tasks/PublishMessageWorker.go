package tasks

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"hubble.lead/model"
	"hubble.lead/googlequeue"
)

func ProcessNotificationsForHostWorker(c *gin.Context) {
	hostName := c.Param("host")
	ctx := appengine.NewContext(c.Request)

	messages, err := model.GetNotificationsForHosts(ctx, false, hostName)

	for _,message := range *messages {
		 err = googlequeue.PublishMessageForHost(ctx, hostName, message)
		 if err != nil {
		 	log.Errorf(ctx, err.Error())
		 }
	}

}
