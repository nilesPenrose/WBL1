package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает
//две числовых переменных a,b, значение которых > 2^20.

func main() {
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)

	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Println("Расстояние до галактики Андромеды составляет", distance, "км.") // Выводит: Расстояние до галактики Андромеды составляет 24000000000000000000 км.

	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)
	fmt.Println("Это", days, "дня полета на скорости света.") // Выводит: Это 926568346 дня полета на скорости света.
	manyDays := new(big.Int)
	manyDays.Add(days, days)
	fmt.Println("Это", manyDays, "дня полета на скорости света туда и обратно.")

	energy := new(big.Int)
	energy.Mul(lightSpeed, lightSpeed)
	fmt.Println("Это", energy, "которой обладает один киллограмм материи.")

}
