package googlequeue

import (
	"cloud.google.com/go/pubsub"
	"google.golang.org/appengine/log"
	"golang.org/x/net/context"
)


func ResetSubscriberForHost(ctx context.Context, host string) error {
	const topicName = TOPIC_SUBS_POSTFIX
	//ctx := context.Background()

	t := GetTopicForHost(ctx, host)

	if t == nil {
		log.Errorf(ctx, "Failed to create topic for host " + host)
		return nil
	}
	//Remove subsrciption if exist
	deleteSubscription(ctx, host)

	config := pubsub.SubscriptionConfig{
		Topic: GetTopicForHost(ctx, host),
		AckDeadline: ACK_DEAD_LINE,
	}
	_, err := GetClientInstance(ctx).CreateSubscription(ctx, topicName, config)

	if err != nil {
		log.Errorf(ctx, "Error in creating topic for host " + host)
		return err
	}
	return nil
}


func deleteSubscription(ctx context.Context, host string) {
	c := GetClientInstance(ctx)
	sub := c.Subscription(getTopicSubName(host))

	if err := sub.Delete(ctx); err == nil {
		log.Infof(ctx,"Removed subscriber for host -- " + host)
	}
}

