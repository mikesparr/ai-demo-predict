package message

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mikesparr/ai-demo-predict/models"

	"cloud.google.com/go/pubsub"
)

func (producer Producer) UpdateBatch(br *models.BatchResponse) error {
	ctx := context.Background()
	fmt.Println("I ran UpdateBatch !!!")

	topic := producer.Topic
	feedbackJson, err := json.Marshal(br)
	if err != nil {
		return err
	}

	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(feedbackJson),
	})
	if _, err := res.Get(ctx); err != nil {
		return err
	}

	return nil
}
