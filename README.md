# simple-http-server

It's a Go implementation of a simple HTTP server that can be used to serve a single HTML file.

## Building
To build a binary go to the cloned repository and run
```
go build
```
A ready-to-use binary should appear in your working directory.

## Usage

To serve a file `file.html` on port 8080 you should run
```
./simple-http-server run --file file.html
```
