package server

import (
	"context"
	"dimitrisCBR/bookie-api/v2/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

type authHelper struct {
	secret string
}

type claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type contextKey string

func (c contextKey) String() string {
	return "mypackage context key " + string(c)
}

var (
	contextKeyAuthtoken = contextKey("server-token")
)

func (a *authHelper) newToken(user model.User) string {
	expireTime := time.Now().Add(time.Hour * 1)
	c := claims{
		user.Username,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "localhost!",
		}}

	token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString([]byte(a.secret))

	return token
}

func (a *authHelper) validate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		cookie, err := req.Cookie("Auth")
		if err != nil {
			Error(res, http.StatusUnauthorized, "No authorization cookie")
			return
		}

		token, err := jwt.ParseWithClaims(cookie.Value, &claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected siging method")
			}
			return []byte(a.secret), nil
		})

		if err != nil {
			Error(res, http.StatusUnauthorized, "Invalid token")
			return
		}

		if claims, ok := token.Claims.(*claims); ok && token.Valid {
			ctx := context.WithValue(req.Context(), contextKeyAuthtoken, *claims)
			next(res, req.WithContext(ctx))
		} else {
			Error(res, http.StatusUnauthorized, "Unauthorized")
			return
		}
	})
}
