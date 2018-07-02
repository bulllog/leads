package services

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"strings"
	"net/http"
	"time"
	"encoding/json"
	"fmt"
	"hubble.lead/model"
	"hubble.lead/controllers"
)

const (
	GOOGLE_TOKEN_HEADER string = "x-google-channel-token"
	OAUTH_ID string = "oauthId"
	HOST_NAME string = "hostName"
	USER_LOGIN string = "userLogin"
)

func SendConfirmation(c *gin.Context) {
	c.Status(http.StatusOK)
	return
}

func WebhookPostHandler(c *gin.Context) {
	ctx := appengine.NewContext(c.Request)
	log.Infof(ctx, "********************* Receive Notification ******************")
	contentType := c.ContentType()
	c.Status(http.StatusAccepted)

	if contentType == "application/json" {
		data, err := c.GetRawData()
		if err != nil {
			log.Errorf(ctx, "Error getting post data" + err.Error())
		}

		var jsonData map[string]interface{}
		json.Unmarshal(data, &jsonData)

		changeId := fmt.Sprint(jsonData["id"])


		//postDataStr = "{ \"kind\": \"drive#change\", \"id\": \"82755\", \"selfLink\": \"https://www.googleapis.com/drive/v2/changes/82755\"}"

		// No need to get channel Id and Resource Id as the notification
		// will be processed only once for evey user.
		googleToken := c.GetHeader(GOOGLE_TOKEN_HEADER)

		gTokenMap := getGoogleTokenParams(googleToken)

		if len(gTokenMap) != 3 || changeId == "" {
			log.Errorf(ctx, "Parameter count is less than 3")
			return
		}

		oauthId := gTokenMap[OAUTH_ID]
		hostName := gTokenMap[HOST_NAME]
		userId := gTokenMap[USER_LOGIN]

		notification := &model.NotificationEntity{
			HostName: hostName,
			UserId: userId,
			ChangeId: changeId,
			OAuthId: oauthId,
			RecievedAt: time.Now(),
			HasConsumed: false,
		}

		response := controllers.RecieveWebhookNotification(ctx, notification)

		if response.IsError == true {
			log.Errorf(ctx,
				"Error saving notification object " + response.ErrMessage)
		}
	}

}

func getGoogleTokenParams(gHeader string) map[string]string {
	tokenMap := make(map[string]string)
	pairs := strings.Split(gHeader, "&")
	for _, element := range pairs {
		keyValuePair := strings.Split(element, "=")
		if len(keyValuePair) == 2 {
			tokenMap[keyValuePair[0]] = keyValuePair[1]
		}
	}
	return tokenMap

}
