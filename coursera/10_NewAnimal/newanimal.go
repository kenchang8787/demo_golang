/*
編寫一個程序，允許用戶創建一組動物並獲取有關這些動物的信息。
每一種動物都有一個名字，可以是牛、鳥或蛇。對於每個命令，
用戶可以創建三種類型之一的新動物，或者用戶可以請求有關他/她已經創建的動物的信息。
每隻動物都有一個由用戶定義的唯一名稱。
請注意，用戶可以定義所選類型的動物，但動物類型僅限於牛、鳥或蛇。下表包含三種動物及其相關數據。

Animal	Food	Locomotion	Sound
cow		grass	walk		moo
bird	worns 	fly			peep
snake	mice	slihter		hsss

您的程序應該向用戶顯示一個提示 > ，以指示用戶可以鍵入請求。
您的程序應該一次接受來自用戶的一個命令，打印出響應，並在新行上打印出一個新提示。
你的程序應該永遠在這個循環中繼續。
來自用戶的每個命令都必須是 newanimal 命令或 query 命令。

每個 newanimal 命令必須是包含三個字符串的一行。
第一個字符串是 newanimal 。
第二個字符串是一個任意字符串，它將是新動物的名稱。
第三個字符串是新動物的類型， 牛 、 鳥 或 蛇 。
您的程序應該通過創建新動物並打印 Created it!
來處理每個 newanimal 命令。屏幕上。

每個 查詢 命令必須是包含 3 個字符串的一行。
第一個字符串是 查詢 。
第二個字符串是動物的名字。
第三個字符串是請求的關於動物的信息的名稱， 吃 、 移動 或 說話 。
您的程序應該通過打印出請求的數據來處理每個查詢命令。

定義一個名為 Animal 的接口類型，它描述了一個動物的方法。
具體來說，Animal 接口應該包含方法 Eat()、Move() 和 Speak()，它們不帶參數也不返回值。
Eat() 方法應該打印動物的食物，Move() 方法應該打印動物的運動，Speak() 方法應該打印動物說話的聲音。
定義三種類型的牛、鳥和蛇。
對於這三種類型中的每一種，定義 Eat()、Move() 和 Speak() 方法，以便 Cow、Bird 和 Snake 類型都滿足 Animal 接口。
當用戶創建動物時，創建適當類型的對象。
當用戶發出查詢命令時，您的程序應該調用適當的方法。
*/
package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// 全局變數, 儲存動物
var _animals = make(map[string]Animal)

func main() {

	for {
		fmt.Printf(">")

		input := readInput()

		DoCommand(input)
	}
}

// 讀取用戶輸入的字串
func readInput() []string {
	// 讀取用戶輸入
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	ifErrorThenExit(err, "在讀取用戶輸入時發生錯誤, 請再操作一遍")

	// 將字串用空白截斷組成陣列
	textArray := strings.Split(text, " ")

	// 檢查長度是否為 3
	if len(textArray) != 3 {
		ifErrorThenExit(errors.New(""), "用戶輸入格式不正確, 請再操作一遍")
	}

	// 將輸入字串統一為小寫、並略過空白鍵
	result := []string{}

	for _, t := range textArray {
		t = strings.TrimSpace(t)
		t = strings.ToLower(t)
		result = append(result, t)
	}

	return result
}

// 執行用戶指令
func DoCommand(inputs []string) {
	// read first string
	switch first := inputs[0]; first {
	case "newanimal":
		NewAnimal(inputs[1], inputs[2])
	case "query":
		QueryAnimal(inputs[1], inputs[2])
	default:
		ifErrorThenExit(errors.New(""), "指令不正確, 請再操作一遍")
	}
}

// 創建動物
func NewAnimal(name, animal string) {
	succeed := false

	switch animal {
	case "cow":
		_animals[name] = &Cow{"grass", "walk", "moo"}
		succeed = true
	case "bird":
		_animals[name] = &Bird{"worms", "fly", "peep"}
		succeed = true
	case "snake":
		_animals[name] = &Snake{"mice", "slither", "hsss"}
		succeed = true
	}

	if succeed {
		fmt.Println("Create it!")
	} else {
		ifErrorThenExit(errors.New(""), "創建動物失敗, 請再操作一遍")
	}
}

// 查詢動物
func QueryAnimal(name, query string) {
	if a, exist := _animals[name]; exist {
		switch query {
		case "eat":
			a.Eat()
		case "move":
			a.Move()
		case "speak":
			a.Speak()
		default:
			ifErrorThenExit(errors.New(""), "查詢的動物存在, 但沒有此命令, 請再操作一遍")
		}
	} else {
		ifErrorThenExit(errors.New(""), "查詢的動物不存在, 請再操作一遍")
	}
}

// 若是發生錯誤, 就離開程序
func ifErrorThenExit(err error, errMsg string) {
	if err != nil {
		fmt.Println("ERROR: ", errMsg)
		os.Exit(0)
	}
}

// 動物介面
type Animal interface {
	Eat()
	Move()
	Speak()
}

// 類別牛
type Cow struct {
	food, locomotion, noise string
}

// 類別鳥
type Bird struct {
	food, locomotion, noise string
}

// 類別蛇
type Snake struct {
	food, locomotion, noise string
}

// 牛吃
func (m *Cow) Eat() {
	fmt.Println(m.food)
}

// 牛移動
func (m *Cow) Move() {
	fmt.Println(m.locomotion)
}

// 牛說話
func (m *Cow) Speak() {
	fmt.Println(m.noise)
}

// 鳥吃
func (m *Bird) Eat() {
	fmt.Println(m.food)
}

// 鳥移動
func (m *Bird) Move() {
	fmt.Println(m.locomotion)
}

// 鳥說話
func (m *Bird) Speak() {
	fmt.Println(m.noise)
}

// 蛇吃
func (m *Snake) Eat() {
	fmt.Println(m.food)
}

// 蛇移動
func (m *Snake) Move() {
	fmt.Println(m.locomotion)
}

// 蛇說話
func (m *Snake) Speak() {
	fmt.Println(m.noise)
}
