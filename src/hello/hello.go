package hello

import "fmt"

func Print() string {
	msg := "hello print"
	go fmt.Println(msg)
	return msg
}

func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
