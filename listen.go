package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
	"io"
	"os"
	"sync/atomic"
	"time"
)

func main() {
	projectId := "golang-dev-1"
	subId := "my-sub"
	err := pullMsgs(os.Stdout, projectId, subId)
	if err != nil {
		panic(err)
	}
}
func pullMsgs(w io.Writer, projectID, subID string) error {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %w", err)
	}
	defer client.Close()

	sub := client.Subscription(subID)

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	var received int32
	err = sub.Receive(ctx, func(_ context.Context, msg *pubsub.Message) {
		fmt.Fprintf(w, "Got message: %q\n", string(msg.Data))
		atomic.AddInt32(&received, 1)
		msg.Ack()
	})
	if err != nil {
		return fmt.Errorf("sub.Receive: %w", err)
	}
	fmt.Fprintf(w, "Received %d messages\n", received)

	return nil
}
