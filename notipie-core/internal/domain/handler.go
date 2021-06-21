package domain

type Handler interface {
	HandleNotification(Notification) error
	HandleError(error)
}

type HandlerError struct {
	msg string
}

func (h HandlerError) Error() string {
	return h.msg
}
