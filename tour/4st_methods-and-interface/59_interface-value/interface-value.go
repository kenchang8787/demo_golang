/*
接口值

接口也是值，在 Golang 中接口可以作為函數的參數或返回值。

變數的實作在 Golang 中，
是由一個值欄位和一個類型欄位組成的結構如下:

變數結構: {value, type}

ex1. 整數變數: {1, int32}
ex2. 字串變數: {hello, string}

而接口儲存於變數時，
接口變數的值會保存其底層類型的具體值， (下面範例為類型 T 與類型 F)
接口變數的類型則為 packageName.typeName (下面範例為 *main.T 與 main.F)
*/
package main

import (
	"fmt"
	"math"
)

// 符合 I 接口需實現 M() 方法
type I interface {
	M()
}

// 定義 T 類型
type T struct {
	S string
}

// 實作 T 類型的 M() 方法
func (t *T) M() {
	fmt.Println(t.S)
}

// 定義 F 類型
type F float64

// 實作 F 類型的 M() 方法
func (f F) M() {
	fmt.Println(f)
}

func main() {
	// 宣告變數 i 為 I 接口類型
	var i I
	// 將 T{ S: "Hello"} 的指針賦予變數 i
	i = &T{"Hello"}
	describe(i)
	i.M()
	// 將 F(math.Pi) 的值賦予變數 i
	i = F(math.Pi)
	describe(i)
	i.M()
}

// 參數為接口的函數
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
