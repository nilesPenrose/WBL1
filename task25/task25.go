package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	// текущий поток заблокируется, пока канал, возвращенный из time.After не вернет свое значение.
	// Это произойдет только после d времени
	<-time.After(d)
}

func main() {
	now := time.Now()

	Sleep(3 * time.Second)

	fmt.Printf("ended after %v\n", time.Now().Sub(now))

	// поспим еще 3 секунды с использованием функции из библиотеки time для сравнения
	now2 := time.Now()
	time.Sleep(3 * time.Second)
	fmt.Printf("ended after %v\n", time.Now().Sub(now2))
}
