package main

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/lucasmarqs/sqs-playground/conf"
	"github.com/sirupsen/logrus"
)

const (
	maxNumberOfMessages = 10
)

func main() {
	fmt.Printf(`
 .d8888b.   .d88888b.   .d8888b.
d88P  Y88b d88P" "Y88b d88P  Y88b
Y88b.      888     888 Y88b.
 "Y888b.   888     888  "Y888b.
    "Y88b. 888     888     "Y88b.
      "888 888 Y8b 888       "888
Y88b  d88P Y88b.Y8b88P Y88b  d88P
 "Y8888P"   "Y888888"   "Y8888P"
                  Y8b


8888888b.  888                                                                   888
888   Y88b 888                                                                   888
888    888 888                                                                   888
888   d88P 888  8888b.  888  888  .d88b.  888d888 .d88b.  888  888 88888b.   .d88888
8888888P"  888     "88b 888  888 d88P"88b 888P"  d88""88b 888  888 888 "88b d88" 888
888        888 .d888888 888  888 888  888 888    888  888 888  888 888  888 888  888
888        888 888  888 Y88b 888 Y88b 888 888    Y88..88P Y88b 888 888  888 Y88b 888
888        888 "Y888888  "Y88888  "Y88888 888     "Y88P"   "Y88888 888  888  "Y88888
                             888      888
                        Y8b d88P Y8b d88P
                         "Y88P"   "Y88P"


`)

	logrus.Info("setting up environment")

	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(endpoints.UsEast1RegionID),
		Endpoint: aws.String(conf.SQSEndpoint),
	}))
	svc := sqs.New(sess)

	receiveMessageInput := sqs.ReceiveMessageInput{QueueUrl: aws.String(conf.PlaygroundQueueURL)}
	receiveMessageInput.SetMaxNumberOfMessages(maxNumberOfMessages)

	ctx := context.Background()

	msgCh := make(chan *sqs.Message, maxNumberOfMessages)
	go func() {
		for msg := range msgCh {
			go handleNewMessage(ctx, svc, msg)
		}
	}()

	logrus.Info("waiting for new incoming messages...")
	for {
		received, err := svc.ReceiveMessageWithContext(ctx, &receiveMessageInput)
		if err != nil {
			panic(err)
		}

		for _, msg := range received.Messages {
			msgCh <- msg
		}
	}

}

func handleNewMessage(ctx context.Context, svc *sqs.SQS, msg *sqs.Message) {
	receipt := *msg.ReceiptHandle
	logrus.WithFields(logrus.Fields{
		"receipt": receipt,
		"body":    strings.Trim(*msg.Body, "\n"),
	}).Info("Message received")
	work := time.Duration(rand.Float64() * 3000)
	logrus.WithField("receipt", receipt).Infof("processing in %d ms", work)
	<-time.After(work * time.Millisecond)

	input := sqs.DeleteMessageInput{
		QueueUrl:      aws.String(conf.PlaygroundQueueURL),
		ReceiptHandle: &receipt,
	}
	_, err := svc.DeleteMessageWithContext(ctx, &input)
	if err != nil {
		logrus.WithField("receipt", receipt).Error("could not delete message")
		return
	}

	logrus.WithField("receipt", receipt).Info("work done and message deeted")
}
