package newapp

import (
	"io/ioutil"
	"path/filepath"

	"github.com/cumulusware/grab/helpers"
)

func createAuthRoutes(appPath string) {
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

func SetAuthenticationRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/token-auth", controllers.Login).Methods("POST")
	router.Handle("/refresh-token-auth", negroni.New(
		negroni.HandlerFunc(authentication.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.RefreshToken),
	)).Methods("GET")
	router.Handle("/logout", negroni.New(
		negroni.HandlerFunc(authentication.RequireTokenAuthentication),
		negroni.HandlerFunc(controllers.Logout),
	)).Methods("GET")

	return router
}`)

	filename := filepath.Join(appPath, "routers", "authentication.go")
	ioutil.WriteFile(filename, data, filePerm)
}
