/*Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
 По завершению программа должна выводить итоговое значение счетчика.*/
 package main

import (
	"fmt"
	//"log"
	"sync"
)
 func main() {
	c := 0
n := 200
m := sync.Mutex{}
wg := sync.WaitGroup{}
wg.Add(n)
for i := 0; i < n; i++ {
	go func(i int) {
		m.Lock()
		c++
		m.Unlock()
		wg.Done()
	}(i)
}
wg.Wait()

fmt.Println(c)

 }



 // Заморочистый вариант


// type Counter struct {
// 	num int
// 	sync.Mutex
// }

// func (c *Counter) Inc() {
// 	c.Lock()
// 	defer c.Unlock()
// 	c.num += 1
// }

// func (c *Counter) Value() int {
// 	return c.num
// }

// func main() {
// 	cnt := &Counter{
// 		num: 0,
// 	}

// 	finish := make(chan struct{})

// 	go Do(cnt, finish)

// 	select {
// 	case <-finish:
// 		log.Printf("Work done with count: %d", cnt.Value())
// 	}
// }

// func Do(cnt *Counter, finish chan struct{}) {
// 	wg := sync.WaitGroup{}

// 	// Hard work in goroutine and finish work
// 	for i := 0; i < 20; i++ {
// 		wg.Add(1)

// 		// Тяжелая работа, о заверешении которой скинется сигнал в канал
// 		go func(num int, cnt *Counter, wg *sync.WaitGroup) {
// 			defer wg.Done()

// 			fmt.Printf("Worker %d starting\n", num)
// 			cnt.Inc()
// 			fmt.Printf("Worker %d done\n", num)
// 		}(i, cnt, &wg)
// 	}

// 	wg.Wait()
// 	finish <- struct{}{}
// 	close(finish)
// }