/*
用 Go 編寫冒泡排序程序。該程序
應提示用戶輸入最多 10 個整數的序列。該程序
應該在一行上打印整數，按排序順序，從最小到
最大的。使用您最喜歡的搜索工具查找有關泡沫如何產生的描述
排序算法有效。

作為這個程序的一部分，你應該寫一個
名為 BubbleSort() 的函數
將一片整數作為參數，不返回任何內容。 BubbleSort() 函數應該修改切片，以便元素排序
命令。

冒泡排序算法中的一個循環操作是
Swap 操作，它交換兩個相鄰元素的位置
片。您應該編寫一個執行此操作的 Swap() 函數。你的掉期（）
函數應該接受兩個參數，一個整數切片和一個索引值 i
表示切片中的位置。 Swap() 函數應該什麼都不返回，但它應該交換
位置 i 中切片的內容與位置 i+1 中的內容。
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
	fmt.Print("請輸入一系列整數, 用空白作為間隔, 最多十個: ")

	intArray := readInput()

	// 檢查長度
	if len(intArray) > 10 {
		fmt.Println("整數的數量不可超過 10 個")
	} else {
		bubleSort(intArray)
		fmt.Printf("%v\n", intArray)
	}
}

// 讀取用戶輸入的字串轉為整數陣列
func readInput() []int {
	// 讀取用戶輸入
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	ifErrorThenExit(err, "在讀取用戶輸入時發生錯誤, 請再操作一遍")

	// 用空白作為間隔, 轉換為字串陣列
	textArray := strings.Split(text, " ")

	// 將字串陣列轉為整數陣列
	result := []int{}

	for _, t := range textArray {
		integer, err := strconv.Atoi(strings.TrimSpace(t))

		ifErrorThenExit(err, "將字串轉換為整數時發生錯誤, 請重新操作")

		result = append(result, integer)
	}

	return result
}

// 若是發生錯誤, 就離開程序
func ifErrorThenExit(err error, errMsg string) {
	if err != nil {
		fmt.Println("ERROR: ", errMsg)
		os.Exit(0)
	}
}

// 將整數陣列進行冒泡排序
func bubleSort(input []int) {
	// i = run times
	for i := len(input); i > 0; i-- {
		// j = index of input
		for j := 0; j < i-1; j++ {
			if input[j] > input[j+1] {
				swap(input, j)
			}
		}
	}
}

// 交換
func swap(input []int, index int) {
	temp := input[index]
	input[index] = input[index+1]
	input[index+1] = temp
}
