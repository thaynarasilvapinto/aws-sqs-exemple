package sheduled

import (
	"src/sqs"

	"github.com/jasonlvhit/gocron"
)

func Bath() {
	gocron.Every(3).Seconds().Do(task)
	<-gocron.Start()
}

func task() {
	sqs.ReadMessage()
}
