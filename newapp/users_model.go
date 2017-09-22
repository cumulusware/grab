package newapp

import (
	"io/ioutil"
	"path/filepath"
)

func createUserModel(appPath string) {

	data := []byte(`package models

type User struct {
	UUID 		 string ` + "`" + `json:"uuid" form:"-"` + "`" + `
	Username string ` + "`" + `json:"username" form:"username"` + "`" + `
	Password string ` + "`" + `json:"password" form:"password"` + "`" + `
}`)

	filename := filepath.Join(appPath, "models", "users.go")
	ioutil.WriteFile(filename, data, filePerm)
}
