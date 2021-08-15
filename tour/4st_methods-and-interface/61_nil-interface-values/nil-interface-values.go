/*
nil 接口值

nil 接口值既不保存值也不保存具體類型，

並當 nil 接口發生在代碼中且運行時，會發生錯誤。
*/
package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	// nil 接口，故會發生運行錯誤
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
