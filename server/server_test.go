package server_test

import (
	"io"
	"io/ioutil"
	"github.com/krzysiek27/simple-http-server/server"
	"testing"
	"net/http/httptest"
)

func TestNoFileError(t *testing.T){
	err := server.ServeHtmlFile("file_that_do_not_exist.html")
	if err.Error() != "File do not exist." {
		t.Fail()
	}
}

func TestHandler(t *testing.T){
	var filePath string = "../test_files/test.html"

	request := httptest.NewRequest("GET", "/", nil)
	writer := httptest.NewRecorder()
	handler := server.ServeFileHandler{filePath}
	handler.ServeHTTP(writer, request)
	response := writer.Result()

	buf := make([]byte, 10000)
	n, _ := io.ReadFull(response.Body, buf)
	body := string(buf[:n])

	original_b, _ := ioutil.ReadFile(filePath)
	original := string(original_b)

	if(body != original){
		t.Fail()
	}
}