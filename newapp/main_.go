package newapp

import (
	"io/ioutil"
	"path/filepath"
)

func createMain(appPath, appName string) {
	data := []byte(`package main

import (
	"flag"
	"fmt"

	"github.com/codegangsta/negroni"
)

func main() {
	var (
		portFlag = flag.Int("port", 9090, "port to serve ` + appName + `.")
	)
	flag.Parse()
	port := *portFlag
	addr := fmt.Sprintf(":%d", port)
	handler := InitializeRouter()
	n := negroni.Classic()
	n.UseHandler(handler)
	n.Run(addr)
}`)
	mainGoFile := filepath.Join(appPath, "main.go")
	ioutil.WriteFile(mainGoFile, data, filePerm)
}
