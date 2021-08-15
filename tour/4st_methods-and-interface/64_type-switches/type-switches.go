/*
類型選擇 (type-switches)

類型選擇是一種按照順序從幾種類型斷言中選擇分支的結構如下

switch v := i.(type) {
case T:
    // v 的類型為 T
case S:
    // v 的類行為 S
default:
    // 沒有匹配，v 與 i 的類型相同
}

類型選擇的結構與 switch 相似，不過類型選擇中的 case 為類型而非值，
switch 中的聲明與類型斷言 i.(T) 的語法相同，只是具體類型被替換為關鍵字 type。
*/
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
