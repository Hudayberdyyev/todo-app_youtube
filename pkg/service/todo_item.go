package service

import (
	"github.com/Hudayberdyyev/todo-app-Maksim-"
	"github.com/Hudayberdyyev/todo-app-Maksim-/pkg/repository"
)

type TodoItemService struct {
	repos     repository.TodoItem
	listRepos repository.TodoList
}

func NewTodoItemService(repos repository.TodoItem, listRepos repository.TodoList) *TodoItemService {
	return &TodoItemService{
		repos:     repos,
		listRepos: listRepos,
	}
}

func (s *TodoItemService) Create(userId, listId int, input todo.TodoItem) (int, error) {
	_, err := s.listRepos.GetById(userId, listId)
	if err != nil {
		// list does not exists or does not belongs to user
		return 0, err
	}
	return s.repos.Create(listId, input)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return s.repos.GetAll(userId, listId)
}

func (s *TodoItemService) GetById(userId, listId, itemId int) (todo.TodoItem, error) {
	_, err := s.listRepos.GetById(userId, listId)

	if err != nil {
		// list does not exists or does not belongs to user
		return todo.TodoItem{}, err
	}

	return s.repos.GetById(userId, listId, itemId)
}
