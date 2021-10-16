package domain_test

import "github.com/jazzsewera/notipie/core/internal/domain"

type MockCommandHandler struct {
	Command domain.Command
}

func (h *MockCommandHandler) HandleCommand(command domain.Command) {
	h.Command = command
}

func NewTestApp() domain.App {
	return domain.App{
		ID:   "1",
		Name: "TestApp",
	}
}
