package domain

type User struct {
	ID                 string
	Username           string
	repo               NotificationRepository
	tags               []*Tag
	lastNotificationID string
}

func (u *User) Receive(notification Notification) {
	if notification.ID != u.lastNotificationID {
		u.repo.SaveNotification(notification)
		u.lastNotificationID = notification.ID
	}
}

func (u *User) SubscribeToTag(tag *Tag) {
	u.tags = append(u.tags, tag)
	tag.RegisterUser(u)
}

// TODO: Add UnsubscribeFromTag func

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
