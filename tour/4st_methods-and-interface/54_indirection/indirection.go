/*
方法與指針重新定向

	比較上述兩個範例，就會發現到帶指針參數的函數必須接受一個指針，
	而以指針為接收者的方法被調用時，接收者既能為值也能為指針。(上述範例中，編譯警告與編譯錯誤的差別)
	這是因為使用固定接收者方法時，GoLang 會自動 v.Scale(10) 解釋為 (&v).Scale(10)。
*/
package main

import "fmt"

type Vertex struct {
	X, Y float64
}

// 指針接收者方法
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// 一般函數
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}

	//ScaleFunc(v, 10) // 編譯錯誤
	ScaleFunc(&v, 10) // OK

	v.Scale(10) // OK, 因為被解釋為 (&v).Scale(10)

	p := &v
	p.Scale(10) // OK, 因為被解釋為 (p=&v).Scale(10)

	fmt.Println(v, p)
}
