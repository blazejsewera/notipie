package domain

import "fmt"

type App struct {
	ID           string
	Name         string
	SmallIconURL string
	BigIconURL   string
	tags         []*Tag
}

func (a *App) Send(notification Notification) error {
	if len(a.tags) == 0 {
		return SendError{
			App:          *a,
			Notification: notification,
		}
	}

	var emptyTags []*Tag

	for _, tag := range a.tags {
		err := tag.Broadcast(notification)
		if err != nil {
			emptyTags = append(emptyTags, tag)
		}
	}

	if len(emptyTags) != 0 {
		return SendError{
			App:          *a,
			Notification: notification,
			Tags:         emptyTags,
		}
	}

	return nil
}

func (a *App) AddTag(tag *Tag) {
	a.tags = append(a.tags, tag)
	tag.RegisterApp(a)
}

func (a *App) RemoveTag(tag Tag) (err error) {
	a.tags, err = removeTag(a.tags, tag)
	return
}

type SendError struct {
	App
	Notification
	Tags []*Tag
}

func (e SendError) Error() string {
	if len(e.Tags) == 0 {
		return fmt.Sprintf(noTagsWhenSendErrorFormat, e.App.Name, e.App.ID, e.Notification)
	}

	var tags []string
	for _, tag := range e.Tags {
		tags = append(tags, tag.Name)
	}

	return fmt.Sprintf(noUsersInTagsWhenSendErrorFormat, tags, e.App.Name, e.App.ID, e.Notification)
}

const (
	noTagsWhenSendErrorFormat        = "no tags for %s#%s when sending %s"
	noUsersInTagsWhenSendErrorFormat = "tags: %v for %s#%s did not have registered users when sending %s"
)
