package main

func main() {
	//не сохраняет порядок
	a := []string{"A", "B", "C", "D", "E"}
	i := 2

	a[i] = a[len(a)-1]

	a[len(a)-1] = ""
	a = a[:len(a)-1]
	//сохраняет порядок
	a = []string{"A", "B", "C", "D", "E"}
	i = 2

	copy(a[i:], a[i+1:])

	a[len(a)-1] = ""

	a = a[:len(a)-1]

}
