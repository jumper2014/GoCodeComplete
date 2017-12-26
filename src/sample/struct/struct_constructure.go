package main

import "fmt"

func NewUser(name string, age int32, mess string) *User {
	return &User{Name: name, Age: age, mess: mess}
}

func main() {
	//user := new(User)
	user := NewUser("suoning", 18, "lover")
	fmt.Println(user, user.mess, user.Name, user.Age)
}
