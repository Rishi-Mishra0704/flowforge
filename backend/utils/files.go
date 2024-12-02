package utils

import (
	"io"
	"log"
	"os"
)

// OpenFile opens a file and returns the file pointer.
func OpenFile(filePath string) (*os.File, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return file, nil
}

// ReadFileContent reads content from a file and returns it as a byte slice.
func ReadFileContent(file *os.File) ([]byte, error) {
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return content, nil
}

// ReadMultipleFiles combines the content of multiple files provided as a slice of strings.
func ReadMultipleFiles(contents []string) (string, error) {
	var contentBuilder string

	// Iterate over the provided file contents
	for _, content := range contents {
		contentBuilder += content + "\n"
	}

	return contentBuilder, nil
}
