package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
//
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

func UniqueString(in string) bool {
	str := strings.ToLower(in)

	check := make(map[rune]struct{})

	for _, c := range str {
		if _, ok := check[c]; ok {
			return false
		}

		check[c] = struct{}{}
	}

	return true
}

func main() {
	tests := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
		"aA",
		"AbCd",
	}

	for _, t := range tests {
		fmt.Printf("%s - %t\n", t, UniqueString(t))
	}
}
