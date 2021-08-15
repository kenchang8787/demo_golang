/*
結構 (struct)

	一個結構就是一組字段 filed，
	字段用 . 來訪問。
*/
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
