package domain

import (
	"database/sql"
)

type Todo struct {
	BaseModel
	CompletedAt sql.NullTime   `gorm:"index" json:"completed_at"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
}

type TodoRepository interface {
	FindAll(paginationRequest PaginationRequest) (*TodoPaginatedResponse, error)
	FindById(id int) (Todo, error)
	Create(todo Todo) (Todo, error)
	Update(todo Todo) (Todo, error)
	Delete(id int) error
	MarkAsCompleted(id int) (Todo, error)
	MarkAsUncompleted(id int) (Todo, error)
	FindAllDeleted(paginationRequest PaginationRequest) (*TodoPaginatedResponse, error)
	FindDeletedById(id int) (Todo, error)
	Recover(id int) error
}

type TodoService interface {
	FindAll(paginationRequest PaginationRequest) (*TodoPaginatedResponse, error)
	FindById(id int) (Todo, error)
	Create(request CreateOrUpdateTodoRequest) (Todo, error)
	Update(id int, request CreateOrUpdateTodoRequest) (Todo, error)
	Delete(id int) error
	MarkAsCompleted(id int) (Todo, error)
	MarkAsUncompleted(id int) (Todo, error)
	FindAllDeleted(paginationRequest PaginationRequest) (*TodoPaginatedResponse, error)
	Recover(id int) error
}

type CreateOrUpdateTodoRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
}

type TodoPaginatedResponse struct {
	Meta PaginationMetaResponse `json:"meta"`
	Data []Todo                 `json:"data"`
}
