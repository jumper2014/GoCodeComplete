package main

import "fmt"

type User struct {
	Name string
	Age  int32
	mess string
}


func main() {
	var user User
	user.Name = "nick"
	user.Age = 18
	user.mess = "lover"

	var user1 *User = &User{
		Name: "dawn",
		Age:  21,
	}
	fmt.Println(*user1)                    //{dawn 21 }
	fmt.Println(user1.Name, (*user1).Name) //dawn dawn

	var user2 *User = new(User)
	user2.Name = "suoning"
	user2.Age = 18
	fmt.Println(user2)                     //&{suoning 18 }
	fmt.Println(user2.Name, (*user2).Name) //suoning suoning
}