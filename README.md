# SQS Playground

Playing around with AWS SQS and Go.

### How to play

Get it running by executing docker-compose and applying the terraform:

```terminal
# start localstack container with SQS
docker-compose up -d

# wait a sec until localstack is ready then create the queues
terraform apply -auto-approve
```

The consumer command listen to the queue "playground" created by terraform.
You can build it with `make`.

```terminal
# make and running the consumer
make && builds/consumer
```

Producing messages with awscli is easy:

```terminal
aws --endpoint http://localhost:4576 sqs send-message --queue-url 'http://localhost:4576/queue/playground' --message-body 'hello world'
```

### License

This project is licensed under [GLWTPL](https://github.com/me-shaon/GLWTPL/blob/master/LICENSE)
