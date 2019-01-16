package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/pat"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/twitter"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Struct for parsing JSON configuration
type Configuration struct {
	TwitterKey     string
	TwitterSecret  string
	FacebookKey    string
	FacebookSecret string
}

var config Configuration

// Read configuration values from config.json
func init() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err := decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
}

func callbackAuthHandler(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	t, _ := template.New("userinfo").Parse(userTemplate)
	t.Execute(w, user)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("index").Parse(indexTemplate)
	t.Execute(w, nil)
}

func main() {
	//Register OAuth2 providers with Goth
	goth.UseProviders(
		twitter.New(config.TwitterKey, config.TwitterSecret, "http://localhost:8080/auth/twitter/callback"),
		facebook.New(config.FacebookKey, config.FacebookSecret, "http://localhost:8080/auth/facebook/callback"),
	)
	//Routing using Pat package
	r := pat.New()
	r.Get("/auth/{provider}/callback", callbackAuthHandler)
	r.Get("/auth/{provider}", gothic.BeginAuthHandler)
	r.Get("/", indexHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}

// view templates
var indexTemplate = `
<p><a href="/auth/twitter">Log in with Twitter</a></p>
<p><a href="/auth/facebook">Log in with Facebook</a></p>
`

var userTemplate = `
<p>Name : {{.Name}}</p>
<p>Email : {{.Email}}</p>
<p>NickName : {{.NickName}}</p>
<p>Location : {{.Location}}</p>
<p>AvatarURL : {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description : {{.Description}}</p>
<p>UserID : {{.UserID}}</p>
<p>AccessToken : {{.AccessToken}}</p>
`
