package domain

import (
	"fmt"
	"github.com/jazzsewera/notipie/core/pkg/lib/util"
	"time"
)

type Notification struct {
	App       *App
	Timestamp time.Time
	Title     string
	Body      string
	Urgency   Urgency
}

func (n Notification) String() string {
	t := n.Timestamp.Format(time.RFC3339)
	var b string
	for _, line := range util.SplitLines(n.Body) {
		b += fmt.Sprintf("|> %s\n", line)
	}
	return fmt.Sprintf("[%s#%s@%s|%s] %s\n%s", n.App.Name, n.App.ID, t, n.Urgency.ShortString(), n.Title, b)
}

type Urgency int

const (
	Info Urgency = iota
	Low
	Medium
	High
	Fatal
)

func (u Urgency) ShortString() string {
	return [...]string{"I", "L", "M", "H", "F"}[u]
}

func (u Urgency) String() string {
	return [...]string{"Info", "Low", "Medium", "High", "Fatal"}[u]
}
