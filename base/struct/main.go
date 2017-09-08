package main

import (
	"log"
	"fmt"
)

type User struct {
	Name  string
	Email string
}
type Admin struct {
	User
	Level string
}

func (u *User) Notify() error {
	log.Printf("User: Sending User Email To %s<%s>\n",
		u.Name,
		u.Email)

	return nil
}
func (u *User) Show() error {
	fmt.Println("打印Name:" + u.Name)

	return nil
}

type Notifier interface {
	Notify() error
	Show() error
}

func SendNotification(notify Notifier) error {
	return notify.Notify()
}

func main() {
	user := &User{
		Name:  "AriesDevil",
		Email: "ariesdevil@xxoo.com",
	}

	SendNotification(user)
	fmt.Println("-----------------------")
	user.Notify()

	fmt.Println("-----------------------")
	admin := &Admin{
		User: User{
			Name:  "AriesDevil",
			Email: "ariesdevil@xxoo.com",
		},
		Level: "super",
	}

	SendNotification(admin)
}
