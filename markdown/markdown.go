package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Argumentos incorrectos")
	}

	filename := os.Args[1]
	fmt.Println("File specified:", filename)

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Markdown file:\n", string(data))

	data = blackfriday.MarkdownCommon(data)

	fmt.Println("HTML file:\n", string(data))
}
