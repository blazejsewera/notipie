package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func AsyncAssert(t testing.TB, done chan struct{}) *assert.Assertions {
	t.Helper()
	a := assert.New(t)
	select {
	case <-done:
		return a
	case <-time.After(200 * time.Millisecond):
		a.FailNow("test blocked for over 200ms")
		return a
	}
}

func AsyncRun(f func()) (done chan struct{}) {
	done = make(chan struct{})
	go func() {
		f()
		done <- struct{}{}
	}()
	return
}
