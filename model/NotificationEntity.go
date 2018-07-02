package model

import (
	"time"
	"context"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"errors"
)

const (
	NOTIFICATION_ENTITY_NAME = "NotificationEntity"
	MESSAGE_ID = "MessageId"
	USER_ID = "UserId"
	CHANGE_ID = "ChangeId"
	RECIEVED_AT = "Recieved_At"
	HAS_CONSUMED = "Hash_Consumed"
)

type NotificationEntity struct {
	HostName string `json:HostName`
	UserId string `json:UserId`
	ChangeId string `json:ChangeId`
	OAuthId string `json:OAuthId`
	Token string `json:Token`
	RecievedAt time.Time `json:ReceivedAt`
	HasConsumed bool
}


func Insert(ctx context.Context,
	entity *NotificationEntity) (output bool,err error) {
	key := datastore.NewIncompleteKey(ctx, NOTIFICATION_ENTITY_NAME, nil)

	_, err = datastore.Put(ctx, key, entity)

	if err != nil {
		log.Errorf(ctx, "ERROR INSERTING Host: %v", err.Error())
		return false, err
	}

	return true, nil
}

func GetNotificationsForHosts(ctx context.Context,
	hasConsumed bool, hostName string)(*[]NotificationEntity, error)  {
	query := datastore.NewQuery(NOTIFICATION_ENTITY_NAME).Filter(HOST_NAME, hostName)

	if hasConsumed == false {
		query.Filter(HAS_CONSUMED, false)
	}

	var notifications[]NotificationEntity

	_, err := query.GetAll(ctx, &notifications)

	if err != nil {
		return nil, errors.New(" Error Getting messages from datastore " + hostName + err.Error())
	}

	return &notifications, nil
}
