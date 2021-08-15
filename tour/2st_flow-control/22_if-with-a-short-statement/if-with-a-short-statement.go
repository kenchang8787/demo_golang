/*
if 的簡短語句

	跟 for 一樣，if 可以在條件式之前執行一個簡單的語句，
	該語句聲明的變量的作用域僅在 if 內。
*/
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// 簡單語句, v 的作用域僅在 if 內
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
