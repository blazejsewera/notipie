package domain

import "time"

type Notification struct {
	Timestamp   time.Time
	Application Application
	Title       string
	Body        string
	Urgency     Urgency
}

type Urgency int

const (
	Low Urgency = iota
	Medium
	High
	Fatal
)

func (u Urgency) String() string {
	return [...]string{"Low", "Medium", "High", "Fatal"}[u]
}

type NotificationBuilder struct {
	timestamp   time.Time
	application Application
	title       string
	body        string
	urgency     Urgency
}

func NewNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (b *NotificationBuilder) WithTimestamp(timestamp time.Time) *NotificationBuilder {
	b.timestamp = timestamp
	return b
}

func (b *NotificationBuilder) WithApplication(application Application) *NotificationBuilder {
	b.application = application
	return b
}

func (b *NotificationBuilder) WithTitle(title string) *NotificationBuilder {
	b.title = title
	return b
}

func (b *NotificationBuilder) WithBody(body string) *NotificationBuilder {
	b.body = body
	return b
}

func (b *NotificationBuilder) WithUrgency(urgency Urgency) *NotificationBuilder {
	b.urgency = urgency
	return b
}

func (b *NotificationBuilder) Build() *Notification {
	var timestamp time.Time
	if b.timestamp.IsZero() {
		timestamp = time.Now()
	} else {
		timestamp = b.timestamp
	}

	return &Notification{
		Timestamp:   timestamp,
		Application: b.application,
		Title:       b.title,
		Body:        b.body,
		Urgency:     b.urgency,
	}
}
