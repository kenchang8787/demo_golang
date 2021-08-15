/*
函數 (function)

	函數可以接收 0-N 個參數，
	注意參數類型會在參數名之後。
*/
package main

import "fmt"

// 函數 (x參數 x類型, y參數 y類型) 返回值類型
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
