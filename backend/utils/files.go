package utils

import (
	"io"
	"log"
	"os"
)

func OpenFile() (*os.File, error) {
	file, err := os.Open("test.py")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return file, nil
}

func ReadFileContent(file *os.File) ([]byte, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return content, nil
}
