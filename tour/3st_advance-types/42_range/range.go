/*
Range

	for 循環的 range 形式可以迭代切片或映射，
	用法如下。
*/
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// range 用法其實就是 foreach 概念,
	// i=index, v=value
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
