package main

import "fmt"

type void struct {
}

func main() {
	var exist void
	var interception []string

	data1 := []string{"cat", "cat", "dog", "cat", "tree", "plant", "gigachad"}
	data2 := []string{"cat", "racoon", "men", "alien", "plant", "megachad"}

	set1 := make(map[string]void)
	set2 := make(map[string]void)

	for _, info := range data1 {
		set1[info] = exist
	}

	for _, info := range data2 {
		set2[info] = exist
	}

	for key1, _ := range set1 {
		find := false

		for key2, _ := range set2 {
			if key1 == key2 {
				interception = append(interception, key1)
				find = true
				continue
			}

			if find {
				continue
			}
		}
	}

	fmt.Println(interception)
}
