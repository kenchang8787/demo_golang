/*
修改映射

	插入或修改元素
	map[key] = element

	獲取元素
	element = map[key]

	刪除元素
	delete(map, key)

	檢查鍵值是否存在
	element, ok = map[key]

	若 key 在 map 中，ok 為 true ; 否則 ok 為 false，
	若 key 不在 map 中，則 element 為其類型的零值。
*/
package main

import "fmt"

func main() {
	m := make(map[string]int)
	// 插入鍵為 "Answer", 值為 42 的映射
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])
	// 修改 "Answer" 鍵的值為 48
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	// 刪除 "Answer" 鍵
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])
	// 檢查 "Answer" 鍵是否存在
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
