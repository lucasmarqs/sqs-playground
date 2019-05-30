package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/lucasmarqs/sqs-playground/conf"
)

func main() {
	fmt.Println("--=[ SQS Playground | Consumer ]=--")

	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(endpoints.UsEast1RegionID),
		Endpoint: aws.String(conf.SQSEndpoint),
	}))
	svc := sqs.New(sess)

	receiveMessageInput := sqs.ReceiveMessageInput{QueueUrl: aws.String(conf.PlaygroundQueueURL)}
	receiveMessageInput.SetMaxNumberOfMessages(10)
	received, err := svc.ReceiveMessage(&receiveMessageInput)
	if err != nil {
		panic(err)
	}

	for _, msg := range received.Messages {
		fmt.Println("[received]", *msg)
	}
}
