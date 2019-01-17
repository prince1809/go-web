package controllers

import "github.com/prince1809/go-web/TaskManager/models"

// Models for JSON resources
type (
	// For Post - /users/register
	UserResource struct {
		Data models.User `json:"data"`
	}

	// For Post - /users/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}

	// Response for authorized user post - /users/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}

	// For Post/Put - /tasks
	// For Get - /task/id
	TaskResource struct {
		Data models.Task `json:"data"`
	}

	// For Get - /tasks
	TasksResource struct {
		Data []models.Task `json:"data"`
	}

	// For Post/Put - /notes
	NoteResource struct {
		Data NoteModel `json:"data"`
	}

	NotesResource struct {
		Data []models.TaskNote `json:"data"`
	}

	// Model for authentication
	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Model for authorized user with access token
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}

	//Model for a TaskNote
	NoteModel struct {
		TaskId      string `json:"taskid"`
		Description string `json:"description"`
	}
)
