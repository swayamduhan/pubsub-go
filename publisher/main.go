package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/swayamduhan/pubsub-go/publisher/client"
)

func main() {
	log.Println("starting client..")
	client.InitClient()
	defer client.RedisClient.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter channel name : ")
		channel, _ := reader.ReadString('\n')
		channel = strings.TrimSpace(channel)
		if channel == "" {
			fmt.Println("Invalid channel name")
			continue
		}

		fmt.Println("Enter message : ")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)

		client.PublishMessage(msg, channel)

	}
}