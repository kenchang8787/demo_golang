/*
While

	初始化語句和後置語句是可選的，
	剩下條件式的 for 循環結構就是 while 的概念。
*/
package main

import "fmt"

func main() {
	sum := 1
	// 只有條件式，就是 while 的概念了
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
