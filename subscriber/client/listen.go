package client

import (
	"context"
	"fmt"
)

func ListenToChannel(channels []string) {
	ctx := context.Background()
	pubsub := RedisClient.Subscribe(ctx, channels...)
	defer pubsub.Close()

	fmt.Printf("Listening to channels %v - \n", channels)

	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Printf("Message received from channel %s : %s\n", msg.Channel, msg.Payload)
	}
}