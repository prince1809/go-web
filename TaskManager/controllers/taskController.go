package controllers

import (
	"encoding/json"
	httpcontext "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/prince1809/go-web/TaskManager/common"
	"github.com/prince1809/go-web/TaskManager/data"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

// CreateTask insert a new task document
// Handler for HTTP Post - "/task"
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var dataResource TaskResource
	// Decode the incoming Task json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Task data", 500)
		log.Println(err)
		return
	}
	task := &dataResource.Data
	context := NewContext()
	defer context.close()
	val := r.Context().Value("user")
	context.User = val.(string)
	log.Printf("Context : %s", httpcontext.GetAll(r))
	task.CreatedBy = context.User
	col := context.DbCollection("tasks")
	repo := &data.TaskRepository{C: col}
	// insert a task document
	repo.Create(task)
	j, err := json.Marshal(TaskResource{Data: *task})
	if err != nil {
		common.DisplayAppError(w, err, "An unexptected eror has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

// GetsTasks returns all Task document
// Handler for HTTP Get - "/tasks"
func GetTasks(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.close()
	col := context.DbCollection("tasks")
	repo := data.TaskRepository{C: col}
	tasks := repo.GetAll()
	log.Printf("Context : %s", context.User)
	j, err := json.Marshal(TasksResource{Data: tasks})
	if err != nil {
		common.DisplayAppError(w, err, "An unexptected error occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

//GetTaskById returns a single Task document by id
// Handler for HTTP Get - "/tasks/{id}"
func GetTaskByID(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.close()

	col := context.DbCollection("tasks")
	repo := &data.TaskRepository{C: col}
	task, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNoContent)
		} else {
			common.DisplayAppError(w, err, "An unexpected error occurred", 500)
			return
		}
	}
	j, err := json.Marshal(task)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error occurred", 500)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// GetTasksByUser returns all Tasks created by a User
// Handler for HTTP Get - "/tasks/users/{id}"
func GetTasksByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := vars["id"]
	context := NewContext()
	defer context.close()
	col := context.DbCollection("tasks")
	repo := &data.TaskRepository{C: col}
	tasks := repo.GetByUser(user)
	j, err := json.Marshal(TasksResource{Data: tasks})
	if err != nil {
		common.DisplayAppError(w, err, "An unexptected error occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// UpdateTask update an existing Task document
// Handler for HTTP Put - "/task/{id}"
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := bson.ObjectId(vars["id"])
	var dataResoure TaskResource
	err := json.NewDecoder(r.Body).Decode(&dataResoure)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid task data", 500)
		return
	}
	task := &dataResoure.Data
	task.Id = id
	context := NewContext()
	defer context.close()
	col := context.DbCollection("tasks")
	repo := &data.TaskRepository{C:col}
	if err := repo.Update(task); err != nil {
		common.DisplayAppError(w, err, "An unexptected error occurred", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// DeleteTask delete an existing Task document
// Handler for HTTP Delete - "/tasks/{id}"
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext()
	defer context.close()
	col := context.DbCollection("tasks")
	repo := &data.TaskRepository{C:col}
	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(w, err, "An Unexptected error occurred", 500)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
