package domain

type User struct {
	ID                 string
	Username           string
	NotificationChan   chan Notification
	repo               NotificationRepository
	tags               []*Tag
	lastNotificationID string
}

func (u *User) Listen() {
	if u.NotificationChan == nil {
		u.NotificationChan = make(chan Notification)
	}
	go func() {
		for {
			u.Receive(<-u.NotificationChan)
		}
	}()
}

func (u *User) Receive(notification Notification) {
	if notification.ID != u.lastNotificationID {
		u.repo.SaveNotification(notification)
		u.lastNotificationID = notification.ID
	}
}

func (u *User) SubscribeToTag(tag *Tag) {
	u.tags = append(u.tags, tag)
	tag.registerUser(u)
}

func (u *User) UnsubscribeFromTag(tag Tag) (err error) {
	u.tags, err = removeTag(u.tags, tag)
	return
}

func (u *User) GetAllNotifications() []Notification {
	return u.repo.GetNotifications(0, u.repo.GetNotificationCount())
}

func (u *User) GetLastNotifications(n int) []Notification {
	return u.repo.GetLastNotifications(n)
}

func (u *User) GetNotifications(from int, to int) []Notification {
	return u.repo.GetNotifications(from, to)
}

func (u *User) GetNotificationCount() int {
	return u.repo.GetNotificationCount()
}

type NotificationRepository interface {
	SaveNotification(notification Notification)
	GetLastNotifications(n int) []Notification
	GetNotifications(from, to int) []Notification
	GetNotificationCount() int
}
