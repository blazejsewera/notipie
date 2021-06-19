package domain

import (
	"time"
)

type Listener struct {
	URL        string
	LastUpdate time.Time
}

func NewListener(url string) *Listener {
	return &Listener{
		URL:        url,
		LastUpdate: time.Now(),
	}
}

func (l *Listener) UpdateTime() {
	l.LastUpdate = time.Now()
}

func (l *Listener) TimeDelta() time.Duration {
	return time.Since(l.LastUpdate)
}
