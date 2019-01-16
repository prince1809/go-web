package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// using  asymmetric crypto/RSA keys
const (
	privKeyPath = "keys/app.rsa"     // opensssl genrsa -out app.rsa 1024
	pubKeyPath  = "keys/app.rsa.pub" //openssl rsa -in app.rsa -pubout > app.rsa.pub
)

// verify key and sign key
var (
	verifyKey, signKey []byte
)

//struct User for parsing login credentials
type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type AppClaims struct {
	userName string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

//read the key files before starting http handlers
func init() {
	var err error

	signKey, err = ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal("Error reading private key")
		return
	}
	verifyKey, err = ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal("Error reading public key")
		return
	}
}

// reads the login credentials, checks them and creates JWT token
func loginHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	//decode into User struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error in request body")
		return
	}

	//validates user credentials
	if user.UserName != "prince" && user.Password != "pass" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Wrong credntials")
		return
	}

	//create the claims
	claims := AppClaims{
		user.UserName,
		"Member",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			Issuer:    "admin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Sorry, error while signing token!")
		log.Println("Token signing error: %v\n", err)
		return
	}
	response := Token{Token: tokenString}
	jsonResponse(response, w)
}

//only accessible with a valid token
func authHandler(w http.ResponseWriter, r *http.Request) {
	//get token from request
	//token, err := jwt.Parse

	//})

}

type Response struct {
	Text string `json:"text"`
}

type Token struct {
	Token string `json:"token"`
}

func jsonResponse(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/auth", authHandler).Methods("POST")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()

}
