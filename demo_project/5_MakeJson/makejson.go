/*
編寫一個程序，提示用戶首先輸入姓名，然後輸入地址。
您的程序應該創建一個地圖，並分別使用 name 和 address 鍵將名稱和地址添加到地圖中。
您的程序應該使用 Marshal() 從地圖創建一個 JSON 對象，然後您的程序應該打印 JSON 對象。
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 創建鍵與值都為字串的地圖
	var inputMap = make(map[string]string)

	fmt.Print("請輸入姓名: ")
	inputMap["Name"] = readInput()

	fmt.Printf("請輸入地址: ")
	inputMap["Address"] = readInput()

	// 將地圖轉為 Json
	jsonData, err := json.Marshal(inputMap)

	ifErrorThenExit(err, "將用戶輸入轉為json時發生錯誤, 請再操作一遍")

	fmt.Printf("你的輸入在json格式為:\n%s", jsonData)
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
