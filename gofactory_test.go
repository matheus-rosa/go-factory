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
					accountData := factory.GetField("account")
					accounts := make([]*account, len(accountData))
					factory.CreateN("account", &accounts, accountData...)

					return &user{
						Name:      factory.String("name", "default name"),
						Age:       factory.Int("age", 20),
						Accounts:  accounts,
						Active:    factory.Bool("active", true),
						userGroup: factory.Create("userGroup", factory.GetField("userGroup")...).(*userGroup),
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
				"userGroup": func(factory Factory) *userGroup {
					return &userGroup{
						Name:      factory.String("name", "default group name"),
						Active:    factory.Bool("active", true),
						CreatedAt: factory.Time("createdAt", time.Now()),
						UpdatedAt: factory.Time("updatedAt", time.Now()),
						DeletedAt: gorm.DeletedAt{},
					}
				},
			}
		},
		Gorm: nil,
	})
}

func TestFactory_Create(t *testing.T) {
	t.Run("it should fabricate data correctly", func(t *testing.T) {
		u, ok := factory.Create("user", Fields{
			"name": "John Doe",
		}).(*user)

		assert.True(t, ok)
		assert.IsType(t, &user{}, u)
		assert.Equal(t, "John Doe", u.Name)
		assert.Equal(t, 20, u.Age)

		layout := "2006-01-02T15:04:05.000Z"
		userGroupCreatedAt, _ := time.Parse(layout, "2021-10-10T11:45:26.371Z")
		userGroupUpdatedAt, _ := time.Parse(layout, "2021-10-11T11:45:26.371Z")

		u, ok = factory.Create("user", Fields{
			"name": "Albus Dumbledore",
			"age":  100,
			"account": []Fields{
				{"name": "first account"},
				{"name": "second account"},
				{"name": "third account"},
			},
			"userGroup": Fields{
				"name":      "main user group",
				"createdAt": userGroupCreatedAt,
				"updatedAt": userGroupUpdatedAt,
			},
		}).(*user)

		assert.True(t, ok)
		assert.IsType(t, &user{}, u)
		assert.Equal(t, "Albus Dumbledore", u.Name)
		assert.Equal(t, 100, u.Age)
		assert.Len(t, u.Accounts, 3)

		a := u.Accounts[0]
		assert.Equal(t, "first account", a.Name)

		assert.Equal(t, "main user group", u.userGroup.Name)
		assert.True(t, userGroupCreatedAt.Equal(u.userGroup.CreatedAt))
		assert.True(t, userGroupUpdatedAt.Equal(u.userGroup.UpdatedAt))
	})
}
