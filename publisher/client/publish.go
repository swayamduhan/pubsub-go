package client

import (
	"context"
	"log"
)


func PublishMessage(msg string, channel string) {
	ctx := context.Background()
	err := RedisClient.Publish(ctx, channel, msg).Err()
	if err != nil {
		log.Printf("error publishing message to channel %v", channel)
	}
	log.Printf("Message published to channel %v successfully!", channel)
}