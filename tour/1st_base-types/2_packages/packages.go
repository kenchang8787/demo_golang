/*
套件 (package)

	每個 Go 程式都是由 package 構成的，
	程式從 package main 的 func main 開始運行。

	下列代碼通過導入 fmt 和 math/rand 套件，
	以使用其套件內定義好的功能。

	需要注意的是，不管在使用別人的或自己的套件，
	在使用時，都會以套件的最後一個元素作為開頭，
	本例中 fmt.Println() 與 rand.Intn() 即是。
*/
// main 套件
package main

// 導入 fmt、math/rand 套件
import (
	"fmt"
	"math/rand"
)

// 程式起始點
func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
