package ChatTool

type chatConnect interface {
	FindUserByEmail(email string) error
	PostMessage(message string) error
}
