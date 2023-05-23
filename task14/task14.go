package main

//Разработать программу, которая в рантайме способна определить тип переменной:
//int, string, bool, channel из переменной типа interface{}.

func main() {
	var (
		num  = 4
		flag = false
		addr = "address"
	)
	ch := make(chan int)

	checkType(num)
	checkType(flag)
	checkType(addr)
	checkType(ch)
	close(ch)
}

func checkType(i interface{}) {
	switch i.(type) {
	case int:
		println("is integer")

	case string:
		println("is string")

	case bool:
		println("is bool")

	case chan int:
		println("is chan int")

	default:
		println("has unknown type")
	}
}
