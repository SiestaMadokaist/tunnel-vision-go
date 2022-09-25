package ClientHub

import (
	consumer "github.com/academy-software/go-aws-sqs-consumer"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type ClientHub struct {
	consumer *consumer.Consumer
}

func New(queueURL string) ClientHub {
	sess := session.Must(session.NewSessionWithOptions(session.Options{ SharedConfigState: session.SharedConfigEnable }));
	config := consumer.Config{ AwsSession: sess, Receivers: 1, SqsMaxNumberOfMessages: 10, SqsMessageVisibilityTimeout: 20, PollDelayInMilliseconds: 100}
	c := consumer.New(queueURL, handle, &config)
	return ClientHub{ consumer: &c }
}

func handle(m *sqs.Message) error {
	return nil;
}

