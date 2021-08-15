/*
隱含實現

	Golang 不使用 implements 關鍵字來實現接口，
	而是通過實現一個接口 (interface) 的所有方法來實現該接口。
*/
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示 T 類型實現了 I 接口
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
