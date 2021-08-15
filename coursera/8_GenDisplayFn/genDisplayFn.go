/*
讓我們假設以下公式
位移 s 作為時間 t 的函數，加速度 a，初始速度 vo，
和初始位移如此。

s = ½ a t2 + vot + so

編寫一個程序，首先提示用戶
輸入加速度、初始速度和初始位移的值。
然後程序應提示用戶輸入時間值和
程序應計算輸入時間後的位移。

您將需要定義和使用一個函數
稱為 GenDisplaceFn() 需要三個 float64
參數、加速度 a、初始速度 vo 和初始
位移所以。 GenDisplaceFn()
應該返回一個計算位移作為時間函數的函數，
假設給定值加速度、初始速度和初始
移位。 GenDisplaceFn() 返回的函數應該帶一個 float64 參數 t，代表時間，並返回一個
float64 參數，它是時間 t 後行進的位移。

例如，假設我想假設
以下加速度、初始速度和初始值
位移：a = 10，vo = 2，so = 1。我可以使用
以下語句調用 GenDisplaceFn() 到
生成一個函數 fn，它將計算位移作為時間的函數。

fn := GenDisplaceFn(10, 2, 1)

然後我可以使用以下語句
3 秒後打印位移。

fmt.Println(fn(3))

我可以使用以下語句打印
5 秒後的位移。

fmt.Println(fn(5))
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("請輸入加速度: ")
	a := readInput()
	fmt.Print("請輸入速度: ")
	v0 := readInput()
	fmt.Print("請輸入初始位置: ")
	s0 := readInput()

	fn := GenDisplayFn(a, v0, s0)

	for {
		fmt.Print("請輸入經過秒數: ")
		t := readInput()

		fmt.Printf(
			"經過 %s 秒後, 位置在: %s\n",
			strconv.FormatFloat(t, 'f', -1, 64),
			strconv.FormatFloat(fn(t), 'f', -1, 64))
	}
}

// 讀取用戶輸入的字串
func readInput() float64 {
	// 讀取用戶輸入
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	ifErrorThenExit(err, "在讀取用戶輸入時發生錯誤, 請再操作一遍")

	result, err := strconv.ParseFloat(strings.TrimSpace(text), 64)

	ifErrorThenExit(err, "在轉換字串為浮點數時發生錯誤, 請再操作一遍")

	return result
}

// 輸入加速度、初始速度、初始位置, 回傳計算位移的公式
func GenDisplayFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		if t < 0 {
			ifErrorThenExit(errors.New(""), "時間不可以小於 0")
		}
		return a*t*t/2 + v0*t + s0
	}
}

// 若是發生錯誤, 就離開程序
func ifErrorThenExit(err error, errMsg string) {
	if err != nil {
		fmt.Println("ERROR: ", errMsg)
		os.Exit(0)
	}
}
