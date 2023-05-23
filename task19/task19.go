package main

import (
	"fmt"
	"unicode/utf8"
)

//Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
//Символы могут быть unicode.

func main() {

	someString := "WASD format123 @!#капкан 狐 jumped over the lazy 犬"
	fmt.Println(reverse(someString))
}

func reverse(someString string) string {
	r := make([]rune, len(someString))
	start := len(someString)
	for _, c := range someString {
		// quietly skip invalid UTF-8
		if c != utf8.RuneError {
			start--
			r[start] = c
		}
	}
	return string(r[start:])
}
