/*
多個返回值

	函數可以返回任意數量的返回值，
	下例中 swap 函數返回了兩個字符串。
*/
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
