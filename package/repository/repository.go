package repository

import (
	todo "firstRESTApi/package"

	"github.com/jmoiron/sqlx"
)

type Autorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Delete(userId, listid int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}

type ToDoItems interface {
}

type Repository struct {
	Autorization
	TodoList
	ToDoItems
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autorization: NewAuthService(db),
		TodoList:     NewTodooListPostgres(db),
	}
}
