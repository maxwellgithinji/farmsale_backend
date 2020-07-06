package auth

import (
	"github.com/maxwellgithinji/farmsale_backend/utils"
	"net/http"
	"os"
	"strings"
)

var (
	decodebs = []byte(os.Getenv("TOKEN_SECRET"))
)

//Exception struct
type Exception utils.Exception

func verifyTokenHelper(w http.ResponseWriter, r *http.Request) string {
	var tokenString = r.Header.Get("Authorization") //Grab the token from the header
	//ensure there is a token
	if len(tokenString) == 0 {
		return ""
	}
	// if the token has a bearer, slice it and return the token part
	if len(strings.Split(tokenString, " ")) == 2 {
		tokenString = strings.Split(tokenString, "Bearer ")[1]
		return tokenString
	} else if len(strings.Split(tokenString, " ")) > 2 {
		return ""
	}
	return tokenString
}
