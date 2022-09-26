package ClientHub

import (
	"fmt"
	"time"

	consumer "github.com/academy-software/go-aws-sqs-consumer"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type IncomingFields struct {
	QueueURL string;
	Hostname string;
}

type OutgoingFields struct {
	QueueURL string;
	Hostname string;
}

type ClientHub struct {
	incoming *IncomingFields;
	outgoing *OutgoingFields;
	requestConsumer *consumer.Consumer;
}


func New(incoming *IncomingFields, outgoing *OutgoingFields) ClientHub {
	sess := session.Must(session.NewSessionWithOptions(session.Options{ SharedConfigState: session.SharedConfigEnable }));
	config := consumer.Config{ AwsSession: sess, Receivers: 1, SqsMaxNumberOfMessages: 10, SqsMessageVisibilityTimeout: 20, PollDelayInMilliseconds: 100}
	requestConsumer := consumer.New(incoming.QueueURL, handle, &config);
	hub := ClientHub{ incoming: incoming, outgoing: outgoing, requestConsumer: &requestConsumer };
	return hub;
}

func (hub *ClientHub) connect(i int) {
	fmt.Println(i);
	fmt.Println(hub.incoming.QueueURL);
	fmt.Println(hub.incoming.Hostname);
	fmt.Println(hub.outgoing.QueueURL);
	fmt.Println(hub.outgoing.Hostname);
}

func (hub *ClientHub) keepConnect(c chan string) {
	for i := 0; i < 1; i+=0 {
		time.Sleep(time.Second * 1);
		hub.connect(i)
	}
	c <- "ok"
}

func (hub *ClientHub) Start() {
	// var ok string;
	c := make(chan string, 1)
	// go hub.requestConsumer.Start();
	go hub.keepConnect(c);
	ok := <- c;
	fmt.Println(ok);
}

func handle(m *sqs.Message) error {
	fmt.Println(m.Body);
	return nil;
}

