package mock

import "github.com/jazzsewera/notipie/core/internal/domain"

func NewTestTag() domain.Tag {
	return domain.Tag{
		Name: "TestTag",
	}
}
