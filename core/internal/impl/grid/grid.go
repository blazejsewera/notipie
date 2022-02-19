package grid

import "github.com/jazzsewera/notipie/core/internal/domain"

type Grid interface {
	GetRootTag() *domain.Tag
}

type SimpleGrid struct {
	RootTag *domain.Tag
	Tags    []*domain.Tag
}

func NewGrid() *SimpleGrid {
	rootTag := &domain.Tag{Name: "root"}
	rootTag.Listen()
	return &SimpleGrid{RootTag: rootTag}
}

func (g *SimpleGrid) GetRootTag() *domain.Tag {
	return g.RootTag
}
