package newapp

import (
	"io/ioutil"
	"path/filepath"
)

func createRouter(appPath string) {
	data := []byte(`package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// InitializeRouter sets up all the handlers and returns a *mux.Router.
func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(handler)
	}
	return router
}`)
	routerGoFile := filepath.Join(appPath, "router.go")
	ioutil.WriteFile(routerGoFile, data, filePerm)
}
