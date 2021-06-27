package domain

type User struct {
	ID       string
	Username string
	repo     NotificationRepository
	tags     []*Tag
}

func (u User) Receive(notification Notification) {
}

func (u *User) SubscribeTo(tag *Tag) {
	u.tags = append(u.tags, tag)
	tag.RegisterUser(u)
}

type NotificationRepository interface {
	SaveNotification(Notification)
	GetAllNotifications() []Notification
}
