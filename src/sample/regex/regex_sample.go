/**
 * @File Name: test.go
 * @Author:
 * @Email:
 * @Create Date: 2017-12-24 15:12:31
 * @Last Modified: 2017-12-24 16:12:12
 * @Description:
 */

 /*
 //正则结构体
type Regexp struct {
        // contains filtered or unexported fields
}

//初始化结构体对象的方法
func Compile(expr string) (*Regexp, error)
func CompilePOSIX(expr string) (*Regexp, error)
func MustCompile(str string) *Regexp
func MustCompilePOSIX(str string) *Regexp


//结构体方法.常用的几个
//在字符串s中查找完全匹配正则表达式re的字符串.如果匹配到就停止不进行全部匹配，如果匹配不到就输出空字符串
func (re *Regexp) FindString(s string) string

//在字符串s中匹配re表达式，n表示匹配的次数，-1表示匹配整个字符串。返回字符串切片
func (re *Regexp) FindAllString(s string, n int) []string

//在src中匹配re，并替换为repl，该种方式中repl中的$符号会展开实际的变量，通常用在回溯查找中
func (re *Regexp) ReplaceAllString(src, repl string) string

//在src中匹配re，并替换为repl,该方法会按照repl中的字面意思进行替换，不支持高级变量匹配，比如回溯等等
func (re *Regexp) ReplaceAllLiteralString(src, repl string) string


//在字符串中是否匹配到re定义的字符串，匹配返回true
func (re *Regexp) MatchString(s string) bool
  */
package main

import (
"fmt"
"regexp"
)

func main() {
testString := "k8s-test-pod-12343k811sadsxsakxz-test-k8s-container.k8s-@xxbandy.github.io"
re := regexp.MustCompile("k8s?")
fmt.Println("src string:",testString)
fmt.Println("regular expression:",re)

fmt.Println("FindAllString matching all:",re.FindAllString(testString,-1))
fmt.Println("FindAllString matching twice:",re.FindAllString(testString,2))
fmt.Println("FindString:",re.FindString(testString))
fmt.Println("ReplaceAllString:",re.ReplaceAllString(testString,"biaoge"))
fmt.Println("ReplaceAllLiteralString:",re.ReplaceAllLiteralString(testString,"BIAOGE"))
fmt.Println("Match String:",re.MatchString(testString))
}