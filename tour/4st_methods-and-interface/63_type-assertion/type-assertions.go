/*
類型斷言 (type-assertions)

類型斷言提供了訪問接口值底層具體值的方式如下

t := i.(T)

上述語句斷言接口值 i 保存了具體類型 T ，並將其底層類型為 T 的值賦予變量 t ，
但若接口值 i 並未保存 T 類型的值，該語句就會觸發一個恐慌。
類型斷言可以用以下方式判斷一個接口值是否保存了一個特定類型

t, ok := i.(T)

若 i 保存了一個 T 則 t 就會是其底層值，而 ok 則為 true，
反之 ok 為 false 而 t 將為 T 類型的零值。
*/
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // 報錯 (恐慌發生)
	fmt.Println(f)
}
