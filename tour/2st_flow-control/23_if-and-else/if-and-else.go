/*
if else

	在 if 簡短語句終生名的變量可以在任何對應的 else 中使用。
*/
package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	// 宣告變量 v
	if v := math.Pow(x, n); v < lim {
		// 使用變量 v
		return v
	} else {
		// 使用變量 v
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 不能使用變量 v
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
