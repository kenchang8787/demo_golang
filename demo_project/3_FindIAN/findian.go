/*
編寫一個程序，提示用戶輸入一個字符串。
程序在輸入的字符串中搜索字符"i"、"a"和"n"。
程序應該打印"Found！"
如果輸入的字符串以字符"i"開頭，以字符"n"結尾，並且包含字符"a"。
程序應該打印"Not Found！"
否則。 程序不應該區分大小寫，因此字符是大寫還是小寫都沒有關係。

示例：
程序應該打印"Found！" 對於以下示例，輸入字符串"ian"、"Ian"、"iuiygaygn"、"I d skd a efju N"。
程序應該打印"Not Found！" 對於以下字符串，"ihhhhhn"、"ina"、"xian"。
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("請輸入一字符串: ")

	input := readInput()

	if findian(input) {
		fmt.Println("找到!")
	} else {
		fmt.Println("沒找到!")
	}
}

// 讀取用戶輸入的字符串
func readInput() string {
	// 讀取用戶輸入
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	ifErrorThenExit(err, "在讀取用戶輸入時發生錯誤, 請再操作一遍")
	// 轉換輸入為小寫
	result := strings.ToLower(strings.TrimSpace(text))

	return result
}

// 判斷字串是否以 i 為開頭、 n 為結尾並含有 a
func findian(input string) bool {
	if strings.Index(input, "i") == 0 &&
		strings.LastIndex(input, "n")+1 == len(input) &&
		strings.Contains(input, "a") {
		return true
	}

	return false
}

// 若是發生錯誤, 就離開程序
func ifErrorThenExit(err error, errMsg string) {
	if err != nil {
		fmt.Println("ERROR: ", errMsg)
		os.Exit(0)
	}
}
