package service

import "firstRESTApi/package/repository"

type Autorisation interface {
}

type ToDoList interface {
}

type ToDoItems interface {
}

type Service struct {
	Autorisation
	ToDoList
	ToDoItems
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
