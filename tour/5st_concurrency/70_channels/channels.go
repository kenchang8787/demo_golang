/*
頻道 (chennel)

頻道是帶有類型的通道，可以透過頻道操作符 <- 來發送或接收值如以下

ch <- v // 將 v 發送至 ch 頻道
v := <- ch // 從 ch 頻道接收值並賦予 v

箭頭方向就是數據流的方向，
和映射 (map) 與切片 (slice) 一樣，
頻道使用 make 創建如下

ch := make(chan int)

預設情況下，頻道的發送和接收操作在另一端準備好之前都會阻塞 (block)，
這使得 goroutine 可以在沒有使用互斥鎖 (mutex) 或有競爭條件變量 (race condition) 的情況下進行同步。
*/
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 將 sum 送入頻道 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 從頻道 c 接收並賦值給 x, y

	fmt.Println(x, y, x+y)
}
