package middlewares

import (
	"net/http"
)

type AuthenticatedHandler func(http.ResponseWriter, *http.Request, User)

type User struct {
	Id int
}

type Auth struct {
	handler AuthenticatedHandler
}

// Incoming requests are attached with user details
func (auth *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user, err := getAuthenticatedUser("")
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	auth.handler(w, r, user)
}

func NewAuth(handlerToWrap AuthenticatedHandler) *Auth {
	return &Auth{handlerToWrap}
}

// TODO
func getAuthenticatedUser(jwt string) (User, error) {
	// decode jwt and return user
	return User{}, nil
}
