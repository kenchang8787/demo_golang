/*
常數 (const)

	常數的聲明與變量類似，只不過是使用 const 關鍵字，

	常數可以是字符、字符串、布林值或數值，

	常數不可以使用簡短聲明 :=。
*/
package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
