package domain

import (
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"go.uber.org/zap"
	"sync"
)

type User struct {
	ID                 string
	Username           string
	NotificationChan   chan Notification
	Tags               []*Tag
	repo               NotificationRepository
	tagsMutex          sync.Mutex
	lastNotificationID string
	l                  *zap.Logger
}

func NewUser(id, username string, repo NotificationRepository) *User {
	return &User{
		ID:       id,
		Username: username,
		repo:     repo,
		l:        log.For("domain").Named("user").With(zap.String("userID", id), zap.String("username", username)),
	}
}

func (u *User) Start() {
	if u.NotificationChan == nil {
		u.NotificationChan = make(chan Notification)
	}
	go func() {
		for {
			u.Receive(<-u.NotificationChan)
		}
	}()

	u.l.Debug("started user")
}

func (u *User) Receive(notification Notification) {
	u.logRxNotification(notification)
	if notification.ID != u.lastNotificationID {
		u.repo.SaveNotification(notification)
		u.lastNotificationID = notification.ID
		u.logNotificationSaved(notification)
	} else {
		u.l.Debug("notificationID same as lastNotificationID", zap.String("notificationID", notification.ID))
	}
}

func (u *User) logRxNotification(notification Notification) {
	u.l.Debug(
		"received notification",
		zap.String("notificationID", notification.ID),
		zap.String("lastNotificationID", u.lastNotificationID),
	)
}

func (u *User) logNotificationSaved(notification Notification) {
	u.l.Info(
		"notification saved",
		zap.String("notificationID", notification.ID),
		zap.String("notificationTitle", notification.Title),
		zap.String("appName", notification.App.Name),
	)
}

func (u *User) SubscribeToTag(tag *Tag) {
	u.tagsMutex.Lock()
	defer u.tagsMutex.Unlock()

	u.Tags = append(u.Tags, tag)
	tag.registerUser(u)
	u.l.Info("subscribed to tag", zap.String("tagName", tag.Name))
}

func (u *User) UnsubscribeFromTag(name string) error {
	var err error
	var tag *Tag

	u.tagsMutex.Lock()
	defer u.tagsMutex.Unlock()

	u.Tags, tag, err = removeTag(u.Tags, name)
	if err != nil {
		u.l.Warn("could not unsubscribe from tag", zap.String("tagName", name), zap.Error(err))
		return err
	}
	tag.unregisterUser(u.ID)
	u.l.Info("removed tag from user", zap.String("tagName", name))
	return nil
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
