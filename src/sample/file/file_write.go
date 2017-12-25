/**
 * @File Name: writefile.go
 * @Author:
 * @Email:
 * @Create Date: 2017-12-17 12:12:09
 * @Last Modified: 2017-12-17 12:12:30
 * @Description:使用多种方式将数据写入文件
 */
package main
import (
	"fmt"
	"io/ioutil"
)

func main() {
	name := "/Users/zyt/git/github/GoCodeComplete/src/sample/file/test.txt"
	content := "Hello, xxbandy.github.io!\n"
	WriteWithIoutil(name,content)
}

//使用ioutil.WriteFile方式写入文件,是将[]byte内容写入文件,如果content字符串中没有换行符的话，默认就不会有换行符
func WriteWithIoutil(name,content string) {
	data :=  []byte(content)
	if ioutil.WriteFile(name,data,0644) == nil {
		fmt.Println("写入文件成功:",content)
	}
}
