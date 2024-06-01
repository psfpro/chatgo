package domain

import "github.com/gofrs/uuid"

type Task struct {
	ID                uuid.UUID
	Title             string
	Description       string
	State             State
	TechnicalSolution string
	Comments          string
	Files             map[string]*File
	Iteration         int32
}

func NewTask(id uuid.UUID, title string, description string) *Task {
	return &Task{ID: id, Title: title, Description: description, State: StateNewRequest, Files: make(map[string]*File)}
}

type File struct {
	Path     string
	Language string
	Content  string
}