package domain

import (
	"fmt"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"go.uber.org/zap"
	"sync"
)

type App struct {
	ID             string
	Name           string
	IconURI        string
	CommandChan    chan Command
	tags           []*Tag
	tagsMutex      sync.Mutex
	commandHandler CommandHandler
	l              *zap.Logger
}

func NewApp(id, name, iconURI string, commandHandler CommandHandler) *App {
	return &App{
		ID:             id,
		Name:           name,
		IconURI:        iconURI,
		commandHandler: commandHandler,
		l:              log.For("domain").Named("app").With(zap.String("appID", id), zap.String("appName", name)),
	}
}

func (a *App) Start() {
	if a.CommandChan == nil {
		a.CommandChan = make(chan Command)
	}

	go func() {
		for {
			a.commandHandler.HandleCommand(<-a.CommandChan)
		}
	}()

	a.l.Debug("started app")
}

func (a *App) Send(notification Notification) error {
	if len(a.tags) == 0 {
		a.l.Warn("no tags attached to app")
		return &SendError{
			App:          a,
			Notification: notification,
		}
	}

	notification.App = a

	for _, tag := range a.tags {
		tag.NotificationChan <- notification
		a.logSentNotificationToTag(tag, notification)
	}

	return nil
}

func (a *App) logSentNotificationToTag(tag *Tag, notification Notification) {
	a.l.Info(
		"sent notification to tag",
		zap.String("tagName", tag.Name),
		zap.String("notificationID", notification.ID),
	)
}

func (a *App) AddTag(tag *Tag) {
	a.tagsMutex.Lock()
	defer a.tagsMutex.Unlock()

	a.tags = append(a.tags, tag)
	tag.registerApp(a)
	a.l.Info("added tag to app", zap.String("tagName", tag.Name))
}

func (a *App) RemoveTag(name string) error {
	var err error
	var tag *Tag

	a.tagsMutex.Lock()
	defer a.tagsMutex.Unlock()

	a.tags, tag, err = removeTag(a.tags, name)
	if err != nil {
		a.l.Warn("could not remove tag from app", zap.String("tagName", name), zap.Error(err))
		return err
	}
	tag.unregisterApp(a.ID)
	a.l.Info("removed tag from app", zap.String("tagName", name))
	return nil
}

func (a *App) GetTags() []*Tag {
	return a.tags
}

type SendError struct {
	*App
	Notification
	Tags []*Tag
}

func (e *SendError) Error() string {
	if len(e.Tags) == 0 {
		return fmt.Sprintf(NoTagsWhenSendErrorFormat, e.App.Name, e.App.ID, e.Notification)
	}

	var tags []string
	for _, tag := range e.Tags {
		tags = append(tags, tag.Name)
	}

	return fmt.Sprintf(NoUsersInTagsWhenSendErrorFormat, tags, e.App.Name, e.App.ID, e.Notification)
}

const (
	NoTagsWhenSendErrorFormat        = "no tags for %s#%s when sending %s"
	NoUsersInTagsWhenSendErrorFormat = "tags: %v for %s#%s did not have registered users when sending %s"
)

type CommandHandler interface {
	HandleCommand(command Command)
}
