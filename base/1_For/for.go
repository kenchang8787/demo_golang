/*
For
	跟一般的 for 迴圈語法結構差不多，結構如下
	起始句; 條件句; 迭代句	(i:=0;	i<10;	i++)

While
	Golang 的 while 其實就是將起始句與迭代句省略的 for

無窮迴圈
	就把所有 for 條件式都拿掉就可以了

Foreach
	Golang 裡面的 foreach 比較特別,
	要配合 range 關鍵字來執行, 結構如下
	index, element := range someSlice (項次, 迭代的值 := 關鍵字 你的陣列)
*/

package main

import "fmt"

func main() {
	forExample()
	whileExample()
	infiniteExample()
	foreachExample()
}

// For 範例
func forExample() {
	fmt.Println("\n# For 迴圈結構範例")

	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
	}
}

// While 範例
func whileExample() {
	// While
	fmt.Println("\n\n# While 範例")

	index := 0
	for index < 3 {
		fmt.Printf("%d ", index)
		index++
	}
}

// 無窮迴圈範例 (展示跑10次)
func infiniteExample() {
	fmt.Println("\n\n# 無窮迴圈範例(隨便跑幾次就好了)")

	index := 0
	for {
		fmt.Printf("%d ", index)
		index++

		if index == 10 {
			break
		}
	}
}

// Foreach 範例
func foreachExample() {
	fmt.Println("\n\n# Foreach 範例")

	nameList := []string{"大哥", "二哥", "三哥", "四哥", "五哥"}

	for index, name := range nameList {
		fmt.Printf("%d.%s ", index, name)
	}
}
