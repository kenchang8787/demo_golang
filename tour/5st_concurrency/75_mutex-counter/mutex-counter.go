/*
sync.Mutex

Golang 標準庫中提供了 sync.Mutex 互斥鎖類型。
*/
package main

import (
	"fmt"
	"sync"
)

// SafeCounter 的結構內含一個互斥鎖
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc 增加给定 key 的计数器的值。
func (c *SafeCounter) Inc(key string) {
	// 互斥鎖 Lock() 後，同時間只有一個 goroutine 可以訪問
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

// Value 返回给定 key 的计数器的当前值。
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	// 創建 1000 個 goroutine, 新增至同一個 safe counter
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	fmt.Println(c.Value("somekey"))
}
