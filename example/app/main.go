// Package main is for test
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"app/routes"

	"github.com/olbrichattila/evmagic/pkg/database/connection"
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

	go process()

	// sendTestMessages(snsPublisher)

	<-sigs // Block here until signal received
	fmt.Println("Term signal received, shutting down...")
}

func process() {
	db, err := connection.Open()
	if err != nil {
		fmt.Println("Cannot connect to the database", err.Error())
		os.Exit(1)
	}
	defer db.Close()
	sqsHandler, snsPublisher, err := initQueues(db)
	if err != nil {
		// TODO error log
		fmt.Println(err)
		return
	}

	routes.InitRoutes(sqsHandler, snsPublisher)
	sqsHandler.Run()
}
