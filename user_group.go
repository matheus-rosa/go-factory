package go_factory

import "fmt"

type UserGroup struct {
	ID    int
	Name  string
	Users []*User
}

func (g UserGroup) String() string {
	return fmt.Sprintf("ID: %d, Name: %s", g.ID, g.Name)
}
