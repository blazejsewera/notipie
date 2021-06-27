package domain

type App struct {
	ID           string
	Name         string
	SmallIconURL string
	BigIconURL   string
	tags         []*Tag
}

func (a *App) Send(notification Notification) {
}
