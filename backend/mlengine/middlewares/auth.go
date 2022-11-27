package middlewares

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt"
)

type AuthenticatedHandler func(http.ResponseWriter, *http.Request, User)

type User struct {
	Id int
}

type Auth struct {
	handler AuthenticatedHandler
	secret  string
}

func NewAuth(handlerToWrap AuthenticatedHandler) *Auth {
	return &Auth{handler: handlerToWrap, secret: os.Getenv("JWT_SECRET")}
}

func (auth *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, err := auth.getAuthenticatedUser(strings.Split(r.Header.Get("Authorization"), "Bearer ")[1])
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	auth.handler(w, r, user)
}

func (auth *Auth) getAuthenticatedUser(tokenStr string) (User, error) {
	// if auth.secret == "" {
	// 	return User{Id: 1}, nil
	// }

	// decode jwt and return user
	claims := jwt.MapClaims{}
	hmacSecret := []byte(auth.secret)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// TODO: check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return User{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		var userId = int(claims["user_id"].(float64))
		return User{Id: userId}, nil
	}

	return User{}, errors.New("Invalid JWT Token")

}
