/*
選擇值或指針作為接收者

	簡單來說，GoLang 可以利用接收者類型 (指針 or 值)，
	來區分固定接收者的方法會不會改動到調用的物件本身 (指針會，值不會)。

	另外因為 GoLang 傳值的方式為 call by value，
	意即每次調用方法時，都會複製一次該物件，再將複製的物件進行傳遞，
	若物件本身較大，就會影響較能。
	所以使用指針接收者可以避免在每次調用方法時複製該值，
	理論上程式會有更高的效能。
*/
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// 以指針作為物件, 這樣接下來都會以 call by reference 的方式傳值
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
