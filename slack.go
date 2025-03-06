package main

import (
	"fmt"
        "os"

	"github.com/slack-go/slack"
)

func notifySlack(name string, namespace string, environment string, tag string, slackmoji string) {
	message := fmt.Sprintf(":%s: Deployed %s %s %s", slackmoji, environment, name, tag)
	api := slack.New(os.Getenv("SLACK_TOKEN"))
	channelID, timestamp, err := api.PostMessage(
		os.Getenv("SLACK_CHANNEL"),
		slack.MsgOptionText(message, false),
		slack.MsgOptionAsUser(true), // Add this if you want that the bot would post message as a user, otherwise it will send response using the default slackbot
	)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
