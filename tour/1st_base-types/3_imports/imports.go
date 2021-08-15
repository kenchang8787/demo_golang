/*
導入 (import)

	下列代碼演示分組導入，
	但也可以編寫為多個導入語句。
	import "fmt"
	import "math"
*/
package main

// 使用分組導入
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}
