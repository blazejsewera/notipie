package domain

import "sync"

type User struct {
	ID                 string
	Username           string
	NotificationChan   chan Notification
	Tags               []*Tag
	repo               NotificationRepository
	tagsMutex          sync.Mutex
	lastNotificationID string
}

func NewUser(id, username string, repo NotificationRepository) *User {
	return &User{ID: id, Username: username, repo: repo}
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
	u.tagsMutex.Lock()
	u.Tags = append(u.Tags, tag)
	u.tagsMutex.Unlock()
	tag.registerUser(u)
}

func (u *User) UnsubscribeFromTag(name string) (err error) {
	u.tagsMutex.Lock()
	defer u.tagsMutex.Unlock()
	u.Tags, err = removeTag(u.Tags, name)
	return
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

func (u *User) Respond(notification Notification, command Command) {
	app := notification.App
	app.CommandChan <- command
}

type NotificationRepository interface {
	SaveNotification(notification Notification)
	GetLastNotifications(n int) []Notification
	GetNotifications(from, to int) []Notification
	GetNotificationCount() int
}
