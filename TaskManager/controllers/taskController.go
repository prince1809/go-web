package controllers

import "net/http"

// CreateTask insert a new task document
// Handler for HTTP Post - "/task"
func CreateTask(w http.ResponseWriter, r *http.Request) {

}

// GetsTasks returns all Task document
// Handler for HTTP Get - "/tasks"
func GetTasks(w http.ResponseWriter, r *http.Request) {

}

//GetTaskById returns a single Task document by id
// Handler for HTTP Get - "/tasks/{id}"
func GetTaskByID(w http.ResponseWriter, r *http.Request) {

}

// GetTasksByUser returns all Tasks created by a User
// Handler for HTTP Get - "/tasks/users/{id}"
func GetTasksByUser(w http.ResponseWriter, r *http.Request) {

}

// UpdateTask update an existing Task document
// Handler for HTTP Put - "/task/{id}"
func UpdateTask(w http.ResponseWriter, r *http.Request) {

}


// DeleteTask delete an existing Task document
// Handler for HTTP Delete - "/tasks/{id}"
func DeleteTask(w http.ResponseWriter, r *http.Request) {

}
