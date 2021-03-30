package main

import (
	"github.com/krzysiek27/simple-http-server/server"
	"fmt"
	"os"
)

func main(){
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "No arguments")
		return
	}

	switch command := os.Args[1]; command {
	case "version":
		fmt.Println(server.GetVersionString())
	case "help":
		fmt.Print(server.GetHelpString())
	case "run":
		if len(os.Args) < 4 || os.Args[2] != "--file" {
			fmt.Fprintln(os.Stderr, "Invalid arguments")
			return
		}

		fmt.Println("Serving file " + os.Args[3] + " on port 8080")
		err := server.ServeHtmlFile(os.Args[3])
		fmt.Fprintln(os.Stderr, err)
	default:
		fmt.Fprintln(os.Stderr, "Invalid arguments")
	}
}