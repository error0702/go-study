package hello

import (
	"fmt"
	"math"
)

// 类型转换demo
func TypeConvert() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
