package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"sync"
	"syscall"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые
//читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

func main() {
	sigQuit := make(chan os.Signal, 1)
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)

	runtime.GOMAXPROCS(0)
	data := os.Args

	N, err := strconv.Atoi(data[1])
	if err != nil {
		log.Fatalf("usage: only int number of goroutines. err: %s", err.Error())
	}

	in := make(chan interface{})
	wg := sync.WaitGroup{}
	for i := 0; i < N; i++ {
		wg.Add(1)
		go worker(&wg, in)
	}

	func() {
		for {
			select {
			case <-sigQuit:
				close(in)
				return
			default:
				in <- "qeqw1"
				time.Sleep(1 * time.Second)
			}
		}
	}()
	wg.Wait()
	fmt.Println("exit")
}

func worker(wg *sync.WaitGroup, in chan interface{}) {
	for data := range in {
		fmt.Println(data)
	}
	wg.Done()
}
