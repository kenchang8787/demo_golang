/*
函數值

	函數也是值，在 GoLang 中函數可以作為其他函數的參數或返回值。
*/
package main

import (
	"fmt"
	"math"
)

// 將方法作為參數傳遞
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	// 定義 hypot 變數 (變數內容為方法)
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
