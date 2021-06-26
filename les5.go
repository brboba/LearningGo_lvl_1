package main

import "fmt"

func fibonacci(i int) (ret int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}
m := make(map[int]int)

func fibonacci1(i int) (ret int) {
	if m[i] == 0 {
		return 0
		println(m[i])
	}
	if m[i] == 1 {
		return 1
		println(m[i])
	}
	return fibonacci1(m[i]-1) + fibonacci1(m[i]-2)
}


func main() {
	var N int
	m := make(map[int]int)

	println("Вычислим чему равно N число в последовательности фибоначи," +
		" для этого введите N")
	fmt.Scanln(&N)

	println(N, "число в последовательности фибоначи будет равно", fibonacci(N))
	println(N, "число в последовательности фибоначи будет равно", fibonacci1(N))
}
