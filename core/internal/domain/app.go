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
	if a.tags == nil {
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

	if emptyTags != nil {
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

type SendError struct {
	App
	Notification
	Tags []*Tag
}

func (e SendError) Error() string {
	if e.Tags == nil {
		// TODO: extract format constant
		return fmt.Sprintf("no tags for %s#%s when sending %s", e.App.Name, e.App.ID, e.Notification)
	}

	tags := "[ "
	for _, tag := range e.Tags {
		tags += fmt.Sprintf("%s ", tag.Name)
	}
	tags += "]"

	// TODO: extract format constant
	return fmt.Sprintf("tags: %s for %s#%s did not have registered users when sending %s",
		tags, e.App.Name, e.App.ID, e.Notification)
}
