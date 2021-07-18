package main

import (
	"./les9"
	"fmt"
	"os"
)

func main() {
	conf, err := ReadConf.GetConfig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(conf)

}
