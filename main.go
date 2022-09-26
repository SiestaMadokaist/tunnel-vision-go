package main

import (
	"github.com/SiestaMadokaist/tunnel-vision-go/service/ClientHub"
)

func main() {
	// signalChan := make(chan os.Signal, 1)
	// wg := sync.WaitGroup{};
	// wg.Add(1);
	// signal.Notify(signalChan, os.Interrupt)
	// go func() {
	// 	defer wg.Done();
	// 	<- signalChan
		
	// }();

	requestQueue := "https://sqs.ap-southeast-1.amazonaws.com/674152176016/ramadoka-request-1"
	responseQueue := "https://sqs.ap-southeast-1.amazonaws.com/674152176016/ramadoka-response-1"

	incoming := ClientHub.IncomingFields{ QueueURL: requestQueue, Hostname:  "incoming-host" };
	outgoing := ClientHub.OutgoingFields{ QueueURL: responseQueue, Hostname: "outgoing-host" };
	hub := ClientHub.New(&incoming, &outgoing);
	hub.Start()
}