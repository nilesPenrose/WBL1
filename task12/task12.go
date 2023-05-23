package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

type void struct {
}

func main() {
	var exist void
	data := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]void)

	for _, info := range data {
		set[info] = exist
	}

	fmt.Println(set)
}
