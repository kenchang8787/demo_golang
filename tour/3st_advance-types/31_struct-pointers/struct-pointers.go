/*
結構指針

	結構也可以通過結構指針來訪問。
*/
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	// 創建一個指向 v 結構地址的指針 p
	p := &v
	// 透過指針 p 訪問 v 結構的欄位
	p.X = 1e9
	fmt.Println(v)
}
