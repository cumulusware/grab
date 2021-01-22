package newapp

import (
	"io/ioutil"
	"path/filepath"
)

func createRoutes(appPath string) {
	data := []byte(`package main

import "net/http"

// Route represents a REST API route.
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

// Routes is a slice of Route.
type Routes []Route

var routes = Routes{
	Route{"ReadAPI", "GET", "/api", ReadAPI},
	Route{"DescribeAPI", "OPTIONS", "/api", DescribeAPI},
}`)
	routesGoFile := filepath.Join(appPath, "routes.go")
	ioutil.WriteFile(routesGoFile, data, filePerm)
}
