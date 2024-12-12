package server

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/Rishi-Mishra0704/flowforge/backend/models"
	"github.com/Rishi-Mishra0704/flowforge/backend/utils"
	"github.com/labstack/echo/v4"
)

func (s *Server) GetFlowChartHandler(c echo.Context) error {
	// Parse the uploaded file
	file, err := c.FormFile("codebase")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse uploaded file")
	}

	src, err := file.Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to open uploaded file")
	}
	defer src.Close()

	// Read the entire file into memory
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, src)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to read uploaded file")
	}

	// Open the zip archive from the buffer
	zipReader, err := zip.NewReader(bytes.NewReader(buf.Bytes()), file.Size)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to read zip file")
	}

	// Concurrently process files
	contentChan := make(chan string)
	errChan := make(chan error, len(zipReader.File))
	var wg sync.WaitGroup

	for _, f := range zipReader.File {
		if f.FileInfo().IsDir() {
			continue // Skip directories
		}

		wg.Add(1)
		go func(f *zip.File) {
			defer wg.Done()
			rc, err := f.Open()
			if err != nil {
				errChan <- fmt.Errorf("failed to open file in zip: %v", err)
				return
			}
			defer rc.Close()

			fileContent, err := io.ReadAll(rc)
			if err != nil {
				errChan <- fmt.Errorf("failed to read file in zip: %v", err)
				return
			}

			contentChan <- string(fileContent)
		}(f)
	}

	// Close the channel after all files are processed
	go func() {
		wg.Wait()
		close(contentChan)
		close(errChan)
	}()

	var content []string
	for {
		select {
		case err := <-errChan:
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		case data, ok := <-contentChan:
			if !ok {
				contentChan = nil
			} else {
				content = append(content, data)
			}
		}
		if contentChan == nil && len(errChan) == 0 {
			break
		}
	}

	// Process the combined content
	combinedContent, err := utils.ReadMultipleFiles(content)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to read multiple files")
	}

	// Generate the flowchart via AI
	part := utils.AskAI(combinedContent, s.config)

	var flowchart models.Flowchart
	err = json.Unmarshal([]byte(part), &flowchart)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to decode JSON")
	}

	return c.JSON(http.StatusOK, flowchart)
}
