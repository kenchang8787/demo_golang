/*
零值

	沒有明確初始值的變量聲明會賦予它們零值。
	數值類型為 0 (int、uint、float、complex)
	布林類型為 false (bool)
	字串為 "" (string)
*/
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
