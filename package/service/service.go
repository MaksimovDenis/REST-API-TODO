package service

import (
	todo "firstRESTApi/package"
	"firstRESTApi/package/repository"
)

type Autorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listid int) (todo.TodoList, error)
	Delete(userId, listid int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}

type ToDoItems interface {
}

type Service struct {
	Autorization
	TodoList
	ToDoItems
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
		TodoList:     NewTodoListService(repos.TodoList),
	}
}
