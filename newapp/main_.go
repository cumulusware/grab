package newapp

import (
	"io/ioutil"
	"path/filepath"

	"github.com/cumulusware/grab/helpers"
)

func createMain(appPath, appName string) {
	_, importPath, err := helpers.DetermineImportPath(appPath)
	if err != nil {
		// FIXME(mdr): Need to do something more than just panic.
		panic(err)
	}

	data := []byte(`package main

import (
	"` + importPath + `/routers"
	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {
	router := routers.InitRouters()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)
}`)
	mainGoFile := filepath.Join(appPath, "main.go")
	ioutil.WriteFile(mainGoFile, data, filePerm)
}
