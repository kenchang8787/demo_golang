/*
方法與指針重新定向 (續)

	接受一個值作為參數的函數，則必須接受一個指定類型的值。
	而以值為接收者的方法被調用時，接收者既能為值也能為指針。
*/
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 接收者方法被調用時，接收者可為值和指針
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 函數則只能為指定類型的值
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	// 接收者為指定類型
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	// 接收者為指針
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}
