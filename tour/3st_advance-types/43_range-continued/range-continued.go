/*
Range (續)

可以將下標或值賦予 _ 來忽略它。
*/
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	// 不需要 index, 故直接省略
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	// 未使用 index, 則用 _ 忽略
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
