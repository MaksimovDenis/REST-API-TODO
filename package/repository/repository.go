package repository

import (
	todo "firstRESTApi/package"

	"github.com/jmoiron/sqlx"
)

type Autorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type ToDoList interface {
}

type ToDoItems interface {
}

type Repository struct {
	Autorization
	ToDoList
	ToDoItems
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuthService(db),
	}
}
