// Package main is for test
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"app/routes"
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

	sqsHandler, snsPublisher, err := initQueues()
	if err != nil {
		fmt.Println(err)
		return
	}

	routes.InitRoutes(sqsHandler)
	go sqsHandler.Run()
	sendTestMessages(snsPublisher)

	<-sigs // Block here until signal received
	fmt.Println("Term signal received, shutting down...")
}
