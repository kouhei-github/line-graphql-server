package teamSlack

import (
	"fmt"
	"github.com/slack-go/slack"
)

func (connect Connector) FindUserByEmail(email string) error {
	user, err := connect.Client.GetUserByEmail(email)
	if err != nil {
		return err
	}
	fmt.Println("user")
	fmt.Println(user)
	return nil
}

func (connect Connector) PostMessage(channelId string, userId string, message string) error {
	_, _, err := connect.Client.PostMessage(
		"C0410SY5JUT",
		slack.MsgOptionPostEphemeral("U03P9NC0MFC"),
		slack.MsgOptionText("エラー通知", true),
	)
	if err != nil {
		return err
	}
	return nil
}
