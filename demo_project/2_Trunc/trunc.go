/*
編寫一個程序，提示用戶輸入浮點數並打印整數，
該整數是輸入的浮點數的截斷版本。 截斷是刪除小數點右邊的數字的過程。
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("請輸入一個浮點數: ")
	input := readInput()
	fmt.Printf("截斷後的整數為: %d\n", int(input))
}

// 讀取用戶輸入的浮點數
func readInput() float64 {
	// 讀取用戶輸入
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	ifErrorThenExit(err, "在讀取用戶輸入時發生錯誤, 請再操作一遍")
	// 轉換輸入為浮點數
	result, err := strconv.ParseFloat(strings.TrimSpace(text), 64)

	ifErrorThenExit(err, "在將輸入轉換為浮點數時發生錯誤, 請再操作一遍")

	return result
}

// 若是發生錯誤, 就離開程序
func ifErrorThenExit(err error, errMsg string) {
	if err != nil {
		fmt.Println("ERROR: ", errMsg)
		os.Exit(0)
	}
}
