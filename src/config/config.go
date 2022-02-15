package config

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type Config struct {
	Queue    *string
	Timeout  *int64
	QueueURL *string
	Svc      *sqs.SQS
	Handle   string
}

func Configuration() *Config {

	var config Config = Config{}

	var queue = flag.String("q", "terraform-example-queue", "The name of the queue")
	config.Timeout = flag.Int64("t", 5, "How long, in seconds, that the message is hidden from others")
	flag.Parse()

	if *queue == "" {
		fmt.Println("You must supply the name of a queue (-q QUEUE)")
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	config.Svc = sqs.New(sess)

	urlResult, _ := config.Svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queue,
	})

	config.QueueURL = urlResult.QueueUrl

	return &config
}
