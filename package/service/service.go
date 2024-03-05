package service

import (
	todo "firstRESTApi/package"
	"firstRESTApi/package/repository"
)

type Autorization interface {
	CreateUser(user todo.User) (int, error)
}

type ToDoList interface {
}

type ToDoItems interface {
}

type Service struct {
	Autorization
	ToDoList
	ToDoItems
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Autorization: NewAuthService(repos.Autorization),
	}
}
