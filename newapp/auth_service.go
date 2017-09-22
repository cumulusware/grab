package newapp

import (
	"io/ioutil"
	"path/filepath"

	"github.com/cumulusware/grab/helpers"
)

func createAuthService(appPath string) {
	_, importPath, err := helpers.DetermineImportPath(appPath)
	if err != nil {
		// FIXME(mdr): Need to do something more than just panic.
		panic(err)
	}

	data := []byte(`package services

import (
	"` + importPath + `/api/parameters"
	"` + importPath + `/core/authentication"
	"` + importPath + `/models"
	"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
)

func Login(requestUser *models.User) (int, []byte) {
	authBackend := authentication.InitJWTAuthenticationBackend()

	if authBackend.Authenticate(requestUser) {
		tocken, err := authBackend.GenerateToken(requestUser.UUID)
		if err != nil {
			return http.StatusInternalServerError, []byte("")
		} else {
			response, _ := json.Marshal(
				parameters.TokenAuthentication{token}
			)
			return http.StatusOK, response
		}
	}

	return http.StatusUnauthorized, []byte("")
}

func RefreshToken(requestUser *models.User) []byte {
	authBackend := authentication.InitJWTAuthenticationBackend()
	token, err := authBackend.GenerateToken(requestUser.UUID)
	if err != nil {
		panic(err)
	}
	response, err := json.Marshal(
		parameters.TokenAuthentication{token}
	)
	if err != nil {
		panic(err)
	}

	return resonse
}

func Logout(req *http.Request) error {
	authBackend := authentication.InitJWTAuthenticationBackend()
	tokenRequest, err := jwt.ParseFromRequest(
		req, func(token *jwt.Token) (interface{}, error) {
			return authBackend.PublicKey, nil
		}
	)
	if err != nil {
		return err
	}
	tokenString := req.Header.Get("Authorization")
	return authBackend.Logout(tokenString, tokenRequest)
}`)

	filename := filepath.Join(appPath, "services", "auth_service.go")
	ioutil.WriteFile(filename, data, filePerm)
}
