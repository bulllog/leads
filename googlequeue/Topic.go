package googlequeue

import (
	"cloud.google.com/go/pubsub"
	"google.golang.org/appengine/log"
	"context"
	"errors"
	"encoding/json"
	"hubble.lead/model"
)

var (
	topicMap map[string]*pubsub.Topic = make(map[string]*pubsub.Topic)
	//googleClient *pubsub.Client = GetClientInstance()
)


func GetTopicForHost(ctx context.Context, hostName string) *pubsub.Topic {
	if topicMap[hostName] == nil {
		t, err := createTopicForHost(ctx, hostName)
		if err == nil {
			topicMap[hostName] = t
			return t
		} else {
			log.Errorf(ctx, err.Error())
			return nil
		}
	} else {
		return topicMap[hostName]
	}
	return nil
}

func createTopicForHost(ctx context.Context, host string) (*pubsub.Topic, error) {
	t, err := GetClientInstance(ctx).CreateTopic(ctx, getTopicSubName(host))

	if err != nil {
		log.Errorf(ctx, "Error in creating topic for host " + host)
		return &(pubsub.Topic{}), nil
	}
	return t, nil
}

func PublishMessageForHost(ctx context.Context, host string, entity model.NotificationEntity) error {
	t := GetTopicForHost(ctx, host)
	if t == nil {
		return errors.New("Topic not found for host - "+host)
	}

	message,err := json.Marshal(entity)
	if err != nil {
		return errors.New("Error in json parsing " + err.Error())
	}

	//Collect Publish Result to get published message id
	t.Publish(ctx, &pubsub.Message{Data:[]byte(message)})

	return nil



}

