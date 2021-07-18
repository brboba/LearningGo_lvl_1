package ReadConf

import (
	"errors"
	"flag"
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

	config := Config{}

	if !isNumeric(*port) {                                         //if isNumeric==false
		return config, errors.New("Port is not number")
	}

	config.Port, _ = strconv.ParseInt(*port, 10, 64)
	config.DbUrl = *dburl

	return config, nil
}

func isNumeric(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}
