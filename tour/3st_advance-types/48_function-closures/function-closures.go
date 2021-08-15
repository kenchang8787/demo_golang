/*
函數的閉包 (真心不知道怎樣翻譯)

	一函數將函數作為返回值時，利用此函數創建函數物件後，
	在未被回收前，函數內的變量值是會被保留的，
	可以想像它是物件的私有成員的那種概念。
*/
package main

import "fmt"

// adder 函數，回傳了函數
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// 建立兩個函數物件
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			// 函數內的變量會保留於 pos 與 neg 物件中
			pos(i),
			neg(-2*i),
		)
	}
	// 當函數物件未被回收前，物件內的函數變量值依然保留
	fmt.Println(pos(1))
	fmt.Println(neg(1))
}
