/*
append 切片

	切片可以用 append 函數來擴展，
	每次擴展時都會將欲加入的元素置於切片的末端，
	當切片對應的數組長度不夠時，
	它就會自動分配一個更大的數組來進行擴展。
*/
package main

import "fmt"

func main() {
	// nil 切片
	var s []int
	printSlice(s)

	// 添加一個元素 0
	s = append(s, 0)
	printSlice(s)

	// 再添加一個元素 1，會發現其對應的數組
	// 每次擴展時，都被重新分配大小
	s = append(s, 1)
	printSlice(s)

	// 可以一次性擴展多個元素
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
