package domain

type User struct {
	ID       string
	Username string
	handler  UserHandler
}

func (u User) Receive(application Application, notification Notification) {
	err := u.handler.Handle(application, notification)
	if err != nil {
		u.handler.HandleError(err)
	}
}

type UserHandler interface {
	Handle(Application, Notification) error
	HandleError(error)
}

type UserHandlerError struct {
	msg string
}

func (e UserHandlerError) Error() string {
	return e.msg
}
