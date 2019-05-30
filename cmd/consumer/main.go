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

	input := sqs.ListQueuesInput{QueueNamePrefix: aws.String("")}

	output, err := svc.ListQueues(&input)
	if err != nil {
		fmt.Println("Failed to list queues", err)
	}

	fmt.Println("Available Queues:")
	for _, name := range output.QueueUrls {
		fmt.Println(*name)
	}
}
