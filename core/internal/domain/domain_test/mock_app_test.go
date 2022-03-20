package domain_test

import (
	"github.com/blazejsewera/notipie/core/internal/domain"
	"github.com/blazejsewera/notipie/core/pkg/lib/util"
	"time"
)

type MockAsyncCommandHandler struct {
	Command        domain.Command
	CommandHandled chan util.Signal
}

func (h *MockAsyncCommandHandler) HandleCommand(command domain.Command) {
	h.Command = command
	select {
	case h.CommandHandled <- util.Ping:
		return
	case <-time.After(200 * time.Millisecond):
		panic("no receiver for h.CommandHandled for 200ms")
	}
}

func NewMockAsyncCommandHandler() *MockAsyncCommandHandler {
	return &MockAsyncCommandHandler{CommandHandled: make(chan util.Signal)}
}

func NewTestApp() (*domain.App, *MockAsyncCommandHandler) {
	commandHandler := NewMockAsyncCommandHandler()
	return domain.NewApp("1", "TestApp", "iconURI", commandHandler), commandHandler
}
