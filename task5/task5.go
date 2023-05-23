package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

type WorkerPool[T any] struct {
	inCh   chan T //гарантирует порядок данных
	wg     sync.WaitGroup
	exitCh chan interface{}
}

func NewWorkerPool[T any](exitCh chan interface{}) *WorkerPool[T] {
	wp := &WorkerPool[T]{
		inCh:   make(chan T),
		wg:     sync.WaitGroup{},
		exitCh: exitCh,
	}

	return wp
}

func (wp *WorkerPool[T]) AddData(data T) {
	wp.inCh <- data
}

func (wp *WorkerPool[T]) RunAndWait(concurrency int, foo func(value T)) {
	for i := 0; i < concurrency; i++ {
		wp.wg.Add(1)

		go func() {
			defer wp.wg.Done()
			for {
				select {
				case <-wp.exitCh:
					return

				case value := <-wp.inCh:
					foo(value)
				}
			}
		}()
	}
	wp.wg.Wait()
	close(wp.inCh)
}

func main() {
	N, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("usage: only int number of goroutines. err: %s", err.Error())
	}
	exitCh := make(chan interface{})

	wp := NewWorkerPool[int](exitCh)
	fmt.Println(N)
	go func() {
		var i int
		timer := time.After(time.Duration(N) * time.Second)
		for {
			select {
			case <-timer:
				println("@@")
				close(exitCh)
				return
			default:
				wp.AddData(i)
				i++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	wp.RunAndWait(1, func(value int) {
		fmt.Println(value)
	})

}
