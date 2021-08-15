/*
Defer 棧

  推遲的函數調用會被儲存於一個棧中,
  當外層函數返回時, 被推遲的函數會按照後進先出的順序調用。
*/
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// 會從最後推入棧中的 i=9 開始執行到 i=0
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
