package go_factory

import "time"

func FactoryRegisterer() map[string]interface{} {
	return map[string]interface{}{
		"user":      userFactory,
		"account":   accountFactory,
		"userGroup": userGroupFactory,
	}
}

func userFactory(factory Factory) *User {
	return &User{
		Name:  factory.String("name", "default name"),
		Email: factory.String("email", "mail@mail.com"),
		Accounts: []*Account{
			factory.Create("account", factory.GetField("account")).(*Account),
		},
		UserGroup: factory.Create("userGroup", factory.GetField("userGroup")).(*UserGroup),
	}
}

func accountFactory(factory Factory) *Account {
	return &Account{
		Name:      factory.String("name", "default account name"),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
		DeletedAt: nil,
	}
}

func userGroupFactory(factory Factory) *UserGroup {
	return &UserGroup{
		Name: factory.String("name", "Group A"),
	}
}
