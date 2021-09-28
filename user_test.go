package go_factory

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	factory := NewFactory(&Options{
		BaseFactory: FactoryRegisterer,
	})

	fields := Fields{
		"name":      "matheus",
		"account":   Fields{"name": "first account"},
		"userGroup": Fields{"name": "that user group"},
	}

	user := factory.Create("user", fields).(*User)
	fmt.Println(user)

}
