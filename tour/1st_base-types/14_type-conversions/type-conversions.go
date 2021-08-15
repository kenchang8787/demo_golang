/*
轉型 (convert)

除了各類型相互轉型的用法要熟悉一下，
沒甚麼好特別注意的
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
