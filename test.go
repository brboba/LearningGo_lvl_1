package main

import (
	"fmt"
<<<<<<< Updated upstream
	"io"
	"log"
	"os"
)

func main() {

	// Пробуем открыть файл
	file, err := os.Open("les9/conf.yaml")
	if err != nil {
		log.Fatalf("Не могу открыть файл: %v", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Не могу закрыть файл: %v", err)
		}
	}()

	ConfRead:= make([]byte, 64)

	for{
		n, err := file.Read(ConfRead)
		if err == io.EOF{   // если конец файла
			break           // выходим из цикла
		}
		fmt.Print(string(ConfRead[:n]))
=======
	"math/big"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix()) // случайные числа действительно случайные
	var N = rand.Intn(10000)       //присваиваем случайные числа
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
			if isPrime == true{
				fmt.Println(i)
			}
		}
>>>>>>> Stashed changes
	}
}