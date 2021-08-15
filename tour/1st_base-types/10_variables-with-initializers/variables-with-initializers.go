/*
變量的初始化

	變量聲明可以包含初始值，
	如果初始值已經存在，
	則可以省略類型。
*/
package main

import "fmt"

// 聲明並賦予初始值
var i, j int = 1, 2

func main() {
	// 變量沒有宣告類型, 但可以由值定義類型
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
