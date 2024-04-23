package repository

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"go-todo-api/domain"
	"gorm.io/gorm"
	"time"
)

type TodoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(app domain.ApplicationType) domain.TodoRepository {
	return TodoRepository{DB: app.GetDB()}
}

func (r TodoRepository) FindAll(paginationRequest domain.PaginationRequest) (*domain.TodoPaginatedResponse, error) {
	var todos []domain.Todo
	var count int64
	query := r.DB.Model(&domain.Todo{}).Where("deleted_at IS NULL")

	err := query.Count(&count).Error

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch todos")
	}

	err = query.Offset(paginationRequest.GetOffset()).Limit(paginationRequest.GetLimit()).Find(&todos).Error

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch todos")
	}

	var meta = domain.PaginationMetaResponse{}.GetPaginationMetaResponse(paginationRequest, int(count), len(todos))
	return &domain.TodoPaginatedResponse{Data: todos, Meta: meta}, nil
}

func (r TodoRepository) FindById(id int) (domain.Todo, error) {
	var todo domain.Todo
	err := r.DB.Model(&domain.Todo{}).Where("id = ? AND deleted_at IS NULL", id).First(&todo).Error

	if err != nil {
		return domain.Todo{}, fiber.NewError(fiber.StatusNotFound, "Todo not found")
	}

	return todo, nil
}

func (r TodoRepository) Create(todo domain.Todo) (domain.Todo, error) {
	err := r.DB.Model(&domain.Todo{}).Create(&todo).Error

	if err != nil {
		return domain.Todo{}, fiber.NewError(fiber.StatusUnprocessableEntity, "Failed to create todo")
	}

	return todo, nil
}

func (r TodoRepository) Update(todo domain.Todo) (domain.Todo, error) {
	err := r.DB.Model(&domain.Todo{}).Where("id = ?", todo.ID).Updates(&todo).Error

	if err != nil {
		return domain.Todo{}, fiber.NewError(fiber.StatusUnprocessableEntity, "Failed to update todo")
	}

	return todo, nil
}

func (r TodoRepository) Delete(id int) error {
	err := r.DB.Model(&domain.Todo{}).Where("id = ?", id).Update("deleted_at", time.Now().UTC()).Error

	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Failed to delete todo")
	}

	return nil
}

func (r TodoRepository) MarkAsCompleted(id int) (domain.Todo, error) {
	todo, err := r.FindById(id)

	if err != nil {
		return domain.Todo{}, err
	}

	todo.CompletedAt = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	dbErr := r.DB.Model(&domain.Todo{}).Where("id = ?", id).Updates(&todo).Error

	if dbErr != nil {
		return domain.Todo{}, fiber.NewError(fiber.StatusUnprocessableEntity, "Failed to mark todo as completed")
	}

	return todo, nil
}

func (r TodoRepository) MarkAsUncompleted(id int) (domain.Todo, error) {
	todo, err := r.FindById(id)

	if err != nil {
		return domain.Todo{}, err
	}

	todo.CompletedAt = sql.NullTime{Time: time.Time{}, Valid: false}
	dbErr := r.DB.Model(&domain.Todo{}).Where("id = ?", id).Updates(&todo).Error

	if dbErr != nil {
		return domain.Todo{}, fiber.NewError(fiber.StatusUnprocessableEntity, "Failed to mark todo as uncompleted")
	}

	return todo, nil
}

func (r TodoRepository) FindAllDeleted(paginationRequest domain.PaginationRequest) (*domain.TodoPaginatedResponse, error) {
	var todos []domain.Todo
	var count int64
	query := r.DB.Model(&domain.Todo{}).Where("deleted_at IS NOT NULL")

	err := query.Count(&count).Error

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch deleted todos")
	}

	err = query.Offset(paginationRequest.GetOffset()).Limit(paginationRequest.GetLimit()).Find(&todos).Error

	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch deleted todos")
	}

	var meta = domain.PaginationMetaResponse{}.GetPaginationMetaResponse(paginationRequest, int(count), len(todos))
	return &domain.TodoPaginatedResponse{Data: todos, Meta: meta}, nil
}

func (r TodoRepository) FindDeletedById(id int) (domain.Todo, error) {
	var todo domain.Todo
	err := r.DB.Model(&domain.Todo{}).Where("id = ? AND deleted_at IS NOT NULL", id).First(&todo).Error

	if err != nil {
		return domain.Todo{}, fiber.NewError(fiber.StatusNotFound, "Deleted todo not found")
	}

	return todo, nil
}

func (r TodoRepository) Recover(id int) error {
	_, err := r.FindDeletedById(id)

	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Deleted todo not found")
	}

	dbErr := r.DB.Model(&domain.Todo{}).Where("id = ?", id).Update("deleted_at", nil).Error

	if dbErr != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "Failed to recover todo")
	}

	return nil
}
