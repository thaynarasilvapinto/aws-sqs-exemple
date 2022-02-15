package sheduled

import (
	"fmt"
	"src/config"
	"src/sqs"
	"time"
)

func Bath(bath string, config *config.Config) {
	fmt.Println("Start bath...")

	for {
		time.Sleep(1 * time.Second)
		sqs.ReadMessage(config, bath)
	}
}
