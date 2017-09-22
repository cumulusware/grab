package newapp

import (
	"io/ioutil"
	"path/filepath"
)

func createHelloController(appPath string) {

	data := []byte(`package controllers

import (
	"net/http"
)

func HelloController(
	w http.ResponseWriter, r *http.Request, next http.HandlerFunc
) {
	w.Write([]byte("Hello, World!"))
}`)

	filename := filepath.Join(appPath, "controllers", "hello_controller.go")
	ioutil.WriteFile(filename, data, filePerm)
}
