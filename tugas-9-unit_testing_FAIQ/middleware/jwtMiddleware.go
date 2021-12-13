package middleware

import (
	"time"
	"tugas9unitTesting/constants"

	"github.com/dgrijalva/jwt-go"
)

type JwtCustomClaims struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.StandardClaims
}

func GenerateToken(userId int, name string, admin bool) (string, error) {
	claims := &JwtCustomClaims{
		Id:    userId,
		Name:  name,
		Admin: admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 22).Unix(),
		},
	}
	// fmt.Println(claims)
	// claims := jwt.MapClaims{}
	// claims["userId"] = userId
	// claims["name"] = name
	// claims["admin"] = admin
	// claims["exp"] = time.Now().Add(time.Minute * 2).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}
