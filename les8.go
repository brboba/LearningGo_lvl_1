package main

import (
	"./les8"
	"flag"
	"fmt"
	"log"
)

var (
	configTypes = flag.String("параметр при запуске", "значение которое передаем в параметре",
		"описание параметра команды")
)

func main() {
	flag.Parse()

	err, conf := ReadConf.GetConfig(*configTypes)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(conf)

}

/*
var port = flag.String("port", "8080", "port_number")
var db_url = flag.String("db_url", "gb.ru", "name")
var login = flag.String("login", "Login-Admin", "name")
var pass = flag.String("pass", "Admin-Pass", "name")
 */