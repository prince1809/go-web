package common

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

type AppClaims struct {
	UserName string `json:"username"`
	Role string `json:"role"`
	jwt.StandardClaims
}


// using asymmetric crypto/RSA keys
//location of private/public key files
const(

)


// GenerateJWT generates a new JWT token
func GenerateJWT(name, role string) (string, error) {
	return "", nil
}

// Authorize Middleware for validating JWT tokens
func Authorize(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

}


// ToeknFromAuthHeader is a "TokenExtractor" that takses a given request and extracts
// the JWT token from the Authorization header.
func TokenFromAuthHeader(r *http.Request) (string, error) {
	return "",nil
}



