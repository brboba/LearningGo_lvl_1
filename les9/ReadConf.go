package ReadConf

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port        int64
	DbUrl       string
	JaegerUrl   string
	SentryUrl   string
	KafkaBroker string
	SomeAppId   string
	SomeAppKey  string
}

func GetConfig() (Config, error) {
	return getArgumentsConfig()
}

func getArgumentsConfig() (Config, error) {
	port := flag.String("port", "80", "Description")
	dburl := flag.String("dburl","ya.ru","adress base")

	flag.Parse()

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

	data := make([]byte, 64)

	for{
		n, err := file.Read(data)
		if err == io.EOF{   // если конец файла
			break           // выходим из цикла
		}
		fmt.Print(string(data[:n]))
	}

	config := Config{}

	config.Port, _ = strconv.ParseInt(*port, 10, 64)
	config.DbUrl = *dburl

	return config, nil

}


