/*
編寫一個程序，提示用戶輸入整數並將整數存儲在一個排序的切片中。
程序應該寫成一個循環。 在進入循環之前，程序應該創建一個大小（長度）為 3 的空整數切片。
在每次循環過程中，程序會提示用戶輸入一個整數以添加到切片中。
該程序將整數添加到切片中，對切片進行排序，並按排序順序打印切片的內容。
切片的大小必須增長以容納用戶決定輸入的任意數量的整數。
當用戶輸入字符 X 而不是整數時，程序應該只退出（退出循環）。
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// 初始整數 slice
	slice := make([]int, 3)[:0]

	for {
		fmt.Print("請輸入一個整數或是 X 離開: ")
		input := readInput()

		// 若是輸入為 X, 則離開
		if input == "X" {
			break
		}

		// 字串轉為整數
		value, err := strconv.Atoi(input)

		ifErrorThenExit(err, "在將輸入字串轉為整數時發生錯誤, 請再操作一次")

		// 將新輸入的整數新增到陣列中, 進行排序, 並打印
		slice = append(slice, value)
		sort.Ints(slice)
		fmt.Printf("%v\n", slice)
	}
}

// 讀取用戶輸入的字串
func readInput() string {
	// 讀取用戶輸入
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	ifErrorThenExit(err, "在讀取用戶輸入時發生錯誤, 請再操作一遍")

	return strings.TrimSpace(text)
}

// 若是發生錯誤, 就離開程序
func ifErrorThenExit(err error, errMsg string) {
	if err != nil {
		fmt.Println("ERROR: ", errMsg)
		os.Exit(0)
	}
}
