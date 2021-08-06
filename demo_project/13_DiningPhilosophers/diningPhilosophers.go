/*
通過以下約束/修改來實現哲學家就餐問題。

應該有 5 個哲學家共用筷子，每對相鄰的哲學家之間有一根筷子。
每個哲學家應該只吃 3 次（不像我們在講座中那樣無限循環）
哲學家以任何順序拿起筷子，而不是從最小的數字開始（我們在講座中就是這樣做的）。
為了吃飯，哲學家必須獲得在其自己的 goroutine 中執行的主機的許可。
主人允許不超過 2 個哲學家同時進餐。
每個哲學家都有編號，從 1 到 5。
當哲學家開始吃東西時（在獲得必要的鎖之後），它會在一行上單獨打印 開始吃 <number> ，其中 <number> 是哲學家的編號。
當哲學家吃完（在它釋放鎖之前）它會在一行上打印 finishingating <number> ，其中 <number> 是哲學家的編號。


	說明:
	在 main 中, 我創建了
	1. wait group, 用於等待所有 goroutine 執行完畢, 才結束程式
	2. chopsticks, 初始化筷子陣列 (內含互斥鎖)
	3. philosophers, 初始化哲學家陣列 (內含編號與左、右兩根筷子對應的值)
	4. c, 創建了一個大小為 2 的 channel, 用於限制同時用餐的哲學家只有兩位
	5. for loop, 創建哲學家數量的 goroutine

	在 Philo.eat 中,
	1. for loop, 每位哲學家用餐 3 次
	2. c <- p.number, 將哲學家的編號填入 channel, 若是 channel 已滿 (大小為 2 個 int), 當前的 goroutine 會等待 channel 空出位置
	3. chopstick.Lock, 將當前編號哲學家兩邊的筷子上鎖
	4. time.Sleep, 用於觀察當前是否只有兩位或以下的哲學家正在用餐
	5. chopstick.unLock, 將用餐完的哲學家兩邊的筷子解鎖
	6. <- c, 將 channel 的位置空出, 讓下一位哲學家用餐
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("##### START #####")

	var wg sync.WaitGroup

	chopsticks := createChopstickArray()
	philosophers := createPhilosopherArray(chopsticks)

	c := make(chan int, 2)

	for _, philo := range philosophers {
		wg.Add(1)
		go philo.eat(&wg, c)
	}

	wg.Wait()

	fmt.Println("##### FINISH #####")
}

// 創建筷子陣列
func createChopstickArray() []*Chopstick {
	result := make([]*Chopstick, 5)

	for i := 0; i < 5; i++ {
		result[i] = new(Chopstick)
	}

	return result
}

// 創建哲學家陣列, 並設定哲學家的編號與對應到此哲學家的左、右筷子值
func createPhilosopherArray(chopsticks []*Chopstick) []*Philo {
	result := make([]*Philo, 5)

	for i := 0; i < 5; i++ {
		result[i] = &Philo{i, chopsticks[i], chopsticks[(i+1)%5]}
	}

	return result
}

// 一根筷子
type Chopstick struct {
	sync.Mutex
}

// 一位哲學家
type Philo struct {
	number                        int
	leftChopstick, rightChopstick *Chopstick
}

func (p Philo) eat(wg *sync.WaitGroup, c chan int) {
	for i := 0; i < 3; i++ {

		c <- p.number

		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		fmt.Printf("starting to eat <%d> --%d\n", p.number, i+1)
		time.Sleep(3000 * time.Millisecond)
		fmt.Printf("finishing eating <%d> --%d\n", p.number, i+1)

		p.leftChopstick.Unlock()
		p.rightChopstick.Unlock()

		<-c
	}

	wg.Done()
}
