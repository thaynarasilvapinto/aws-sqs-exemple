package sqs

import (
	"fmt"
	"src/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func ReadMessage(config *config.Config, bath string) {
	defer recoveryExecution()
	fmt.Println("Starting reading message...", bath)
	msgResult, err := config.Svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            config.QueueURL,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   config.Timeout,
	})

	if err != nil {
		fmt.Println("Error to read sqs")
		panic("erro to read sqs")
	}

	if len(msgResult.Messages) != 0 {
		handle := *msgResult.Messages[0].ReceiptHandle
		fmt.Println("Deleting reading message...", bath)
		DeleteMessage(handle, config)
		message := *msgResult.Messages[0].Body

		fmt.Println("Mensagem recebida: ", message, bath)
	} else {
		fmt.Println("Queue is empty", bath)
	}
}

func DeleteMessage(handle string, config *config.Config) {

	_, err := config.Svc.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      config.QueueURL,
		ReceiptHandle: &handle,
	})

	if err != nil {
		fmt.Println("Error to delete sqs")
		panic("erro to delte sqs")
	}
}

func PublishMessage(message string, config *config.Config) {
	defer recoveryExecution()
	_, err := config.Svc.SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(message),
		QueueUrl:     config.QueueURL,
	})

	if err != nil {
		fmt.Println("Error to publish sqs")
	}
}

func recoveryExecution() {
	if r := recover(); r != nil {
		fmt.Println("Panic recovery")
	}
}
