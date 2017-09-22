package newapp

import (
	"io/ioutil"
	"path/filepath"
)

func createRouter(appPath string) {
	data := []byte(`package routers

import (
	"github.com/gorilla/mux"
)

// InitializeRouter sets up all the handlers and returns a *mux.Router.
func InitializeRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router = SetAuthenticationRoutes(router)
	return router
}`)
	filename := filepath.Join(appPath, "routers", "router.go")
	ioutil.WriteFile(filename, data, filePerm)
}
