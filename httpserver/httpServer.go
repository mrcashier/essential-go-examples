package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port int
var help bool

func main() {
	flag.BoolVar(&help, "help", false, "help usage")
	flag.IntVar(&port, "port", 3000, "The port to run the server on asas")

	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	fmt.Println("Serving files on localhost: %v ", port)

	err := ServeStatic(port)
	if err != nil {
		log.Fatalln(err)
	}
}

func ServeStatic(port int) error {
	host := fmt.Sprintf("localhost:%v", port)
	return http.ListenAndServe(host, http.FileServer(http.Dir(".")))
}
