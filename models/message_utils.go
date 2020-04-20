package models

import (
	"errors"
	"net/http"
)

type TokenCredentils struct {
	Role			string 			`json:"role"`
	Username		string			`json:"username"`
}

func (t *TokenCredentils) GetTokenCredentials(r *http.Request) (error) {

	t.Role = r.Header.Get("role")
	t.Username = r.Header.Get("username")

	if t.Role == "" || t.Username == "" {
		return errors.New("invalid token credentials")
	}

	return nil

}

