package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"app/actions"

	frameworkAction "github.com/olbrichattila/evmagic/pkg/actions/framework-action"
	"github.com/olbrichattila/evmagic/pkg/connector"
)

const (
	awsURI       = "http://localhost:4566"
	awsRegion    = "us-east-1"
	awsAccountID = "000000000000"
	topicName    = "post-processed"
	actionType   = "blog-received"
)

type blogAction struct {
	Blog string `json:"blog"`
}

func main() {
	x := 10000
	for x > 0 {
		x--
		sendMessage(fmt.Sprintf("This is a blog post %d", x))
	}
	// sendMessage(getPrompt())
}

func getPrompt() string {
	scanner := bufio.NewScanner(os.Stdin)
	lines := ""

	for {
		if !scanner.Scan() {
			break
		}
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			break
		}

		lines = lines + text + "\n"

	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading input: %s", err.Error()))
	}

	return lines
}

func sendMessage(messageToPublish string) {
	snsPublisher, err := connector.Publisher(connector.TypeSNS)
	if err != nil {
		fmt.Println(err)
		return
	}

	frameworkAction.PublishFromStruct[actions.BlogReceivedAction](snsPublisher, topicName, actionType, actions.BlogReceivedAction{
		CreatedAt: time.Now(),
		CreatedBy: "John Doe",
		Blog:      messageToPublish,
	}, nil)
}
