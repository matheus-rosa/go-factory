# go-factory

A simple Golang package to help to fill and create your app domain models. 
It is heavily inspired by the usage of [thoughtbot/factory_bot](https://github.com/thoughtbot/factory_bot)
and [ex_machina](https://github.com/thoughtbot/ex_machina) libraries.

## Installation

```shell
go get -u github.com/matheus-rosa/go-factory
```

## Usage

This package integrates with [GORM](https://github.com/go-gorm/gorm) in order to easily insert your models.
If you don't use GORM, probably it isn't made for you (at least if you do want to insert anything on database).

---
Let's assume that you have an `User` model in your app defined like below:

```go
package main

type User struct {
    ID        int
    Name      string
    Email     string
}
```

To use `go-factory` you'll need to inform a base factory which tells how `User` will be factored:
how your models will be created/inserted:

```go
package main

import (
	goFactory "github.com/matheus-rosa/go-factory"
)

func main() {
	goFactory.NewFactory(&goFactory.Options{ 
		// gorm.DB is optional if you don't want to insert records on database 
		Gorm: gorm.DB, 
		BaseFactory: func() map[string]interface{} {
			return map[string]interface{}{
				"user": func(factory goFactory.Factory) *User {
					return &User{
						Name:  factory.String("name", "default name"),
						Email: factory.String("email", "mail@mail.com"),
					}
				},
			}
		},
	})
}
```

For now on you can easily create by using:

```go
goFactory.Create("user", goFactory.Fields{
	"name": "John Doe",
	"email": "johndoe@email.com",
}).(*User)
```

Or if you want to insert that record on database:

```go
user := goFactory.Insert("user", goFactory.Fields{
	"name": "John Doe",
	"email": "johndoe@email.com",
}).(*User)

fmt.Println(u.ID)
```

Or even if you want to create/insert N records:

```go
users := make([]User, 3)
goFactory.CreateN("users", &users)

for _, u := users {
	fmt.Println(u)
}
```

```go
users := make([]User, 3)
goFactory.InsertN("users", &users)

for _, u := users {
	fmt.Println(u.ID)
}
```

## The Fields type

Notice we used the `goFactory.Fields` type on our `Basefactory`.
The `goFactory.Fields` can be used to correctly fill your model values, even default values,
if none were informed. The default value is always the second argument of the type you're using:

```go
factory.String("name", "this is a default name")
factory.Int("age", 29)
```

You can use all Golang's basic types. [Checkout the API](https://github.com/matheus-rosa/go-factory/blob/master/fields.go).

### What about nested models?

It's very unlikely you're only dealing with models without associations.
So let's assume our `User` we defined before have one or more `Account`:

```go
package main

type User struct {
    ID        int
    Name      string
    Email     string
    Accounts  []*Account
}

type Account struct {
    ID        int
    Name      string
}
```

And then a little refactor on our `BaseFactory`:

```go
package main

import goFactory "github.com/matheus-rosa/go-factory"

func main() {
	goFactory.NewFactory(&goFactory.Options{
		// gorm.DB is optional if you don't want to insert records on database 
		Gorm: gorm.DB, 
		BaseFactory: function() map[string]interface{} {
			return map[string]interface{}{
				"user": function(factory goFactory.Factory) *User {
					accountData := factory.GetField("account")
					accounts := make([]*Account, len(accountData))
					factory.CreateN("account", &accounts, accountData...)					
					
					return &User{
						Name:  factory.String("name", "default name"), 
						Email: factory.String("email", "mail@mail.com"), 
						Accounts: accounts,
					}
				}, 
				"account": function(factory goFactory.Factory) *Accounts {
					return &Accounts{
						Name: factory.String("name", "default account name"),
					}
				}
			}
		},
	})
}
```

For now on you can use like:

```go
package main

import goFactory "github.com/matheus-rosa/go-factory"

func main()  {
	goFactory.Create("user", goFactory.Fields{
		"name": "John Doe",
		"email": "johndoe@email.com",
		"account": []goFactory.Fields{
			{"name": "first account"},
			{"name": "second account"},
			{"name": "third account"},
		},
	}).(*User)
}
```
