/*
Switch
	在 Golang 中,
	。switch-case 不須使用 break,
	。switch 後面可以式表達式,
	。case 後面可以是表達式 (例如 x < 1),
	其他部分跟認知中的 switch 用法一樣
*/
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 基本範例
	fmt.Println("\n# 基本範例")

	x := 5

	switch x {
	case 1:
		fmt.Println("x=1")
	case 2:
		fmt.Println("x=2")
	case 3:
		fmt.Println("x=3")
	default:
		fmt.Println("x != 1, 2, 3")
	}

	// switch 表達式範例
	fmt.Println("\n# switch 表達式範例")

	switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Windows 作業系統")
	case "linux":
		fmt.Println("Linux 作業系統")
	default:
		fmt.Printf("其他作業系統: %s\n", os)
	}

	// case 表達式範例
	fmt.Println("\n# case 表達式範例")

	num := 1
	switch start := 0; num {
	case start:
		fmt.Println("0")
	case start + 1:
		fmt.Println("1")
	default:
		fmt.Println("other")
	}

	// switch 後無變數範例
	fmt.Println("\n# switch 後無變數範例")

	currentHour := time.Now().Hour()
	switch {
	case currentHour < 12:
		fmt.Println("Good morning!")
	case currentHour < 17:
		fmt.Println("Good afternoon.")
	case currentHour < 24:
		fmt.Println("Good eventing")
	default:
		fmt.Println("Your time is messy")
	}
}
