/*
切片即是數組的引用

	切片並不儲存任何數據，
	他只是描述對應的數組的某一段元素，
	更改切片元素會變更其對應數組所對應的元素。
*/
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
	// 修改 b[0] 後，其對應的數組與其他切片引用此數組的元素都會變更
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
