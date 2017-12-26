package main

import "fmt"

/*
string底层就是一个byte的数组

string本身是不可变的，因此要改变string中字符，需要如下操作：
 */


func main()  {

	str := "hello world"
	s := []byte(str)
	s[0] = 'o'
	str = string(s)
	fmt.Println(str)

}
