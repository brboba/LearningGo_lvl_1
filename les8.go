package main

import (
	"./les8" // странно но папка с функциями видится только в таком исполнение,
	        // вариант "learningGo/les8" не проходит
	"fmt"
)

func main() {
	greet.Hello()

	fmt.Println(greet.Shark)

	oct := greet.Octopus{
		Name:  "Jesse",
		Color: "orange",
	}

	fmt.Println(oct.String())

	oct.Reset()

	fmt.Println(oct.String())
}