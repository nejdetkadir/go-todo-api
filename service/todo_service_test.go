package service

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go-todo-api/domain"
	"go-todo-api/repository"
	"go-todo-api/test"
	"testing"
	"time"
)

var todoColumns = []string{"id", "title", "description", "created_at", "updated_at", "deleted_at", "completed_at"}

func TestTodoService_Create(t *testing.T) {
	sqlDB, gormDB, mock := test.CreateMockDatabase()
	defer sqlDB.Close()

	todoRepository := repository.TodoRepository{DB: gormDB}

	todoService := NewTodoService(todoRepository)

	t.Run("should create todo", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		request := domain.CreateOrUpdateTodoRequest{
			Title: "Title",
		}

		todo, err := todoService.Create(request)

		assert.Nil(t, err)
		assert.Equal(t, 1, int(todo.ID))
	})

	t.Run("should return error if something wrong", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery("INSERT").WillReturnError(sqlmock.ErrCancelled)
		mock.ExpectRollback()

		request := domain.CreateOrUpdateTodoRequest{
			Title: "Title",
		}

		_, err := todoService.Create(request)

		assert.NotNil(t, err)
		assert.Equal(t, "Failed to create todo", err.Error())
		assert.IsType(t, &fiber.Error{}, err)
	})
}

func TestTodoService_Update(t *testing.T) {
	sqlDB, gormDB, mock := test.CreateMockDatabase()
	defer sqlDB.Close()

	todoRepository := repository.TodoRepository{DB: gormDB}

	todoService := NewTodoService(todoRepository)

	t.Run("should update todo", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(1, "Title"))
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		request := domain.CreateOrUpdateTodoRequest{
			Title: "Title",
		}

		todo, err := todoService.Update(1, request)

		assert.Nil(t, err)
		assert.Equal(t, 1, int(todo.ID))
	})

	t.Run("should return error if todo not found", func(t *testing.T) {
		request := domain.CreateOrUpdateTodoRequest{
			Title: "Title",
		}

		_, err := todoService.Update(1, request)

		assert.NotNil(t, err)
		assert.Equal(t, "Todo not found", err.Error())
		assert.IsType(t, &fiber.Error{}, err)
	})

	t.Run("should return error if something wrong", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(11, "Title"))

		request := domain.CreateOrUpdateTodoRequest{
			Title: "Title",
		}

		_, err := todoService.Update(11, request)

		assert.NotNil(t, err)
		assert.Equal(t, "Failed to update todo", err.Error())
		assert.IsType(t, &fiber.Error{}, err)
	})
}

func TestTodoService_Delete(t *testing.T) {
	sqlDB, gormDB, mock := test.CreateMockDatabase()
	defer sqlDB.Close()

	todoRepository := repository.TodoRepository{DB: gormDB}

	todoService := NewTodoService(todoRepository)

	t.Run("should delete todo", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(1, "Title"))
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := todoService.Delete(1)

		assert.Nil(t, err)
	})

	t.Run("should return error if todo not found", func(t *testing.T) {
		err := todoService.Delete(1)

		assert.NotNil(t, err)
		assert.Equal(t, "Todo not found", err.Error())
		assert.IsType(t, &fiber.Error{}, err)
	})

	t.Run("should return error if something wrong", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(11, "Title"))

		err := todoService.Delete(11)

		assert.NotNil(t, err)
		assert.Equal(t, "Failed to delete todo", err.Error())
		assert.IsType(t, &fiber.Error{}, err)
	})
}

func TestTodoService_MarkAsCompleted(t *testing.T) {
	sqlDB, gormDB, mock := test.CreateMockDatabase()
	defer sqlDB.Close()

	todoRepository := repository.TodoRepository{DB: gormDB}

	todoService := NewTodoService(todoRepository)

	t.Run("should mark todo as completed", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(1, "Title"))
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		todo, err := todoService.MarkAsCompleted(1)

		assert.Nil(t, err)
		assert.Equal(t, 1, int(todo.ID))
	})

	t.Run("should return error if todo not found", func(t *testing.T) {
		_, err := todoService.MarkAsCompleted(1)

		assert.NotNil(t, err)
		assert.Equal(t, "Todo not found", err.Error())
		assert.IsType(t, &fiber.Error{}, err)
	})

	t.Run("should return error if something wrong", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(11, "Title"))

		_, err := todoService.MarkAsCompleted(11)

		assert.NotNil(t, err)
		assert.Equal(t, "Failed to mark todo as completed", err.Error())
		assert.IsType(t, &fiber.Error{}, err)
	})
}

func TestTodoService_MarkAsUncompleted(t *testing.T) {
	sqlDB, gormDB, mock := test.CreateMockDatabase()
	defer sqlDB.Close()

	todoRepository := repository.TodoRepository{DB: gormDB}

	todoService := NewTodoService(todoRepository)

	t.Run("should mark todo as uncompleted", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(1, "Title"))
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		todo, err := todoService.MarkAsUncompleted(1)

		assert.Nil(t, err)
		assert.Equal(t, 1, int(todo.ID))
	})

	t.Run("should return error if todo not found", func(t *testing.T) {
		_, err := todoService.MarkAsUncompleted(1)

		assert.NotNil(t, err)
		assert.Equal(t, "Todo not found", err.Error())
		assert.IsType(t, &fiber.Error{}, err)
	})

	t.Run("should return error if something wrong", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(11, "Title"))

		_, err := todoService.MarkAsUncompleted(11)

		assert.NotNil(t, err)
		assert.Equal(t, "Failed to mark todo as uncompleted", err.Error())
		assert.IsType(t, &fiber.Error{}, err)
	})
}

func TestTodoService_Recover(t *testing.T) {
	sqlDB, gormDB, mock := test.CreateMockDatabase()
	defer sqlDB.Close()

	todoRepository := repository.TodoRepository{DB: gormDB}

	todoService := NewTodoService(todoRepository)

	t.Run("should recover todo", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(1, "Title"))
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := todoService.Recover(1)

		assert.Nil(t, err)
	})

	t.Run("should return error if todo not found", func(t *testing.T) {
		err := todoService.Recover(1)

		assert.NotNil(t, err)
		assert.Equal(t, "Deleted todo not found", err.Error())
		assert.IsType(t, &fiber.Error{}, err)
	})

	t.Run("should return error if something wrong", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(11, "Title"))

		err := todoService.Recover(11)

		assert.NotNil(t, err)
		assert.Equal(t, "Failed to recover todo", err.Error())
		assert.IsType(t, &fiber.Error{}, err)
	})
}

func TestTodoService_FindById(t *testing.T) {
	sqlDB, gormDB, mock := test.CreateMockDatabase()
	defer sqlDB.Close()

	todoRepository := repository.TodoRepository{DB: gormDB}

	todoService := NewTodoService(todoRepository)

	t.Run("should find todo by id", func(t *testing.T) {
		mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(1, "Title"))

		todo, err := todoService.FindById(1)

		assert.Nil(t, err)
		assert.Equal(t, 1, int(todo.ID))
	})

	t.Run("should return error if todo not found", func(t *testing.T) {
		_, err := todoService.FindById(1)

		assert.NotNil(t, err)
		assert.Equal(t, "Todo not found", err.Error())
		assert.IsType(t, &fiber.Error{}, err)
	})
}

func TestTodoService_FindAll(t *testing.T) {
	sqlDB, gormDB, mock := test.CreateMockDatabase()
	defer sqlDB.Close()

	todoRepository := repository.TodoRepository{DB: gormDB}

	todoService := NewTodoService(todoRepository)

	t.Run("should return error when failed to fetch todos", func(t *testing.T) {
		_, err := todoService.FindAll(domain.PaginationRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "Failed to fetch todos", err.Error())
		assert.IsType(t, &fiber.Error{}, err)
	})

	t.Run("should return todos", func(t *testing.T) {
		rows := sqlmock.NewRows(todoColumns).
			AddRow(1, "title", "description", time.Time{}, time.Time{}, nil, nil).
			AddRow(2, "title", "description", time.Time{}, time.Time{}, nil, nil).
			AddRow(3, "title", "description", time.Time{}, time.Time{}, nil, nil)
		mock.ExpectQuery("SELECT count").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(3))
		mock.ExpectQuery("SELECT").WillReturnRows(rows)

		paginationRequest := domain.PaginationRequest{Page: 1, PerPage: 1}

		response, err := todoService.FindAll(paginationRequest)

		assert.Nil(t, err)
		assert.Equal(t, 3, len(response.Data))
		assert.Equal(t, 1, response.Meta.CurrentPage)
		assert.Equal(t, 1, response.Meta.PerPage)
		assert.Equal(t, false, response.Meta.IsEmpty)
	})

	t.Run("should return todos with pagination", func(t *testing.T) {
		rows := sqlmock.NewRows(todoColumns).
			AddRow(1, "title", "description", time.Time{}, time.Time{}, nil, nil)
		mock.ExpectQuery("SELECT count").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(3))
		mock.ExpectQuery("SELECT").WillReturnRows(rows)

		paginationRequest := domain.PaginationRequest{Page: 1, PerPage: 1}

		response, err := todoService.FindAll(paginationRequest)

		var totalPagesCount int
		totalPagesCount = response.Meta.TotalPagesCount

		assert.Nil(t, err)
		assert.Equal(t, 1, len(response.Data))
		assert.Equal(t, 1, response.Meta.CurrentPage)
		assert.Equal(t, 1, response.Meta.PerPage)
		assert.Equal(t, false, response.Meta.IsEmpty)
		assert.Equal(t, 3, response.Meta.TotalPagesCount)
		assert.Equal(t, 3, response.Meta.TotalCount)
		assert.Equal(t, 3, totalPagesCount)
		assert.Nil(t, response.Meta.PrevPage)
	})
}

func TestTodoService_FindAllDeleted(t *testing.T) {
	sqlDB, gormDB, mock := test.CreateMockDatabase()
	defer sqlDB.Close()

	todoRepository := repository.TodoRepository{DB: gormDB}

	todoService := NewTodoService(todoRepository)

	t.Run("should return error when failed to fetch deleted todos", func(t *testing.T) {
		_, err := todoService.FindAllDeleted(domain.PaginationRequest{})

		assert.NotNil(t, err)
		assert.Equal(t, "Failed to fetch deleted todos", err.Error())
		assert.IsType(t, &fiber.Error{}, err)
	})

	t.Run("should return deleted todos", func(t *testing.T) {
		rows := sqlmock.NewRows(todoColumns).
			AddRow(1, "title", "description", time.Time{}, time.Time{}, time.Time{}, nil).
			AddRow(2, "title", "description", time.Time{}, time.Time{}, time.Time{}, nil).
			AddRow(3, "title", "description", time.Time{}, time.Time{}, time.Time{}, nil)
		mock.ExpectQuery("SELECT count").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(3))
		mock.ExpectQuery("SELECT").WillReturnRows(rows)

		paginationRequest := domain.PaginationRequest{Page: 1, PerPage: 1}

		response, err := todoService.FindAllDeleted(paginationRequest)

		assert.Nil(t, err)
		assert.Equal(t, 3, len(response.Data))
		assert.Equal(t, 1, response.Meta.CurrentPage)
		assert.Equal(t, 1, response.Meta.PerPage)
		assert.Equal(t, false, response.Meta.IsEmpty)
	})

	t.Run("should return deleted todos with pagination", func(t *testing.T) {
		rows := sqlmock.NewRows(todoColumns).
			AddRow(1, "title", "description", time.Time{}, time.Time{}, time.Time{}, nil)
		mock.ExpectQuery("SELECT count").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(3))
		mock.ExpectQuery("SELECT").WillReturnRows(rows)

		paginationRequest := domain.PaginationRequest{Page: 1, PerPage: 1}

		response, err := todoService.FindAllDeleted(paginationRequest)

		var totalPagesCount int
		totalPagesCount = response.Meta.TotalPagesCount

		assert.Nil(t, err)
		assert.Equal(t, 1, len(response.Data))
		assert.Equal(t, 1, response.Meta.CurrentPage)
		assert.Equal(t, 1, response.Meta.PerPage)
		assert.Equal(t, false, response.Meta.IsEmpty)
		assert.Equal(t, 3, response.Meta.TotalPagesCount)
		assert.Equal(t, 3, response.Meta.TotalCount)
		assert.Equal(t, 3, totalPagesCount)
		assert.Nil(t, response.Meta.PrevPage)
	})
}
