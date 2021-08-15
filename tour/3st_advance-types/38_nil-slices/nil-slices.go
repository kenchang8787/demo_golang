/*
nil 切片

	切片的零值是 nil，
	nil 切片的長度和容量皆為 0 且沒有對應的數組。
*/
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
