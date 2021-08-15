/*
接口 (interface)

接口類型是由一組方法簽名定義的集合，
接口類型的變量可以保存任何實現了這些方法的值。

*/
package main

import (
	"fmt"
	"math"
)

// 定義接口
// 符合此接口的類型為含有方法 Abs() 的類型
type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat 實現了 Abser
	a = &v // a *Vertex 實現了 Abser

	// v 是一个 Vertex（而不是 *Vertex）
	// 所以没有實現 Abser。
	// a = v

	fmt.Println(a.Abs())
}

// 定義變數
type MyFloat float64

// 定義變數方法
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
