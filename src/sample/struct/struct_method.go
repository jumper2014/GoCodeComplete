package main

import (
	"fmt"
)

type User1 struct {
	Name string
	Age  int32
	mess string
}


func (user *User1) init(name string, age int32, mess string) {
	user.Name = name
	user.Age = age
	user.mess = mess
}

func (user User1) GetName() string {
	return user.Name
}

func main() {
	var user User1
	user.init("nick", 18, "man")
	//(&user).init("nick", 18, "man")
	name := user.GetName()
	fmt.Println(name)
}