package main

import (
	"fmt"
	"github.com/error0702/go-study/v2/src/hello"
)

func main() {
	//fmt.Println("hello world!")
	hello.Print()
	fmt.Println(hello.Split(10))
}
