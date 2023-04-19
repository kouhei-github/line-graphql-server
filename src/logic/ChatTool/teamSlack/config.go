package teamSlack

import (
	"github.com/joho/godotenv"
	"github.com/slack-go/slack"
	"os"
)

type Connector struct {
	Client *slack.Client
}

func NewSlack() (Connector, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return Connector{Client: &slack.Client{}}, err
	}
	slackToken := os.Getenv("SLACK_TOKEN")
	client := slack.New(slackToken)
	return Connector{Client: client}, nil
}
