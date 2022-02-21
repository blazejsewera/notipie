package domain_test

import (
	"github.com/jazzsewera/notipie/core/internal/domain"
	"time"
)

type MockAsyncCommandHandler struct {
	Command        domain.Command
	CommandHandled chan struct{}
}

func (h *MockAsyncCommandHandler) HandleCommand(command domain.Command) {
	h.Command = command
	select {
	case h.CommandHandled <- struct{}{}:
		return
	case <-time.After(200 * time.Millisecond):
		panic("no receiver for h.CommandHandled for 200ms")
	}
}

func NewMockAsyncCommandHandler() *MockAsyncCommandHandler {
	return &MockAsyncCommandHandler{CommandHandled: make(chan struct{})}
}

func NewTestApp() (*domain.App, *MockAsyncCommandHandler) {
	commandHandler := NewMockAsyncCommandHandler()
	return domain.NewApp("1", "TestApp", "iconURI", commandHandler), commandHandler
}
