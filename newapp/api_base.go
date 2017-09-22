package newapp

import (
	"io/ioutil"
	"path/filepath"
)

func createAPIBase(appPath string) {
	data := []byte(`package main

import (
	"net/http"
)

const (
	readLimit = 1048576 // 1 Mbyte
)

// ReadAPI is the handler for the GET /api endpoint.
func ReadAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// DescribeAPI is the handler for the OPTIONS /api endpoint.
func DescribeAPI(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "GET,HEAD,OPTIONS")
	w.WriteHeader(http.StatusOK)
}`)
	filename := filepath.Join(appPath, "api_base.go")
	ioutil.WriteFile(filename, data, filePerm)
}
