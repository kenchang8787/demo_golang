/*
defer

	defer 語句會將函數推遲到外層函數返回後執行。

	推遲調用的函數其參數會立即求值,
	但直到外層函數返回前, 該函數都不會被調用。
*/
package main

import "fmt"

func main() {
	// 被推遲, 所以後執行
	defer fmt.Println("world")
	// 先執行
	fmt.Println("hello")
}
