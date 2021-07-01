package main

import (
	//"flag"
	"fmt"
	"learningGo/lesson8/foo"
	"net/http"
)


func runServer(config foo.Config) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		data := ""
		for i := 0; i < config.Count; i++ {
			data += fmt.Sprintf("hello, %s\n", config.Name)
		}
		_, err := fmt.Fprint(writer, data)
		if err != nil {
			return
		}
	})

	err := http.ListenAndServe("localhost:"+config.Port, nil)
	if err != nil {
		return
	}
}
func main() {
	c := foo.NewConfig()
	fmt.Println(*c)
	runServer(*c)
}
