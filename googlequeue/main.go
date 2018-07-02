package googlequeue

import (
	"cloud.google.com/go/pubsub"
	"sync"
	"time"
)


var (
	pubsubClient *pubsub.Client
	mu sync.Mutex
	//ctx context.Context = appengine.BackgroundContext()
)

const (
	TOPIC_SUBS_POSTFIX = "_webhook_endpoint"
	ACK_DEAD_LINE = 10 * time.Minute
)

func getTopicSubName(host string) string {
	return TOPIC_SUBS_POSTFIX
}


