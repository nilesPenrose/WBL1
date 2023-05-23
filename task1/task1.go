package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	age    int
	sex    string
	height float32
	weight float32
}

func NewHuman(age int, sex string, height float32, weight float32) Human {
	return Human{
		age:    age,
		sex:    sex,
		height: height,
		weight: weight,
	}
}

func (h Human) calculateIMT() float32 {

	return h.weight / (h.height * h.height)
}

type Action struct {
	Human
}

func main() {
	act := Action{Human: NewHuman(10, "boy", 1.20, 40)}
	act.calculateIMT()
	fmt.Println(act.calculateIMT())
}
