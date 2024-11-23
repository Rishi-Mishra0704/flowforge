package main

import (
	"log"

	"github.com/Rishi-Mishra0704/flowforge/backend/utils"
)

func main() {
	file, err := utils.OpenFile()
	if err != nil {
		log.Fatal(err)
	}

	content, err := utils.ReadFileContent(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	log.Println(string(content))
}
