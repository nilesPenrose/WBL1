package main

import (
	"fmt"
	"sync"
)

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2,
//после чего данные из второго канала должны выводиться в stdout.

func main() {
	var (
		dataSet = []int{1, 2, 5, 6, 9, 11, 22, 3}
	)
	in := make(chan int)
	out := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go paste(&wg, dataSet, in)

	go modify(&wg, in, out)

	for data := range out {
		fmt.Printf("data= %d\n", data)
	}
	wg.Wait()

}

func paste(wg *sync.WaitGroup, dataSet []int, in chan int) {
	defer wg.Done()
	defer close(in)
	for _, num := range dataSet {
		in <- num
	}
	fmt.Println("closed IN\n")
}

func modify(_ *sync.WaitGroup, in <-chan int, out chan<- int) {
	defer close(out)
	for num := range in {
		out <- num * 2
	}
	fmt.Println("closed OUT")
}
