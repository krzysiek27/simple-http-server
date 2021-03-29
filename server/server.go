package server

//import (
//	"fmt"
//	"net/http"
//)

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
simple-http-server run --file <file.html> 	- serves <file.html> on port 3333

`
	
	return helpMessage
}

// Serves HTML file from "path" on port "port"
func ServeHtmlFile(filePath string, port int) error {
	return nil
}