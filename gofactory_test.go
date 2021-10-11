package go_factory

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

type user struct {
	ID        uint
	Name      string
	Age       int
	Accounts  []*account
	Active    bool
	userGroup *userGroup
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type account struct {
	ID        uint
	Name      string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type userGroup struct {
	ID        uint
	Name      string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

var factory *Factory

func init() {
	factory = NewFactory(&Options{
		BaseFactory: func() map[string]interface{} {
			return map[string]interface{}{
				"user": func(factory Factory) *user {
					//factory.GetField("account")
					return &user{
						Name:      factory.String("name", "default name"),
						Age:       factory.Int("age", 20),
						Accounts:  factory.CreateN("account", factory.GetField("accounts")).([]*account),
						Active:    factory.Bool("active", true),
						userGroup: nil,
						CreatedAt: factory.Time("createdAt", time.Now()),
						UpdatedAt: factory.Time("updatedAt", time.Now()),
						DeletedAt: gorm.DeletedAt{},
					}
				},
				"account": func(factory Factory) *account {
					return &account{
						Name: factory.String("name", "default account name"),
					}
				},
			}
		},
		Gorm: nil,
	})
}

func TestFactory_Create(t *testing.T) {
	t.Run("it should...", func(t *testing.T) {
		u, ok := factory.Create("user", Fields{
			"name": "John Doe",
		}).(*user)

		assert.True(t, ok)
		assert.IsType(t, &user{}, u)
		assert.Equal(t, "John Doe", u.Name)
		assert.Equal(t, 20, u.Age)

		u, ok = factory.Create("user", Fields{
			"name": "John Doe",
			"age":  18,
			"accounts": []Fields{
				{"name": "first account"},
			},
		}).(*user)

		assert.True(t, ok)
		assert.IsType(t, &user{}, u)
		assert.Equal(t, "John Doe", u.Name)
		assert.Equal(t, 18, u.Age)
		assert.Len(t, u.Accounts, 1)

		a := u.Accounts[0]
		assert.Equal(t, "first account", a.Name)
	})
}
