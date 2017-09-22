// Copyright (c) 2015-2017 The grab developers. All rights reserved.
// Project site: https://github.com/cumulusware/grab
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

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
