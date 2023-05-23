package main

import "fmt"

func main() {
	var a, b int
	a = 10
	b = 20
	fmt.Printf("a = %d b = %d", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("\nswap:\na = %d b = %d", a, b)
}
