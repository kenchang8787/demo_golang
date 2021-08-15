/*
導出名 (export name)

	Go 套件內的公用/私有函數、成員，
	是由函數的第一個字母是否為大寫判斷的，
	此例中 fmt.Println 與 math.Pi 即是 fmt 套件的公用方法與 math 套件的公用成員。
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}
