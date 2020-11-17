package utility

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

// ParseJWT parse jwt and return uuid and permission
func ParseJWT(tokenString string) (uuid.UUID, string, error) {
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
	var uUID uuid.UUID
	var permission string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		uUID = uuid.FromString(fmt.Sprintf("%v", claims["uuid"]))
		permission = fmt.Sprintf("%v", claims["permission"])
	}
	return uUID, permission, nil
}
