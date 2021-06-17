package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
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

	strs := strings.Split(str, ",")
	var ints []int
	for _, s := range strs {
		num, err := strconv.Atoi(s)
		if err == nil {
			ints = append(ints, num)
		}
	}

	fmt.Println("массив до сортировки:      ", ints)

	sort.Ints(ints)
	fmt.Println("массив после сортировки:   ", ints)
}
