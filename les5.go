package main

func fibonacci(i int) (ret int) {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}

func main() {
	var N int
	N = 7
	//	m := make(map[int]int)

	println("Вычислим чему равно ", N, " число в последовательности фибоначи")
	println(N, "число в последовательности фибоначи будет равно", fibonacci(N))
}
