package util

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
)

type Middleware func(http.Handler) http.Handler

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tokenString string
		fmt.Println("hitting auth middleware")
		if tokens, ok := r.Header["Authorization"]; ok {
			tokenString = tokens[0]
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		}

		token, err := jwt.ParseWithClaims(
			tokenString,
			&UserClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return secret, nil
			},
		)

		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("could not parse token"))
			return
		}

		if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
			context.Set(r, "UserId", claims.UserId)
			next.ServeHTTP(w, r)
		} else {
			fmt.Println("not authorized!!!!")
			w.WriteHeader(http.StatusUnauthorized)
		}
	})
}
