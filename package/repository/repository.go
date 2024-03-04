package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
