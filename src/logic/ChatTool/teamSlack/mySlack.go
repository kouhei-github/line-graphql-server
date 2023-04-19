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

func (connect Connector) PostMessage(message string) error {
	_, _, err := connect.Client.PostMessage(
		connect.ChannelId,
		slack.MsgOptionPostEphemeral(connect.UserId),
		slack.MsgOptionText(message, true),
	)
	if err != nil {
		return err
	}
	return nil
}
