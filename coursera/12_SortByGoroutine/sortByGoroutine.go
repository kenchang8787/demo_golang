/*
編寫一個程序對一個整數數組進行排序。
程序應該將數組分成 4 部分，每部分都由不同的 goroutine 排序。
每個分區的大小應大致相等。
然後主 goroutine 應該將 4 個排序的子數組合併為一個大的排序數組。

該程序應提示用戶輸入一系列整數。
每個對數組的 1/4 進行排序的 goroutine 都應該打印它將要排序的子數組。
排序完成後，主 goroutine 應該打印整個排序列表。
*/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Print("請輸入一系列整數: ")

	input := readInput()

	divided := diviedTo4Parts(input)

	// 用 goroutince 排序每一部分
	var wg sync.WaitGroup

	for _, value := range divided {
		wg.Add(1)
		go SortIntArray(value, &wg)
	}
	wg.Wait()

	// 排序整個陣列
	wg.Add(1)
	go SortIntArray(input, &wg)
	wg.Wait()
}

// 讀取用戶輸入的字串
func readInput() []int {
	// 讀取用戶輸入
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	ifErrorThenExit(err, "在讀取用戶輸入時發生錯誤, 請再操作一遍")

	// 將字串用空白截斷組成陣列
	textArray := strings.Split(text, " ")

	// 將輸入字串統一為小寫、並略過空白鍵
	result := []int{}

	for _, t := range textArray {
		value, err := strconv.Atoi(strings.TrimSpace(t))

		ifErrorThenExit(err, "將輸入轉為整數時發生錯誤, 請再操作一遍")

		result = append(result, value)
	}

	if len(result) < 4 {
		ifErrorThenExit(errors.New(""), "輸入整數數組長度小於 4, 請再操作一遍")
	}

	return result
}

// 將輸入的陣列分為四部分
func diviedTo4Parts(intArray []int) [][]int {
	var result [][]int

	chunkSizeArray := getChunkSize(len(intArray))

	start := 0

	for i := 0; i < len(chunkSizeArray); i++ {
		end := start + chunkSizeArray[i]

		if end > len(intArray) {
			end = len(intArray)
		}

		result = append(result, intArray[start:end])

		start = end
	}

	return result
}

// 取得每一部分的大小
func getChunkSize(arrayLength int) [4]int {
	result := [4]int{arrayLength / 4, arrayLength / 4, arrayLength / 4, arrayLength / 4}

	remainder := arrayLength % 4

	switch remainder {
	case 1:
		result[3] = result[3] + 1
	case 2:
		result[3] = result[3] + 1
		result[2] = result[3]
	case 3:
		result[3] = result[3] + 1
		result[2] = result[3]
		result[1] = result[3]
	}

	return result
}

// 排序整數陣列
func SortIntArray(intArray []int, wg *sync.WaitGroup) []int {
	sort.Ints(intArray)

	fmt.Println(intArray)

	wg.Done()

	return intArray
}

// 若是發生錯誤, 就離開程序
func ifErrorThenExit(err error, errMsg string) {
	if err != nil {
		fmt.Println("ERROR: ", errMsg)
		os.Exit(0)
	}
}
