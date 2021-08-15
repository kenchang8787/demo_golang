/*
方法 (續)

	除了結構與接口外，也可以為其他類型定義方法
*/
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

// 為 MyFloat 浮點類型定義其方法
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
