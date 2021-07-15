package jwtInterface

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func JWTencode(claims map[string]interface{}, secret, alg string) string {
	var key = []byte(secret)
	var jwtClaims = jwt.MapClaims(claims)

	algorithm := jwt.GetSigningMethod(alg)

	token := jwt.NewWithClaims(algorithm, jwtClaims)

	tokenString, err := token.SignedString(key)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return tokenString
}

func JWTdecode(tokenString string) *jwt.Token {
	var token *jwt.Token
	var err error

	parser := new(jwt.Parser)
	token, _, err = parser.ParseUnverified(tokenString, jwt.MapClaims{})

	if err != nil {
		fmt.Errorf("[%v] Invalid token", err)
		return nil
	}
	return token
}
