/*
結構初始化語法

	語法如下
*/
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // 創建一個 Vertex 結構
	v2 = Vertex{X: 1}  // v2.Y 被預設賦予為 0
	v3 = Vertex{}      // v3.X v3.Y 被預設賦予為 0
	p  = &Vertex{1, 2} // 創建一個 Vertex 結構的指針
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
