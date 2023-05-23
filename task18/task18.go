package main

import (
	"fmt"
	"sync"
	"time"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

type Counter struct {
	num   int
	mutex sync.RWMutex
}

func NewCounter() *Counter {
	return &Counter{
		num:   0,
		mutex: sync.RWMutex{},
	}
}
func (c *Counter) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.num++
}

func (c *Counter) Value() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.num

}

func main() {
	fmt.Println("started work")
	cnt := NewCounter()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(cnt, i, &wg)
	}
	wg.Wait()
	fmt.Println(cnt.Value())
}

func worker(cnt *Counter, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("i=%d\n", i)
	time.Sleep(1 * time.Second)
	cnt.Inc()

}
