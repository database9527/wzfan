package services

import (
	"math/rand"
	"os"
	"path/filepath" // Corrected from "path/filepath" to ensure it's standard

	"github.com/flosch/pongo2/v6"
)

const filesDir = "./public/static/files/" // Relative to project root

// FilterRandFile selects a random file from the 'filesDir' and returns its web-accessible path.
// Usage: {{ ""|rand_file }}
func FilterRandFile(in *pongo2.Value, param *pongo2.Value) (*pongo2.Value, *pongo2.Error) {
	entries, err := os.ReadDir(filesDir)
	if err != nil {
		// Log the error or handle it more gracefully in a real app
		// For now, return an empty string if directory can't be read
		return pongo2.AsValue(""), &pongo2.Error{
			Sender:    "filter:rand_file",
			OrigError: err,
			Message:   "Error reading files directory: " + filesDir,
		}
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	if len(files) == 0 {
		return pongo2.AsValue(""), nil // No files found, return empty string
	}

	randomIndex := rand.Intn(len(files))
	randomFileName := files[randomIndex]

	// Construct the web path.
	// Assuming app.Static("/", "./public") is used, so paths are relative to "./public"
	// The filesDir is "./public/static/files/", so web path is "/static/files/filename.ext"
	webPath := filepath.Join("/static/files", randomFileName) // filepath.Join handles path separators

	return pongo2.AsValue(webPath), nil
}
