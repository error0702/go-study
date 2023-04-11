package hello

import (
	"fmt"
	"math/cmplx"
)

/**
基本类型：
	bool

	string

	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte // uint8 的别名

	rune // int32 的别名
    	// 表示一个 Unicode 码点

	float32 float64

	complex64 complex128
*/

var (
	// 申明类型别名
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func PrintType() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

/*
*
打印各类型的零值
*/
func PrintZeroVal() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
