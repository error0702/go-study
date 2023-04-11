package main

import (
	"fmt"
	"github.com/error0702/go-study/v2/src/hello"
)

// 成员变量
var i, j int = 1, 2

func main() {
	//fmt.Println("hello world!")
	hello.Print()
	fmt.Println(hello.Split(10))
	// 定义局部变量
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
