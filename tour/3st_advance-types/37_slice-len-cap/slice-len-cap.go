/*
切片的長度與容量

切片有 長度(length) 與 容量(capability)，

長度即是它所包含的元素個數，

容量則為從它第一個元素開始算起，到它所對應的數組的最後一個元素的個數。
*/
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 長度為 0
	s = s[:0]
	printSlice(s)

	// 拓展其長度
	s = s[:4]
	printSlice(s)

	// 截斷前兩個元素
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
