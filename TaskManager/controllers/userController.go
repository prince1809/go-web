package controllers

import (
	"encoding/json"
	"github.com/prince1809/go-web/TaskManager/common"
	"github.com/prince1809/go-web/TaskManager/data"
	"github.com/prince1809/go-web/TaskManager/models"
	"net/http"
)

// Handler for HTTP Post - "/users/register
// Add a new User document
func Register(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
	//decode the incoming user json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid user data", 500)
		return
	}
	user := &dataResource.Data
	context := NewContext()
	defer context.close()
	col := context.DbCollection("users")
	repo := &data.UserRepository{C: col}
	// Insert user document
	repo.CreateUser(user)
	user.HashPassword = nil
	j, err := json.Marshal(UserResource{Data: *user})
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

// Login authenticates the HTTP request with username and password
// handler for HTTP Post - "/users/login"
func Login(w http.ResponseWriter, r *http.Request) {
	var dataResource LoginResource
	var token string
	// Decode the incoming login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Login data", 500)
		return
	}

	loginModel := dataResource.Data
	loginUser := models.User{
		Email: loginModel.Email,
		Password: loginModel.Password,
	}

	context := NewContext()
	defer context.close()
	col := context.DbCollection("users")
	repo := &data.UserRepository{C: col}
	// Authneticate the loging user
	user, err := repo.Login(loginUser)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid login credentials", 401)
		return
	}
	// Generate JWT token
	token, err = common.GenerateJWT(user.Email, "member")
	if err != nil {
		common.DisplayAppError(w, err, "Error while generating the access token", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//clean up the hashpassword to eliminate it from response json
	user.HashPassword = nil
	authUser := AuthUserModel{
		User:user,
		Token:token,
	}
	j, err := json.Marshal(AuthUserResource{Data: authUser})
	if err != nil {
		common.DisplayAppError(w, err, "An unexptected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
