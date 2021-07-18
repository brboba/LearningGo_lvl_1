package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
<<<<<<< Updated upstream
	var str string
	fmt.Println("Введите последовательность целых чисел в формате: \"7,2,47,242,475,12,42\", если пропустить" +
		"ввод то будет использована последовательность из примера ")
	fmt.Scanln(&str)

	if len(str) == 0 { // первый вариант проверки, что строка не пустая
		str = "7,2,47,242,475,12,42" // если строка пустая, то пишем свою последовательность из примера
	}

	//if ttt == "" {            // второй вариант проверки, что строка не пустая
	//ttt = {7,2,47,242,475,12,42}
	//}

=======

	var str string

	fmt.Println("введите последовательность для сортировки в виде: \"1,33,5,56,8,42,11\" " +
		"или нажмите enter чтобы использовать последовательность из примера")
	fmt.Scanln(&str)

	 fmt.Println("считали:      ", str)

	//if len(str) == 0 {			 // метод проверки 1
	//	str= "1,33,5,56,8,42,11"
	//}

	if str == "" {					 // метод проверки 2
		str = "1,33,5,56,8,42,11"
	}

	fmt.Println("итого:      ", str)

>>>>>>> Stashed changes
	strs := strings.Split(str, ",")
	var ints []int
	for _, s := range strs {
		num, err := strconv.Atoi(s)
		if err == nil {
			ints = append(ints, num)
		}
	}

	fmt.Println("массив до сортировки:      ", ints)
<<<<<<< Updated upstream

	sort.Ints(ints)
	fmt.Println("массив после сортировки:   ", ints)
}
=======
	sort.Ints(ints)
	fmt.Println("массив после сортировки:   ", ints)

}
>>>>>>> Stashed changes
