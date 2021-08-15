/*
For

	Go 的循環結構只有 for 一種，
	基本的 for 循環由三個部分組成，彼此用分號隔開:

	初始化語句 - 在第一次迭代時執行
	條件表達式 - 在每次迭代前檢查
	後置語句 - 在每次迭代後執行
*/
package main

import "fmt"

func main() {
	sum := 0
	// 初始句; 條件句; 後置句
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
