package googlequeue

import (
	"cloud.google.com/go/pubsub"
	"os"
	"context"
	"google.golang.org/appengine/log"
)

func GetClientInstance(ctx context.Context) (c *pubsub.Client) {
	mu.Lock()                    // <--- Unnecessary locking if instance already created
	defer mu.Unlock()

	if nil == pubsubClient {
		pubsubClient = getClient(ctx)
	}
	return pubsubClient
}

func getClient(ctx context.Context) (c *pubsub.Client) {
	proj := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if proj == "" {
		log.Errorf(ctx,"GOOGLE_CLOUD_PROJECT environment variable must be set.\n")
		os.Exit(1)
	}
	client, err := pubsub.NewClient(ctx, proj)
	if err != nil {
		log.Errorf(ctx, "Could not create pubsub Client: %v", err)
		return nil
	}

	return client
}