/*
指針接收者

	接收者 (receiver) 可以為指針類型，
	代表著對於某類型 T，接收者類型可以用 *T 的文法。
	需要注意的是，T 類型不能是 *int 這樣的指針。

	兩者著差異為，
	指針接收者是可以改變物件內容值的方法，
	而一般接收者則否。
*/
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 使用指針接收者的方法
// 嘗試將 * 號移除後，會發現編譯警告
// 這時接收者變為值而非指針
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
