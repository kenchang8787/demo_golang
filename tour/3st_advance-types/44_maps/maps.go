/*
映射 (map)

	概念上就跟字典一樣，映射的零值為 nil。
	nil 映射既沒有鍵值，也不能添加鍵值。
*/
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// 鍵為 "字串", 值為 Vertex 結構的映射
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
