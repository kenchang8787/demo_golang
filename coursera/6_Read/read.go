/*
編寫一個程序，該程序從文件中讀取信息並將其表示在一個結構片段中。
假設有一個包含一系列名稱的文本文件。
文本文件的每一行都有一個名字和一個姓氏，按這個順序，在行上用一個空格分隔。

你的程序將定義一個名稱結構，它有兩個字段，名字是 fname，姓氏是 lname。
每個字段將是一個大小為 20（字符）的字符串。

您的程序應該提示用戶輸入文本文件的名稱。
您的程序將連續讀取文本文件的每一行並創建一個結構，其中包含在文件中找到的名字和姓氏。
創建的每個結構都將添加到一個切片中，在從文件中讀取所有行之後，您的程序將有一個切片，
其中包含文件中每一行的一個結構。 從文件中讀取所有行後，您的程序應該遍歷您的結構切片並打印在每個結構中找到的名字和姓氏。
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("請輸入名稱文件路徑: ")

	filePath := readInput()

	nameArray := readFile(filePath)

	for index, name := range nameArray {
		fmt.Printf("%d. 姓: %s, 名: %s\n", index+1, name.fname, name.lname)
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

// 讀取用戶輸入的文件路徑, 並回傳 Name 陣列
func readFile(filePath string) []Name {
	// 開啟文件流
	file, err := os.Open(filePath)

	ifErrorThenExit(err, "在開啟文件時發生錯誤, 請再操作一遍")

	// 一行一行讀取文件
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	// 將讀取每行都存至陣列內
	dataArray := make([]Name, 0)

	for scanner.Scan() {
		each_line := scanner.Text()
		// 用空白鍵分為字串
		words := strings.Fields(each_line)
		// 創建 Name 物件
		obj := NewName(words[0], words[1])
		// 新增至陣列
		dataArray = append(dataArray, obj)
	}
	// 關閉文件流
	file.Close()

	return dataArray
}

// 若是發生錯誤, 就離開程序
func ifErrorThenExit(err error, errMsg string) {
	if err != nil {
		fmt.Println("ERROR: ", errMsg)
		os.Exit(0)
	}
}

// Name 結構
type Name struct {
	fname string
	lname string
}

// 創建 Name 物件
func NewName(fname string, lname string) Name {
	result := new(Name)

	maxNameLen := 20

	var runeArray []rune

	runeArray = []rune(fname)
	result.fname = string(runeArray[:maxNameLen])

	runeArray = []rune(lname)
	result.lname = string(runeArray[:maxNameLen])

	return *result
}
