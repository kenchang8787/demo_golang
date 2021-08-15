/*
range 和 close

發送者可以透過 close() 方法來關閉頻道，
接收者可以用以下方式來判斷頻道是否被關閉，
ok 為 true 則頻道尚未關閉; 為 false 則為關閉。

value, ok := <-ch

循環 for i := range c 會不斷從頻道接收值，直到頻道被關閉。
需要注意的是，頻道只能由發送者關閉。
另外就是不像是文件流那樣，頻道通常不需要關閉，
只有在必須告訴接收者不再需要接收訊息時才有必要，例如終止一個 range 循環。
*/
package main

import (
	"fmt"
	"time"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
		time.Sleep(200 * time.Millisecond)
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
