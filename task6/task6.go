package main

import (
	"fmt"
	"sync"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

func main() {
	wg := sync.WaitGroup{}

	done := make(chan interface{})
	wg.Add(1)
	go func() {
		defer wg.Done()
		now := time.Now()
		select {
		case <-done:
			break
		}
		fmt.Printf("exit from goroutine %v\n", time.Now().Sub(now))
	}()
	time.Sleep(2 * time.Second)
	close(done)
	wg.Wait()
	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("make some work")
		return
	}()
	wg.Wait()

}
