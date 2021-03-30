package server

import (
	"net/http"
	"os"
	"errors"
)

// Returns a string containing name and version of this program
func GetVersionString() string {
	return "simple-http-server 1.0"
}

// Returns a help message
func GetHelpString() string {

	helpMessage := 
`
A simple command-line interface HTTP server.

USAGE:
simple-http-server version			- shows version of this program
simple-http-server help				- shows this message
simple-http-server run --file <file.html> 	- serves <file.html> on port 8080

`
	
	return helpMessage
}

type ServeFileHandler struct{
	FileName string 
}

func (t *ServeFileHandler)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, t.FileName)
}

// Serves HTML file from "filePath"
func ServeHtmlFile(filePath string) error {
	_, fileErr := os.Stat(filePath)
	if os.IsNotExist(fileErr) {
		return errors.New("File do not exist.")
	}

	handler := ServeFileHandler{filePath}
	err := http.ListenAndServe(":8080", &handler)
	return err
}