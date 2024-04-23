package service

import (
	"database/sql"
	"go-todo-api/domain"
)

type TodoService struct {
	TodoRepository domain.TodoRepository
}

func NewTodoService(todoRepository domain.TodoRepository) domain.TodoService {
	return TodoService{TodoRepository: todoRepository}
}

func (s TodoService) FindAll(paginationRequest domain.PaginationRequest) (*domain.TodoPaginatedResponse, error) {
	return s.TodoRepository.FindAll(paginationRequest)
}

func (s TodoService) FindAllDeleted(paginationRequest domain.PaginationRequest) (*domain.TodoPaginatedResponse, error) {
	return s.TodoRepository.FindAllDeleted(paginationRequest)
}

func (s TodoService) FindById(id int) (domain.Todo, error) {
	return s.TodoRepository.FindById(id)
}

func (s TodoService) Create(request domain.CreateOrUpdateTodoRequest) (domain.Todo, error) {
	todo := domain.Todo{
		Title:       request.Title,
		Description: sql.NullString{String: request.Description, Valid: request.Description != ""},
	}

	return s.TodoRepository.Create(todo)
}

func (s TodoService) Update(id int, request domain.CreateOrUpdateTodoRequest) (domain.Todo, error) {
	todo, err := s.TodoRepository.FindById(id)

	if err != nil {
		return domain.Todo{}, err
	}

	todo.Title = request.Title
	todo.Description = sql.NullString{String: request.Description, Valid: request.Description != ""}

	return s.TodoRepository.Update(todo)
}

func (s TodoService) Delete(id int) error {
	todo, err := s.TodoRepository.FindById(id)

	if err != nil {
		return err
	}

	return s.TodoRepository.Delete(int(todo.ID))
}

func (s TodoService) MarkAsCompleted(id int) (domain.Todo, error) {
	return s.TodoRepository.MarkAsCompleted(id)
}

func (s TodoService) MarkAsUncompleted(id int) (domain.Todo, error) {
	return s.TodoRepository.MarkAsUncompleted(id)
}

func (s TodoService) Recover(id int) error {
	return s.TodoRepository.Recover(id)
}
