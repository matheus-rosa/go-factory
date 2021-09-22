package go_factory

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	factory := NewFactory(&Options{
		BaseFactory: FactoryRegisterer,
	})

	user := factory.Create("user", Fields{
		"name": "matheus",
	}).(*User)

	fmt.Println(user)
}
