/*
切片的切片

	切片可包含任何類型，也可以包含其他的切片。
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 創一個 9 宮格 (圈圈叉叉遊戲)
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}

	// 兩個玩家分別輸入 O 和 X
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
