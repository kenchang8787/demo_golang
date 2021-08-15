/*
變量 (variable)

	var 語句用來聲明一個變量列表，
	跟函數的參數列表一樣，類型在最後。
*/
package main

import "fmt"

// 全域變量
var c, python, java bool

func main() {
	// 函數變量
	var i int
	fmt.Println(i, c, python, java)
}
