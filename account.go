package go_factory

import (
	"fmt"
	"time"
)

type Account struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func (a Account) String() string {
	return fmt.Sprintf("ID: %d, Name: %s", a.ID, a.Name)
}
