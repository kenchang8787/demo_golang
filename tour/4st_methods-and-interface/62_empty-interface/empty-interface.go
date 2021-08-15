/*
空接口

指定了零個方法的接口 interface { } 被稱為空接口，
空接口可以保存任何的類型的值 (因為每個類型都至少實現了零個方法)，
空接口可以被用來處理未知類型的值。

例如 fmt.Print 可接受類行為 interface { } 的任意數量的參數。
*/
package main

import "fmt"

func main() {
	// 定義空接口 i, 可填入任意類型
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
