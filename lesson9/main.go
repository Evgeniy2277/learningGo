package main

import (
	"fmt"
	"gihub.com/gouser/configNet/configNet"
	"log"
)

func main() {

	cfgPath, err := configNet.ParseFlags()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfgPath)
	cfg, err := configNet.NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	// Run the server
	cfg.Run()
	//fmt.Println(cfg.Server.Name)
	//fmt.Println(cfg.Server.Count)
}
