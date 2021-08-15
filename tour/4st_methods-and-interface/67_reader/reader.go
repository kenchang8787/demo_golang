/*
Reader

io 套件指定了 io.Reader 接口，它表示從數據流的末尾進行讀取，
Go 標準庫包含了該接口的許多實現，包括文件、網路連接、壓縮和加密等等。

io.Reader 接口有一個 Read 方法:

func (T) Read(b []byte) (n int, err error)

Read 用數據填充給定的字節切片並返回填充的字節數和錯誤值，
遇到數據流結尾時，會返回一個 io.EOF 的錯誤。
*/
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")
	// 每次以 8 個字節的速度讀取的切片
	b := make([]byte, 8)
	for {
		// 利用 b 作為讀取 r 的媒介
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		// 碰到數據流結尾
		if err == io.EOF {
			break
		}
	}
}
