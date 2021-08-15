/*
類型推導

	上面講到變量時有提到，
	在變量聲明時不指定其類型，
	變量的類型由右值推導得出。

	不過當右邊包含未指名類型的數值常量時，
	變量就有可能是 int、float64 或 complex128，
	這時就取決於常量的精度。
*/
package main

import "fmt"

func main() {
	v := 42 // 修改此數值的精度，即可看見不同的類型宣告
	fmt.Printf("v is of type %T\n", v)
}
