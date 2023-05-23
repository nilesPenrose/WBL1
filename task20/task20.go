package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func main() {
	dataSentence := "snow dog sun"
	dataSet := strings.Split(dataSentence, " ")
	start := len(dataSet)
	var reversedDataSet []string
	for i := start - 1; i >= 0; i-- {
		reversedDataSet = append(reversedDataSet, dataSet[i])
	}

	fmt.Println(strings.Join(reversedDataSet, " "))
}
