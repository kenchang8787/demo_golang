/*
數值常數

	數值常數是高精度的值，
	一個未指定類型的常量是由上下文來決定其類型的。
*/
package main

import "fmt"

const (
	// 將 1 向左移動 100 個位數，以創建一個非常大的數字
	// 即這個數字的二進制為 1 後面加上 100 個 0
	Big = 1 << 100
	// 往右移動 99 個位數，即 Small = 1 << 1，或者說 Small = 2
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	// 由此可知，常數可以在執行階段動態的決定其類型
}
