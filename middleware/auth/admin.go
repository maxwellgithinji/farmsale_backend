package auth

import (
	"context"
	"encoding/json"
	"github.com/maxwellgithinji/farmsale_backend/models/jwtmodel"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

//AdminVerify verifies user is admin
func AdminVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := verifyTokenHelper(w, r)
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "User not authorized"})
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &jwtmodel.Token{}, func(token *jwt.Token) (interface{}, error) {
			return decodebs, nil
		})

		if claims, ok := token.Claims.(*jwtmodel.Token); ok && token.Valid {
			if !claims.Isadmin {
				w.WriteHeader(http.StatusForbidden)
				json.NewEncoder(w).Encode(Exception{Message: "User not authorized, Admin only"})
				fmt.Println(err)
				return
			}
			ctx := context.WithValue(r.Context(), "user", token)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(Exception{Message: "User not authorized"})
			fmt.Println(err)
			return
		}
	})
}
