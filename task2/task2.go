package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из
//массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	var data = []int{2, 4, 6, 8, 10, 11, 40}

	wg := sync.WaitGroup{}
	in := make(chan int)
	out := make(chan int)

	for i := 1; i < 4; i++ {
		wg.Add(1)
		go square(in, out, i, &wg)
	}

	go func() {
		for res := range out {
			fmt.Println(res)
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
