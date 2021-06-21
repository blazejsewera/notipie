package domain

type Application struct {
	ID           string
	Name         string
	SmallIconURL string
	BigIconURL   string
	handler      Handler
}

func (a Application) Send(notification Notification) {
	err := a.handler.HandleNotification(notification)
	if err != nil {
		a.handler.HandleError(err)
	}
}

func (a Application) Equals(app *Application) bool {
	return a.ID == app.ID
}
