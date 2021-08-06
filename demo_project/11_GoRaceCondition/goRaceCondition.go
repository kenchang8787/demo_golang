/*
編寫兩個並發執行時具有競爭條件的 goroutine。 解釋競態條件是什麼以及它是如何發生的。

在這個程序中，它運行了兩個方法，add1000 和 add5000，
操作不同的全局參數 var a 和 var b，
在主函數中，我調用了 time.Sleep 讓結果最終可以打印在終端上。（或者你可以添加睡眠時間等待結果顯示)
在add1000和add5000的邏輯中，我使用 go 關鍵字讓每個循環成為不同的goroutine。
所以會導致兩個或多個goroutine試圖操作同一個地址的值，而在處理器中，a++或b++的操作如下：
1. 從a地址中取一個值。
2. 添加一個值
3. 將一個值放回地址中。
因此，goroutine 在 1、2、3 進程之間交錯運行。
所以結果是 add1000 和 add500 永遠不可能是 a = 1000 和 b = 5000，其是不可預測的。
*/

package main

import (
	"fmt"
	"time"
)

var a = 0
var b = 0

func main() {
	add1000()
	add5000()

	time.Sleep(3 * time.Second)
}

func add1000() {
	for i := 0; i < 1000; i++ {
		go func() {
			a++
		}()
	}

	fmt.Println("add 1000 end:", a)
}
func add5000() {
	for i := 0; i < 5000; i++ {
		go func() {
			b++
		}()
	}

	fmt.Println("add 5000 end:", b)
}
