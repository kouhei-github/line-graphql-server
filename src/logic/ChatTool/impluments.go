package ChatTool

type chatConnect interface {
	FindUserByEmail(email string) error
	PostMessage(channelId string, userId string, message string) error
}
