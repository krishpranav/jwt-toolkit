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
