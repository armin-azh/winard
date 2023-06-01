package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/armin-azh/winard/keys/auth"
	"github.com/armin-azh/winard/serializer"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
)

func AuthJWTMiddleware(next http.HandlerFunc, secret []byte) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		authHeader, ok := r.Header["Authorization"]
		if ok && len(authHeader) > 0 {
			token, _ := jwt.Parse(authHeader[0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error in %s", "parsing")
				}
				return secret, nil
			})

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				encoded, _ := json.Marshal(serializer.Message{Message: "Couldn't parse claims"})
				_, err := fmt.Fprint(w, encoded)
				if err != nil {
					http.Error(w, err.Error(), http.StatusUnauthorized)
					return
				}
				return
			}
			exp := claims["exp"].(float64)
			if int64(exp) < time.Now().Local().Unix() {
				w.WriteHeader(http.StatusUnauthorized)
				encoded, _ := json.Marshal(serializer.Message{Message: "You token has been expired"})
				_, err := fmt.Fprint(w, string(encoded))
				if err != nil {
					return
				}
				return
			} else {

				// Set User Prime to the header
				r.Header.Set(auth.MidRequestUserKey, claims["user"].(string))
				next(w, r)
			}
		} else {
			encoded, _ := json.Marshal(serializer.Message{Message: "You should be authorized"})
			_, err := fmt.Fprintf(w, "%s", encoded)
			w.WriteHeader(http.StatusUnauthorized)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			return

		}
	}
}
