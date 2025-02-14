package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/swayamduhan/pubsub-go/subscriber/client"
)

func main() {
	fmt.Println("starting redis client ...")
	client.InitClient()
	defer client.RedisClient.Close()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter channel names ( seperated by comma ) to listen :")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("error reading input")
	}

	channels := strings.Split(strings.TrimSpace(input), ",")

	if len(channels) == 0 || channels[0] == "" {
		log.Fatal("enter atleast one channel")
	}

	client.ListenToChannel(channels)
}