package handlers

import (
	"github.com/gofiber/fiber/v2"
	"go-todo-api/bootstrap"
	"go-todo-api/domain"
)

type TodoHandler struct {
	Container *bootstrap.Container
	V         Validator
}

func NewTodoHandler(container *bootstrap.Container, v Validator) TodoHandler {
	return TodoHandler{Container: container, V: v}
}

func (handler TodoHandler) GetTodos(c *fiber.Ctx) error {
	var request domain.PaginationRequest
	err := handler.V.ValidateQueryParams(c, &request)

	if err != nil {
		return err
	}

	result, err := handler.Container.TodoService.FindAll(request)

	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (handler TodoHandler) GetTodoById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Please provide a numeric id")
	}

	result, err := handler.Container.TodoService.FindById(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(result)
}

func (handler TodoHandler) CreateTodo(c *fiber.Ctx) error {
	var request domain.CreateOrUpdateTodoRequest
	err := handler.V.ValidateRequestBody(c, &request)

	if err != nil {
		return err
	}

	result, err := handler.Container.TodoService.Create(request)

	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (handler TodoHandler) UpdateTodoById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Please provide a numeric id")
	}

	var request domain.CreateOrUpdateTodoRequest
	err = handler.V.ValidateRequestBody(c, &request)

	if err != nil {
		return err
	}

	result, err := handler.Container.TodoService.Update(id, request)

	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (handler TodoHandler) DeleteTodoById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Please provide a numeric id")
	}

	err = handler.Container.TodoService.Delete(id)

	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (handler TodoHandler) MarkAsCompleted(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Please provide a numeric id")
	}

	result, err := handler.Container.TodoService.MarkAsCompleted(id)

	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (handler TodoHandler) MarkAsUncompleted(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Please provide a numeric id")
	}

	result, err := handler.Container.TodoService.MarkAsUncompleted(id)

	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (handler TodoHandler) GetDeletedTodos(c *fiber.Ctx) error {
	var request domain.PaginationRequest
	err := handler.V.ValidateQueryParams(c, &request)

	if err != nil {
		return err
	}

	result, err := handler.Container.TodoService.FindAllDeleted(request)

	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (handler TodoHandler) RecoverTodo(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Please provide a numeric id")
	}

	err = handler.Container.TodoService.Recover(id)

	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}
