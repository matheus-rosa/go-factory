package go_factory

func FactoryRegisterer() map[string]interface{} {
	return map[string]interface{}{
		"user": userFactory,
	}
}

func userFactory(fields Fields) *User {
	return &User{
		Name:  fields.String("name", "default name"),
		Email: fields.String("email", "mail@mail.com"),
	}
}
