package repository

import "github.com/jmoiron/sqlx"

type Autorisation interface {
}

type ToDoList interface {
}

type ToDoItems interface {
}

type Repository struct {
	Autorisation
	ToDoList
	ToDoItems
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
