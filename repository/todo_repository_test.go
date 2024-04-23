package repository

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go-todo-api/domain"
	"go-todo-api/test"
	"testing"
	"time"
)

var todoColumns = []string{"id", "title", "description", "created_at", "updated_at", "deleted_at", "completed_at"}
var countColumns = []string{"count"}

func TestTodoRepository_FindAll(t *testing.T) {
	sqlDB, gormDB, mock := test.CreateMockDatabase()
	defer sqlDB.Close()

	repository := TodoRepository{DB: gormDB}

	t.Run("should return error when failed to fetch todos", func(t *testing.T) {
		_, err := repository.FindAll(domain.PaginationRequest{Page: 1, PerPage: 10})

		assert.NotNil(t, err)
		assert.Equal(t, "Failed to fetch todos", err.Error())
	})

	t.Run("should return todos", func(t *testing.T) {
		rows := sqlmock.NewRows(todoColumns).
			AddRow(1, "title", "description", time.Time{}, time.Time{}, nil, nil).
			AddRow(2, "title", "description", time.Time{}, time.Time{}, nil, nil).
			AddRow(3, "title", "description", time.Time{}, time.Time{}, nil, nil)
		mock.ExpectQuery("SELECT count").WillReturnRows(sqlmock.NewRows(countColumns).AddRow(1))
		mock.ExpectQuery("SELECT").WillReturnRows(rows)

		paginationRequest := domain.PaginationRequest{Page: 1, PerPage: 1}

		response, err := repository.FindAll(paginationRequest)

		assert.Nil(t, err)
		assert.Equal(t, 3, len(response.Data))
		assert.Equal(t, 1, response.Meta.CurrentPage)
		assert.Equal(t, 1, response.Meta.PerPage)
		assert.Equal(t, false, response.Meta.IsEmpty)
	})

	t.Run("should return todos with pagination", func(t *testing.T) {
		rows := sqlmock.NewRows(todoColumns).
			AddRow(1, "title", "description", time.Time{}, time.Time{}, nil, nil)
		mock.ExpectQuery("SELECT count").WillReturnRows(sqlmock.NewRows(countColumns).AddRow(3))
		mock.ExpectQuery("SELECT").WillReturnRows(rows)

		paginationRequest := domain.PaginationRequest{Page: 1, PerPage: 1}

		response, err := repository.FindAll(paginationRequest)

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

func TestTodoRepository_FindById(t *testing.T) {
	sqlDB, gormDB, mock := test.CreateMockDatabase()
	defer sqlDB.Close()

	repository := TodoRepository{DB: gormDB}

	t.Run("should return error when failed to fetch todo", func(t *testing.T) {
		_, err := repository.FindById(1)

		assert.NotNil(t, err)
		assert.Equal(t, "Todo not found", err.Error())
		assert.IsType(t, &fiber.Error{}, err)
	})

	t.Run("should return todo", func(t *testing.T) {
		rows := sqlmock.NewRows(todoColumns).
			AddRow(1, "title", "description", time.Time{}, time.Time{}, nil, nil)
		mock.ExpectQuery("SELECT").WillReturnRows(rows)

		todo, err := repository.FindById(1)

		assert.Nil(t, err)
		assert.NotNil(t, todo)
		assert.Equal(t, 1, int(todo.ID))
	})
}

func TestTodoRepository_FindDeletedById(t *testing.T) {
	sqlDB, gormDB, mock := test.CreateMockDatabase()
	defer sqlDB.Close()

	repository := TodoRepository{DB: gormDB}

	t.Run("should return error when failed to fetch todo", func(t *testing.T) {
		_, err := repository.FindDeletedById(1)

		assert.NotNil(t, err)
		assert.Equal(t, "Deleted todo not found", err.Error())
		assert.IsType(t, &fiber.Error{}, err)
	})

	t.Run("should return todo", func(t *testing.T) {
		rows := sqlmock.NewRows(todoColumns).
			AddRow(1, "title", "description", time.Time{}, time.Time{}, nil, nil)
		mock.ExpectQuery("SELECT").WillReturnRows(rows)

		todo, err := repository.FindDeletedById(1)

		assert.Nil(t, err)
		assert.NotNil(t, todo)
		assert.Equal(t, 1, int(todo.ID))
	})
}

func TestTodoRepository_MarkAsCompleted(t *testing.T) {
	sqlDB, gormDB, mock := test.CreateMockDatabase()
	defer sqlDB.Close()

	repository := TodoRepository{DB: gormDB}

	t.Run("should return error when failed to mark todo as completed", func(t *testing.T) {
		rows := sqlmock.NewRows(todoColumns).
			AddRow(1, "title", "description", time.Time{}, time.Time{}, nil, nil)
		mock.ExpectQuery("SELECT").WillReturnRows(rows)
		_, err := repository.MarkAsCompleted(1)

		assert.NotNil(t, err)
		assert.Equal(t, "Failed to mark todo as completed", err.Error())
	})

	t.Run("should mark todo as completed", func(t *testing.T) {

		rows := sqlmock.NewRows(todoColumns).
			AddRow(1, "title", "description", time.Time{}, time.Time{}, nil, nil)
		mock.ExpectQuery("SELECT").WillReturnRows(rows)
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		_, err := repository.MarkAsCompleted(1)

		assert.Nil(t, err)
	})
}

func TestTodoRepository_MarkAsUncompleted(t *testing.T) {
	sqlDB, gormDB, mock := test.CreateMockDatabase()
	defer sqlDB.Close()

	repository := TodoRepository{DB: gormDB}

	t.Run("should return error when failed to mark todo as uncompleted", func(t *testing.T) {
		rows := sqlmock.NewRows(todoColumns).
			AddRow(1, "title", "description", time.Time{}, time.Time{}, nil, nil)
		mock.ExpectQuery("SELECT").WillReturnRows(rows)
		_, err := repository.MarkAsUncompleted(1)

		assert.NotNil(t, err)
		assert.Equal(t, "Failed to mark todo as uncompleted", err.Error())
	})

	t.Run("should mark todo as uncompleted", func(t *testing.T) {

		rows := sqlmock.NewRows(todoColumns).
			AddRow(1, "title", "description", time.Time{}, time.Time{}, nil, nil)
		mock.ExpectQuery("SELECT").WillReturnRows(rows)
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		_, err := repository.MarkAsUncompleted(1)

		assert.Nil(t, err)
	})
}

func TestTodoRepository_Update(t *testing.T) {
	sqlDB, gormDB, mock := test.CreateMockDatabase()
	defer sqlDB.Close()

	repository := TodoRepository{DB: gormDB}

	t.Run("should return error when failed to update todo", func(t *testing.T) {
		todo := domain.Todo{}
		todo.ID = 1
		_, err := repository.Update(todo)

		assert.NotNil(t, err)
		assert.Equal(t, "Failed to update todo", err.Error())
	})

	t.Run("should update todo", func(t *testing.T) {

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		todo := domain.Todo{Title: "title", Description: sql.NullString{String: "description", Valid: true}}
		todo.ID = 1

		_, err := repository.Update(todo)

		assert.Nil(t, err)
	})
}

func TestTodoRepository_Delete(t *testing.T) {
	sqlDB, gormDB, mock := test.CreateMockDatabase()
	defer sqlDB.Close()

	repository := TodoRepository{DB: gormDB}

	t.Run("should return error when failed to delete todo", func(t *testing.T) {
		err := repository.Delete(1)

		assert.NotNil(t, err)
		assert.Equal(t, "Failed to delete todo", err.Error())
	})

	t.Run("should delete todo", func(t *testing.T) {

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repository.Delete(1)

		assert.Nil(t, err)
	})
}

func TestTodoRepository_Create(t *testing.T) {
	sqlDB, gormDB, mock := test.CreateMockDatabase()
	defer sqlDB.Close()

	repository := TodoRepository{DB: gormDB}

	t.Run("should return error when failed to create todo", func(t *testing.T) {
		todo := domain.Todo{}
		_, err := repository.Create(todo)

		assert.NotNil(t, err)
		assert.Equal(t, "Failed to create todo", err.Error())
	})

	t.Run("should create todo", func(t *testing.T) {

		mock.ExpectBegin()
		mock.ExpectQuery("INSERT").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		todo := domain.Todo{Title: "title", Description: sql.NullString{String: "description", Valid: true}}
		_, err := repository.Create(todo)

		assert.Nil(t, err)
	})
}

func TestTodoRepository_FindAllDeleted(t *testing.T) {
	sqlDB, gormDB, mock := test.CreateMockDatabase()
	defer sqlDB.Close()

	repository := TodoRepository{DB: gormDB}

	t.Run("should return error when failed to fetch deleted todos", func(t *testing.T) {
		_, err := repository.FindAllDeleted(domain.PaginationRequest{Page: 1, PerPage: 10})

		assert.NotNil(t, err)
		assert.Equal(t, "Failed to fetch deleted todos", err.Error())
	})

	t.Run("should return deleted todos", func(t *testing.T) {
		rows := sqlmock.NewRows(todoColumns).
			AddRow(1, "title", "description", time.Time{}, time.Time{}, time.Time{}, nil).
			AddRow(2, "title", "description", time.Time{}, time.Time{}, time.Time{}, nil).
			AddRow(3, "title", "description", time.Time{}, time.Time{}, time.Time{}, nil)
		mock.ExpectQuery("SELECT count").WillReturnRows(sqlmock.NewRows(countColumns).AddRow(1))
		mock.ExpectQuery("SELECT").WillReturnRows(rows)

		paginationRequest := domain.PaginationRequest{Page: 1, PerPage: 1}

		response, err := repository.FindAllDeleted(paginationRequest)

		assert.Nil(t, err)
		assert.Equal(t, 3, len(response.Data))
		assert.Equal(t, 1, response.Meta.CurrentPage)
		assert.Equal(t, 1, response.Meta.PerPage)
		assert.Equal(t, false, response.Meta.IsEmpty)
	})

	t.Run("should return deleted todos with pagination", func(t *testing.T) {
		rows := sqlmock.NewRows(todoColumns).
			AddRow(1, "title", "description", time.Time{}, time.Time{}, time.Time{}, nil)
		mock.ExpectQuery("SELECT count").WillReturnRows(sqlmock.NewRows(countColumns).AddRow(3))
		mock.ExpectQuery("SELECT").WillReturnRows(rows)

		paginationRequest := domain.PaginationRequest{Page: 1, PerPage: 1}

		response, err := repository.FindAllDeleted(paginationRequest)

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

	t.Run("should return deleted todos with pagination", func(t *testing.T) {
		rows := sqlmock.NewRows(todoColumns).
			AddRow(1, "title", "description", time.Time{}, time.Time{}, time.Time{}, nil)
		mock.ExpectQuery("SELECT count").WillReturnRows(sqlmock.NewRows(countColumns).AddRow(3))
		mock.ExpectQuery("SELECT").WillReturnRows(rows)

		paginationRequest := domain.PaginationRequest{Page: 1, PerPage: 1}

		response, err := repository.FindAllDeleted(paginationRequest)

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

func TestTodoRepository_Recover(t *testing.T) {
	sqlDB, gormDB, mock := test.CreateMockDatabase()
	defer sqlDB.Close()

	repository := TodoRepository{DB: gormDB}

	t.Run("should return error when failed to recover todo", func(t *testing.T) {
		err := repository.Recover(1)

		assert.NotNil(t, err)
		assert.Equal(t, "Deleted todo not found", err.Error())
	})

	t.Run("should recover todo", func(t *testing.T) {
		rows := sqlmock.NewRows(todoColumns).
			AddRow(1, "title", "description", time.Time{}, time.Time{}, time.Time{}, nil)
		mock.ExpectQuery("SELECT").WillReturnRows(rows)

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repository.Recover(1)

		assert.Nil(t, err)
	})
}
