/*
指針 (pointer)

	Go 擁有指針，指針保存了值的內存地址，
	未被賦值的指針零值為 nil。

	& 操作符會生成一個指向其操作值的指針，
	* 操作符表示指針指向的實際值。
*/
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // 創建一個指向 i 地址的指針 p
	fmt.Println(*p) // 通過指針 p, 讀取 p 地址對應的值 (i)
	*p = 21         // 通過指針 p, 設置 p 地址對應的值 (i)
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 將 j 的地址賦值給 p 指針
	*p = *p / 37   // 通過指針 p, 對 p 地址對應的值 (j), 做除法運算
	fmt.Println(j) // 查看 j 的值
}
