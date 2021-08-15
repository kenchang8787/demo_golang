/*
命名返回值

	Go 的返回值可以被命名，
	它們會被視作定義在函數頂部的變量。

	返回值的名稱應該要具有一定意義，
	它可以做為文檔使用。
*/
package main

import "fmt"

// 返回值被命名為 x, y
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// 當命名返回值時，直接使用 return 即可返回
	return
}

func main() {
	fmt.Println(split(17))
}
