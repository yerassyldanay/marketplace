package models

import "github.com/dgrijalva/jwt-go"

type Token struct {
	Username string `json:"username,omitempty"`
	Roles    string `json:"roles,omitempty"`
	jwt.StandardClaims
}
