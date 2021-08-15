/*
指針接收者 (解釋)

將上一範例轉為一般函數時，
將 * 號移除會發生編譯錯誤，
這是因為 Go 會將方法與指針重新定向。
*/
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 將 * 號移除，會發生編譯錯誤
// 這時參數變為值而非指針
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
