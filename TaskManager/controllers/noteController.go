package controllers

import "net/http"

// CreateNote inserts a new Note document for a TaskId
// Handler for HTTP Post - "/notes"
func CreateNote(w http.ResponseWriter, r *http.Request) {

}


// GetNotesByTask returns all Nodes documents under a TaskId
// Handle for HTTP Get - "/notes/tasks/{id}"
func GetNotesByTask(w http.ResponseWriter, r *http.Request) {

}


// GetNotes returns all Note documents
// Handler for HTTP Get - "/notes"
func GetNotes(w http.ResponseWriter, r *http.Request) {

}

// GetNoteByID returns a single Note document by ID
// Handler for HTTP Get - "/notes/{id}"
func GetNoteByID(w http.ResponseWriter, r *http.Request) {

}

func UpdateNote(w http.ResponseWriter, r *http.Request) {

}

func DeleteNote(w http.ResponseWriter, r *http.Request) {

}
