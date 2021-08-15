/*
Golang 使用 error 值來表示錯誤狀態，其接口為

type error interface {
    Error() string
}

通常函數會返回一個 error 值，判斷他是否為 nil 來進行錯誤處理。

i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)

error 為 nil 表示執行成功，否則失敗。
*/
package main

import (
	"fmt"
	"time"
)

// 定義一個自己的錯誤結構
type MyError struct {
	When time.Time
	What string
}

// 修正其打印出的字串樣式
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
