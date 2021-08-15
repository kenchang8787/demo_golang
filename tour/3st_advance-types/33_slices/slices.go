/*
切片 (slice)

	每個數組的大小都是固定的，
	而切片則為數組提供動態大小、靈活的視角。

	切片通過兩個下標來界定檢視的範圍。
	a[ low : high ]
	以下切片的表達式包含了數組 a 的 a[1]、a[2]、a[3] 共三個元素。
	a[1 : 4]
*/
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
