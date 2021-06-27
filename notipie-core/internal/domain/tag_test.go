package domain

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TagSuite struct {
	suite.Suite
	Tag   Tag
	App1  App
	App2  App
	User1 User
	User2 User
}

func (s *TagSuite) SetupTest() {
	s.Tag = Tag{}
	s.App1 = App{}
	s.App2 = App{}
	s.User1 = User{}
	s.User2 = User{}
}

func TestTag(t *testing.T) {
	suite.Run(t, new(TagSuite))
}
