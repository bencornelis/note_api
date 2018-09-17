package util

import (
	"fmt"
	"time"

	"github.com/bencornelis/note_api/model"
	jwt "github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	UserId uint `json:"id"`
	jwt.StandardClaims
}

var secret = []byte("foobar")

func GenerateToken(user *model.User) (string, error) {
	expiresAt := time.Now().Add(time.Second * 3000000).Unix()

	claims := UserClaims{
		user.ID,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    "testing!",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(secret)
	fmt.Println(ss)
	return ss, err
}
