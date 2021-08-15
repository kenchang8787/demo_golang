/*
select

select 可以使一個 goroutine 等待多個頻道操作，
select 會阻塞道某個分支可以繼續執行為止，這時就會執行該分支。
*/
package main

import (
	"fmt"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		// 直到 quit 頻道可以取值的時候, 迴圈才會結束
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	// 此線程會等 fibonacci 將 x 值傳入頻道 c 時, 才會開始執行
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
