package main

import "fmt"

func main() {
	map1 := make(map[string]interface{})
	map1["hello"] = "world"
	map1["age"] = 10
	fmt.Print(map1)
}
