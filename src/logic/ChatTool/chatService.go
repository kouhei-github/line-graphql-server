package ChatTool

import (
	"fmt"
	"github.com/meetup/logic"
	"strings"
)

func GetUserEmail(connect chatConnect, email string) error {
	if !strings.Contains(email, "gmail.com") {
		return logic.MyError{Message: "Emailアドレスがdym.jpのドメインではありません"}
	}
	err := connect.FindUserByEmail(email)
	if err != nil {
		fmt.Println(err)
		return logic.MyError{Message: "Emailアドレスで検索できませんでした"}
	}

	return nil
}

func PostMessageToChatTool(connect chatConnect, message string) error {
	err := connect.PostMessage(message)
	if err != nil {
		fmt.Println(err)
		return logic.MyError{Message: "メッセージを送信できませんでした"}
	}
	return nil
}
