package domain

type User struct {
	ID       string
	Username string
	repo     NotificationRepository
	tags     []*Tag
}

func (u *User) Receive(notification Notification) {
	u.repo.SaveNotification(notification)
}

func (u *User) SubscribeTo(tag *Tag) {
	u.tags = append(u.tags, tag)
	tag.RegisterUser(u)
}

func (u *User) GetAllNotifications() []Notification {
	return u.repo.GetAllNotifications()
}

func (u *User) GetLastNotifications(n int) []Notification {
	return u.repo.GetLastNotifications(n)
}

type NotificationRepository interface {
	SaveNotification(Notification)
	GetAllNotifications() []Notification
	GetLastNotifications(int) []Notification
}
