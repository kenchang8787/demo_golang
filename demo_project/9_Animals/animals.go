/*
編寫一個程序，允許用戶獲取有關一組預定義動物的信息。
預定義了三種動物：牛、鳥和蛇。
每隻動物都能吃、能動、能說話。
用戶可以發出請求以找出有關動物的三件事之一：
1）牠吃的食物，
2）它的運動方式，
3）它說話時發出的聲音。
下表包含三種動物及其相關數據，這些數據應該硬編碼到您的程序中。

Animal	Food	Locomotion	Sound
cow		grass	walk		moo
bird	worns 	fly			peep
snake	mice	slihter		hsss

您的程序應該向用戶顯示一個提示 > ，以指示用戶可以鍵入請求。
您的程序一次接受來自用戶的一個請求，打印出請求的答案，並打印出一個新的提示。
你的程序應該永遠在這個循環中繼續。
來自用戶的每個請求都必須是包含 2 個字符串的一行。
第一個字符串是動物的名稱，可以是 cow 、 bird 或 snake 。
第二個字符串是請求的關於動物的信息的名稱， 吃 、 移動 或 說話 。
您的程序應該通過打印出請求的數據來處理每個請求。

您將需要一個數據結構來保存有關每隻動物的信息。
創建一個名為 Animal 的類型，它是一個包含三個字段的結構體：food、locomotion 和 noise，
所有這些字段都是字符串。
製作三個方法，稱為 Eat()、Move() 和 Speak()。
所有方法的接收器類型都應該是你的 Animal 類型。
Eat() 方法應該打印動物的食物，
Move() 方法應該打印動物的運動，
Speak() 方法應該打印動物說話的聲音。當用戶發出請求時，您的程序應該調用適當的方法。
*/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// 創建三種動物的結構
var cow = Animal{"grass", "walk", "moo"}
var bird = Animal{"worms", "fly", "peep"}
var snake = Animal{"mice", "slither", "hsss"}

func main() {

	for {
		fmt.Printf(">")

		inputArray := readInput()

		animal := strings.ToLower(strings.TrimSpace(inputArray[0]))
		action := strings.ToLower(strings.TrimSpace(inputArray[1]))

		animalAction(animal, action)
	}
}

// 動物結構
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// 動物結構方法: 吃
func (a *Animal) Eat() {
	fmt.Println(a.food)
}

// 動物結構方法: 移動
func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

// 動物結構方法: 講話
func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

// 動物動作
func animalAction(animal, action string) {
	if animal == "cow" {
		switch action {
		case "eat":
			cow.Eat()
		case "move":
			cow.Move()
		case "speak":
			cow.Speak()
		default:
			fmt.Printf("ERROR: did not contains the action: %s\n", action)
		}
	} else if animal == "bird" {
		switch action {
		case "eat":
			bird.Eat()
		case "move":
			bird.Move()
		case "speak":
			bird.Speak()
		default:
			fmt.Printf("ERROR: did not contains the action: %s\n", action)
		}
	} else if animal == "snake" {
		switch action {
		case "eat":
			snake.Eat()
		case "move":
			snake.Move()
		case "speak":
			snake.Speak()
		default:
			fmt.Printf("ERROR: did not contains the action: %s\n", action)
		}
	} else {
		fmt.Printf("ERROR: did not contains the animal: %s\n", animal)
	}
}

// 讀取用戶輸入的字串
func readInput() []string {
	// 讀取用戶輸入
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	ifErrorThenExit(err, "在讀取用戶輸入時發生錯誤, 請再操作一遍")

	// 將字串用空白截斷組成陣列
	result := strings.Split(text, " ")

	// 檢查長度是否為 2
	if len(result) != 2 {
		ifErrorThenExit(errors.New(""), "用戶輸入格式不正確, 請再操作一遍")
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
