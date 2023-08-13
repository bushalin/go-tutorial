package main

import "fmt"

func main() {
	fmt.Println("methods in golang")

	firstUser := User{"hasib", "hasib@go.dev", true, 15}
	fmt.Println(firstUser.GetName(), firstUser.GetEmail(), firstUser.GetStatus(), firstUser.GetAge())
	firstUser.NewMail()
	fmt.Println(firstUser.GetEmail())
}

type User struct {
	name   string
	email  string
	status bool
	age    int
}

func (u User) GetName() string {
	return u.name
}
func (u User) GetEmail() string {
	return u.email
}
func (u User) GetStatus() bool {
	return u.status
}
func (u User) GetAge() int {
	return u.age
}

// this is a pass by value type
func (u User) NewMail() {
	u.email = "test@go.dev"
	fmt.Println("email of this user is: ", u.email)
}
