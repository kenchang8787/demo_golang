/*
值為 nil 的接口

即便接口內的具體值為 nil，方法仍然會被 nil 的接收者調用，
因為保存了具體值為 nil 的接口其自身並不為 nil。
*/
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	// 處理當具體值為 nil 的狀況
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
