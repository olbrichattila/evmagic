// Package main is for test
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/olbrichattila/evmagic/pkg/actions/action"
	"github.com/olbrichattila/evmagic/pkg/connector"
)

type myAction struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func main() {
	// Listen for termination signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	sqsHandler, err := connector.Handler(connector.TypeSQS)
	if err != nil {
		fmt.Println(err)
		return
	}

	snsPublisher, err := connector.Publisher(connector.TypeSNS)
	if err != nil {
		fmt.Println(err)
		return
	}

	// snsPublisher.Publish("new-posts", []byte("Hello"))

	sqsHandler.Handle("spam-check-queue", func(topic string, message []byte) [][]byte {
		// fmt.Println(topic, string(message))
		// act, err := connector.AsSNSAction(message)
		// fmt.Println(topic, act.Message, err)

		return nil
	})

	sqsHandler.Handle("profanity-check-queue", func(topic string, message []byte) [][]byte {
		// fmt.Println(topic, string(message))
		// act, err := connector.AsSNSAction(message)
		// fmt.Println(topic, act.Message, err)

		// TODO this is the actionType router logic, pass the original message if action type equals
		fullAct, err := connector.AsAction[action.ActionContent]([]byte(message))
		fmt.Println(fullAct.ActionType, err)

		fmt.Println(fullAct.Content)

		return nil
	})

	go sqsHandler.Run()

	// snsPublisher.Publish("new-posts", []byte("Hello1"))
	// snsPublisher.Publish("new-posts", []byte("Hello2"))
	// snsPublisher.Publish("new-posts", []byte("Hello3"))
	// snsPublisher.Publish("new-posts", []byte("Hello4"))
	// snsPublisher.Publish("new-posts", []byte("Hello5"))
	helloAct, _ := action.New("typedata", myAction{Name: "John Doe", Email: "jdoe@jd.com", Address: "United States"})
	snsPublisher.Publish("new-posts", helloAct)

	<-sigs // Block here until signal received
	fmt.Println("Term signal received, shutting down...")

}
