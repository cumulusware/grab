// Copyright (c) 2015-2017 The grab developers. All rights reserved.
// Project site: https://github.com/cumulusware/grab
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

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
