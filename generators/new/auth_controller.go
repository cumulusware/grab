// Copyright (c) 2015-2017 The grab developers. All rights reserved.
// Project site: https://github.com/cumulusware/grab
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package new

import (
	"io/ioutil"
	"path/filepath"

	"github.com/cumulusware/grab/helpers"
)

func createAuthController(appPath string) {
	_, importPath, err := helpers.DetermineImportPath(appPath)
	if err != nil {
		// FIXME(mdr): Need to do something more than just panic.
		panic(err)
	}

	data := []byte(`package controllers

import (
	"` + importPath + `/services"
	"` + importPath + `/models"
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	responseStatus, token := services.Login(requestUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(responseStatus)
	w.Write(token)
}

func RefreshToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	requestUser := new(models.User)
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&requestUser)

	w.Header().Set("Content-Type", "application/json")
	w.Write(services.RefreshToken(requestUser))
}

func Logout(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := services.Logout(r)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}`)

	filename := filepath.Join(appPath, "controllers", "auth_controller.go")
	ioutil.WriteFile(filename, data, filePerm)
}
