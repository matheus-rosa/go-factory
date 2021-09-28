package go_factory

import "fmt"

type User struct {
	ID        int
	Name      string
	Email     string
	Accounts  []*Account
	UserGroup *UserGroup
}

func (u User) String() string {
	var accountsStr []string
	for _, account := range u.Accounts {
		accountsStr = append(accountsStr, account.String())
	}
	return fmt.Sprintf("ID: %d, Name: %s, Email: %s, Accounts %s, UserGroup: %s", u.ID, u.Name, u.Email, accountsStr, u.UserGroup.String())
}
