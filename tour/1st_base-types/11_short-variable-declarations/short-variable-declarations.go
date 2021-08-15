/*
短變量聲明 (:=)

	在函數中，可以使用簡潔的賦值語句 :=，

	可以在類型明確的地方代替 var 聲明。
*/
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// 等同於 var k int = 3 或是 var k = 3
	k := 3
	// 有明確的類型則可以使用簡短聲明
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
