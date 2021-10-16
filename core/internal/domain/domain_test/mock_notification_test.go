package domain_test

import (
	"fmt"
	"github.com/jazzsewera/notipie/core/internal/domain"
	"time"
)

func NewTestNotification(app *domain.App) domain.Notification {
	timestamp, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	return domain.Notification{
		ID:        "1",
		App:       app,
		Timestamp: timestamp,
		Title:     "Test Notification",
		Body:      "First line of body\nSecond line of body",
		Urgency:   domain.Medium,
	}
}

func New5TestNotifications() (notifications []domain.Notification) {
	for i := 1; i <= 5; i++ {
		app := &domain.App{
			ID:   fmt.Sprint(i),
			Name: fmt.Sprintf("TestApp%d", i),
		}
		timestamp, _ := time.Parse(time.RFC3339, fmt.Sprintf("2021-01-0%dT0%d:0%d:0%dZ", i, i, i, i))
		notifications = append(notifications, domain.Notification{
			ID:        fmt.Sprint(i),
			App:       app,
			Timestamp: timestamp,
			Title:     fmt.Sprintf("Test Notification #%d", i),
			Body:      fmt.Sprint(i),
			Urgency:   domain.Urgency(i - 1),
		})
	}
	return
}
