package main

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"os"
	"time"
)

func main() {
	var a, b, res float64
	var op string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)

	fmt.Print("Введите арифметическую операцию (+, -, *, /, pow): ")
	fmt.Scanln(&op)

	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		res = a / b
	case "pow":
		res = math.Pow(a, b)
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}

	fmt.Printf("Результат выполнения операции: %f\n", res)

	fmt.Printf("Задание со звездочкойвывести все простые числа до N, N задаю случайно (лень писать  :)) ) %f\n")

	rand.Seed(time.Now().Unix()) // случайные числа действительно случайные
	var N = rand.Intn(10000)     //присваиваем случайные числа
	fmt.Println("возьмем случайное число N = ", N)

	if N < 2 {
		fmt.Println("число N меньше 2")
	}

	if N == 2 {
		fmt.Println("2")
	}

	if N > 2 {
		for i := 2; (i * i) < (N + 1); i++ {
			j := big.NewInt(int64(i))
			isPrime := j.ProbablyPrime(1)
			if isPrime == true {
				fmt.Println(i)
			}
		}
	}

}
