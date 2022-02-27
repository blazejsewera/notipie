package domain

import (
	"fmt"
	"time"

	"github.com/blazejsewera/notipie/core/pkg/lib/util"
)

type Notification struct {
	ID         string
	App        *App
	Timestamp  time.Time
	Title      string
	Subtitle   string
	Body       string
	Urgency    Urgency
	ExtURI     string
	ReadURI    string
	ArchiveURI string
}

func (n Notification) String() string {
	t := n.Timestamp.Format(time.RFC3339)
	var b string
	for _, line := range util.SplitLines(n.Body) {
		b += fmt.Sprintf("|> %s\n", line)
	}
	return fmt.Sprintf("[%s#%s@%s|%s] %s#%s\n%s", n.App.Name, n.App.ID, t, n.Urgency.ShortString(), n.Title, n.ID, b)
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
