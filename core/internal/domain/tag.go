package domain

type Tag struct {
	Users []*User
	Apps  []*App
}

func (t *Tag) RegisterUser(user *User) {
	t.Users = append(t.Users, user)
}
