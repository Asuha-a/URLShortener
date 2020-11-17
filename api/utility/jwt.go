package utility

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

// ParseJWT parse jwt and return uuid and permission
func ParseJWT(tokenString string) (string, string, error) {
	mySingningKey := []byte("AllYourBase")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return mySingningKey, nil
	})
	if err != nil {
		panic(err)
	}
	var uuid string
	var permission string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uuid = claims["uuid"]
		permission = claims["permission"]
	}
	return uuid, permission, nil
}
