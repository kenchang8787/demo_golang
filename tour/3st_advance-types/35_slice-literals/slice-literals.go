/*
切片初始化語法

切片的初始化語法類似未定義長度的數組。
*/
package main

import "fmt"

func main() {
	// 創建一個切片, 同時創建一個長度為 6 的數組
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
	// 切片也可包含結構
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
