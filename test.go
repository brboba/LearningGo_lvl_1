package main

import "fmt"

func fibGen() func(int) int {
	var cache = map[int]int{}

	var fibRec func(int) int

	fibCache := func(n int) int {
		v, ok := cache[n]
		if !ok {
			v = fibRec(n)
			cache[n] = v
		}
		return v
	}

	fibRec = func(n int) int {
		fmt.Print(n, " ")
		if n <= 1 {
			return n
		}

		return fibCache(n-1) + fibCache(n-2)
	}

	return fibRec
}

func main() {
	fib := fibGen()
	/*for i := 0; i < 10; i++ {
	  fmt.Print(fib(i), ", ")
	  }*/
	fib(9)
}