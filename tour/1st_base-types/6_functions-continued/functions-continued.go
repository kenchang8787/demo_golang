/*
函數 (續)

	當參數的類型相同時，
	除了最後一個參數，
	其他參數類型皆可省略。
*/
package main

import "fmt"

// 將 add(x int, y int) 省略為以下
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
