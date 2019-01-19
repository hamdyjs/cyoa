package main

import (
	"flag"
	"fmt"
	"os"

	cyoa "github.com/hamdyjs/go_cyoa"
)

func main() {
	fileFlag := flag.String("file", "gopher.json", "The json file containing the story")
	flag.Parse()

	file, err := os.Open(*fileFlag)
	if err != nil {
		fmt.Println(err)
		return
	}

	story, err := cyoa.JSONStory(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", story)
}
