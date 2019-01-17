package controllers

import (
	"encoding/json"
	"github.com/prince1809/go-web/TaskManager/common"
	"github.com/prince1809/go-web/TaskManager/data"
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

func Login(w http.ResponseWriter, r *http.Request) {

}
