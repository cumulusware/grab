package newapp

import (
	"io/ioutil"
	"path/filepath"
)

func createAuthParameters(appPath string) {
	data := []byte(`package parameters

import ()

type TokenAuthentication struct {
	Token string ` + "`" + `json:"token" form:"token"` + "`" + `
}`)

	filename := filepath.Join(appPath, "api", "parameters", "auth.go")
	ioutil.WriteFile(filename, data, filePerm)
}
