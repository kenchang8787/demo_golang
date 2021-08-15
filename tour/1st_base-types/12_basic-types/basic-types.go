/*
基本類型

	Go 的基本類型有
	布林值 bool
	字串 string
	整數 int int8 int16 int32 int64
	無符號整數 uint uint8 uint16 uint32 uint64
	byte (uint8 的別名)
	rune (int32 的別名, 表示一個 Unicode)
	浮點數 float32 float64
	複數 complex64 complex128

	下面例子展示了幾種類型的變量，
	與導入(import)一樣，變量也可以分組聲明。
*/
package main

import (
	"fmt"
	"math/cmplx"
)

// 分組聲明
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
