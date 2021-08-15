/*
switch

	switch 是編寫一連串 if-else 語句的簡潔寫法。

	比較特別的是有一關鍵字 fallthrough，
	宣告在 case 內時，不會自動 break 出 switch，
	而是會繼續執行下一個 case 內的語句。

	結構如下。
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
