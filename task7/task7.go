package main

import (
	"fmt"
	"sync"
)

//Реализовать конкурентную запись данных в map.
func main() {

	dataMap := make(map[int]string)
	wg := sync.WaitGroup{}
	rwMutex := sync.RWMutex{}
	N := 10

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			rwMutex.Lock()
			defer rwMutex.Unlock()
			dataMap[i] = fmt.Sprintf("goroutine %d", i)
		}(i)

	}
	wg.Wait()
	fmt.Println(dataMap)
}
