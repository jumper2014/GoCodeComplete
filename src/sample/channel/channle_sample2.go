/*
最后，因为定义是 make(chan int)，所以必须走goroutine方式给赋值。
 */

package channel

import "fmt"

func justequal(a int, c chan int) {
	c <- a
}

func main() {
	c := make(chan int)
	go justequal(100,c)
	go justequal(200,c)
	// c <- 100 // 这样是肯定报错的！
	x, y:= <-c, <-c  // 从 c 中接收
	fmt.Println(x,y)
}
