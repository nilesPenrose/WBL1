package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var (
		data = []float64{-115, -117.3, 112, 110, 1, 2, 5, 6, 9, 11, 22, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	)
	dataSet := make(map[float64][]float64)
	sort.Float64s(data)

	for _, num := range data {
		ost := num / 10
		intpart, _ := math.Modf(ost)
		dataSet[intpart*10] = append(dataSet[intpart*10], num)
	}
	fmt.Println(dataSet)
}
