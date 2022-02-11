package sheduled

import (
	"fmt"
	"src/config"
	"src/sqs"

	"github.com/jasonlvhit/gocron"
)

func Bath(bath string, config *config.Config) {
	fmt.Println("Start bath...")
	gocron.Every(3).Seconds().Do(task, bath, config)
	<-gocron.Start()
}

func task(bath string, config *config.Config) {
	sqs.ReadMessage(config, bath)
}
