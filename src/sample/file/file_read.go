package main

import "fmt"
import "io/ioutil"

func main() {
	b, e := ioutil.ReadFile("/Users/zyt/git/github/GoCodeComplete/src/sample/file/test.txt")
	if e != nil {
		fmt.Println("read file error")
		return
	}
	fmt.Println(string(b))
}