package domain

import (
	"time"
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
