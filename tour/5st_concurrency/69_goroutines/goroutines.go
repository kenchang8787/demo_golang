/*
Go 線程 (goroutine)

goroutine 是 Golang 運行時管理的輕量級線程。

go function()
會啟動一個新的 goroutine 並執行 function()
*/
package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// 宣告的方法為 go func()
	go say("world")
	say("hello")
}
