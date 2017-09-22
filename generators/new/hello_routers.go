package new

import (
	"io/ioutil"
	"path/filepath"

	"github.com/cumulusware/grab/helpers"
)

func createHelloRoutes(appPath string) {
	_, importPath, err := helpers.DetermineImportPath(appPath)
	if err != nil {
		// FIXME(mdr): Need to do something more than just panic.
		panic(err)
	}

	data := []byte(`package routers

import (
	"` + importPath + `/controllers"
	"` + importPath + `/core/authentication"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetHelloRoutes(router *mux.Router) *mux.Router {
	router.Handle(
		"/test/hello",
		negroni.New(
			negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.HelloController),
		),
	).Methods("GET")

	return router
}`)

	filename := filepath.Join(appPath, "routers", "hello.go")
	ioutil.WriteFile(filename, data, filePerm)
}
