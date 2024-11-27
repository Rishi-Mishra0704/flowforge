package utils

import (
	"io"
	"log"
	"os"
	"path/filepath"
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

// ReadMultipleFiles reads all files in a given directory and returns their combined content as a string.
func ReadMultipleFiles(dirPath string) (string, error) {
	var contentBuilder string

	// Walk through the directory and read files
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only process files (skip directories)
		if !info.IsDir() {
			// Open the file using OpenFile function
			file, err := OpenFile(path)
			if err != nil {
				return err
			}
			defer file.Close()

			// Read the content of the file
			content, err := ReadFileContent(file)
			if err != nil {
				return err
			}

			// Append content to the combined content string
			contentBuilder += string(content) + "\n"
		}
		return nil
	})

	// Return the combined content or an error if any occurred
	return contentBuilder, err
}
