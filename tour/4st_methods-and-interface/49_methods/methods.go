/*
方法 (method)

	GoLang 中並沒有類別，但是可以為結構 (struct) 定義方法。
	方法在這裡定義為，固定接收者 (receiver) 參數的函數。
*/
package main

import (
	"fmt"
	"math"
)

// Vertex 結構
type Vertex struct {
	X, Y float64
}

// (接收者) 方法名 (參數) (回傳值)
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	// 用 v.Abs 使用其方法
	fmt.Println(v.Abs())
}
