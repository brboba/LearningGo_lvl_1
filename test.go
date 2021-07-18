package main

import (
	"fmt"
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
	}
}