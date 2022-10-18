package service

import "github.com/Serminaz/GoRun/todo2/pkg/repository"

/* сервисы должны уметь общаться с бд  --> в репозитории работа с бд*/

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

// сервисы обращается к бд
func NewService(repos *repository.Repository) *Service { //  указатель на структуры репозиторий
	return &Service{}
}
