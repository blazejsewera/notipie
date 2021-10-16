package domain

import (
	"fmt"
	"sync"
)

type App struct {
	ID             string
	Name           string
	SmallIconURL   string
	BigIconURL     string
	CommandChan    chan Command
	tags           []*Tag
	tagsMutex      sync.Mutex
	commandHandler CommandHandler
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
}

func (a *App) Send(notification Notification) error {
	if len(a.tags) == 0 {
		return &SendError{
			App:          a,
			Notification: notification,
		}
	}

	for _, tag := range a.tags {
		tag.NotificationChan <- notification
	}

	return nil
}

func (a *App) AddTag(tag *Tag) {
	a.tags = append(a.tags, tag)
	tag.registerApp(a)
}

func (a *App) RemoveTag(name string) (err error) {
	a.tagsMutex.Lock()
	defer a.tagsMutex.Unlock()
	a.tags, err = removeTag(a.tags, name)
	return
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
