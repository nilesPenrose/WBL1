package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0

func main() {

	number, err := strconv.ParseInt((os.Args[1]), 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	if number > math.MaxInt64 || number < math.MinInt64 {
		log.Fatalf("so big or so small.. usage")
	}

	pos, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
	if pos > 63 || pos < 0 {
		log.Fatalf("usage: position must be <= 63 and >= 0")
	}

	bit, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatal(err)
	}

	if bit != 0 && bit != 1 {
		log.Fatalf("usage: bit must be 1 or 0")
	}

	if bit == 1 {
		var mask int64 = 1 << pos // 001 -> 100
		var res = number | mask
		fmt.Println(res)
		return
	}

	var mask int64 = 1 << pos
	var res = (number &^ mask)
	fmt.Println(res)
	return

}
