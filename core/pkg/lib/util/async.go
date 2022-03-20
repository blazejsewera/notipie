package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Signal struct{}

var Ping = Signal{}

func AsyncAssert(t testing.TB, done chan Signal) *assert.Assertions {
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

func AsyncRunAtMost(t testing.TB, f func(), timeout time.Duration) {
	t.Helper()
	a := assert.New(t)
	done := AsyncRun(f)
	select {
	case <-done:
		return
	case <-time.After(timeout):
		a.FailNow(fmt.Sprintf("test blocked for over %s", timeout))
		return
	}
}

func AsyncRun(f func()) (done chan Signal) {
	done = make(chan Signal)
	go func() {
		f()
		done <- Ping
	}()
	return
}
