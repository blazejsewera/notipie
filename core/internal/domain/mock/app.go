package mock

import "github.com/jazzsewera/notipie/core/internal/domain"

type CommandHandler struct {
	Command domain.Command
}

func (h *CommandHandler) HandleCommand(command domain.Command) {
	h.Command = command
}

func NewTestApp() domain.App {
	return domain.App{
		ID:   "1",
		Name: "TestApp",
	}
}
