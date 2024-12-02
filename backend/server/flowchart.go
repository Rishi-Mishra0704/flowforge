package server

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

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

	// In-memory processing of extracted files
	var content []string // Collect all file contents here

	for _, f := range zipReader.File {
		if f.FileInfo().IsDir() {
			continue // Skip directories
		}

		// Open the file inside the zip
		rc, err := f.Open()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to open file in zip")
		}

		// Read the file content into memory
		fileContent, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to read file in zip")
		}

		// Process the content (e.g., store it as a string for further processing)
		content = append(content, string(fileContent))
	}

	// Process the content array as needed
	combinedContent, err := utils.ReadMultipleFiles(content)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to read multiple files")
	}
	log.Println(combinedContent)
	part := utils.AskAI(combinedContent, s.config)

	// Decode the JSON data into the Flowchart struct
	var flowchart models.Flowchart
	err = json.Unmarshal([]byte(part), &flowchart)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to decode JSON")
	}

	// Optionally, iterate over nodes and edges to demonstrate access
	for _, node := range flowchart.Nodes {
		log.Printf("Node: ID=%d, Label=%s, Type=%s", node.ID, node.Label, node.Type)
	}

	for _, edge := range flowchart.Edges {
		log.Printf("Edge: Source=%d, Target=%d, Condition=%s", edge.Source, edge.Target, edge.Condition)
	}

	return c.JSON(http.StatusOK, flowchart)
}
