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

	fmt.Println("Running story from file:", *fileFlag)
	file, err := os.Open(*fileFlag)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	story, err := cyoa.JSONStory(file)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Printf("%+v", story)
}
