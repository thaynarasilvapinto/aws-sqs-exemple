package sqs

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var (
	queue    *string
	timeout  *int64
	queueURL *string
	svc      *sqs.SQS
	handle   string
)

func configuration() {

	queue = flag.String("q", "file-exemplo", "The name of the queue")
	timeout = flag.Int64("t", 5, "How long, in seconds, that the message is hidden from others")
	flag.Parse()

	if *queue == "" {
		fmt.Println("You must supply the name of a queue (-q QUEUE)")
		return
	}

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc = sqs.New(sess)

	urlResult, _ := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queue,
	})

	queueURL = urlResult.QueueUrl

}

func ReadMessage() {
	configuration()
	msgResult, err := svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            queueURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   timeout,
	})

	if err != nil {
		fmt.Println("Error to read sqs")
		panic("erro to read sqs")
	}

	message := *msgResult.Messages[0].Body
	handle = *msgResult.Messages[0].ReceiptHandle

	fmt.Println("Mensagem recebida: ", message)
}

func DeleteMessage(handle string) {
	configuration()
	messageHandle := flag.String("m", handle, "The receipt handle of the message")
	flag.Parse()

	if *queue == "" || *messageHandle == "" {
		fmt.Println("You must supply a queue name (-q QUEUE) and message receipt handle (-m MESSAGE-HANDLE)")
		return
	}

	_, err := svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      queueURL,
		ReceiptHandle: messageHandle,
	})

	if err != nil {
		fmt.Println("Error to delete sqs")
		panic("erro to delte sqs")
	}
}

func PublishMessage(message string) {
	configuration()
	_, err := svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(message),
		QueueUrl:     queueURL,
	})

	if err != nil {
		fmt.Println("Error to publish sqs")
	}
}
