package main

//Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.

import (
	"fmt"
	"sync"
)

func main() {
	var data = []int{2, 4, 6, 8, 10, 11, 40}
	var sum int
	wg := sync.WaitGroup{}
	in := make(chan int)
	out := make(chan int)
	mutex := sync.Mutex{}
	for i := 1; i < 4; i++ {
		wg.Add(1)
		go square(in, out, i, &wg)
	}

	go func() {
		for res := range out {
			mutex.Lock()
			sum += res
			fmt.Println(res)
			fmt.Printf("sum:%d\n", sum)
			mutex.Unlock()
		}
	}()

	go func() {
		for _, val := range data {
			fmt.Printf("val = %d\n", val)
			in <- val
		}
		close(in)
	}()

	wg.Wait()
	close(out)
}
func square(in <-chan int, out chan<- int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for number := range in {
		fmt.Printf("worker %d calculate %d\n", id, number)
		out <- number * number
	}
}
