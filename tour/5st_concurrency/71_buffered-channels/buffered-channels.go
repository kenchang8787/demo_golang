/*
緩衝頻道 (buffered-channel)

頻道是可以帶緩衝的，創建緩衝頻道如下

ch := make(chan int, 100)

當頻道的緩衝區域填滿時 (上述頻道可以儲存 100 個整數值)，
再向其發送數據時才會堵塞 (block); 而當緩衝區為空時，接收方則會堵塞。
*/
package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3 // 這裡會發生堵塞
	ch <- 4
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
