/*
make 切片

切片可以用 make 函數來創建，
make 函數會分配一個零值數組並返回一個引用它的切片。
*/
package main

import "fmt"

func main() {
	// 長度為 5, 容量為 5
	a := make([]int, 5)
	printSlice("a", a)
	// 長度為 0, 容量為 5
	b := make([]int, 0, 5)
	printSlice("b", b)
	// 長度為 2, 容量為 5
	c := b[:2]
	printSlice("c", c)
	// 長度為 3, 容量為 3
	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
