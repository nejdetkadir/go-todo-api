package routes

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-todo-api/bootstrap"
)

type (
	CustomValidator struct {
		Validator *validator.Validate
	}
	CustomValidatorError struct {
		FailedField string
		Tag         string
		Value       interface{}
		Message     string
	}
)

var goValidator = validator.New()

func Setup(container *bootstrap.Container) {
	apiGroup := container.FiberApp.Group("/api")
	v1 := apiGroup.Group("/v1")

	customValidator := CustomValidator{Validator: goValidator}

	DefineHealthCheckRoutes(container)
	DefineHelloRoutes(v1, container)
	DefineTodoRoutes(v1, container, customValidator)
}

func (cv *CustomValidator) Validate(i interface{}) []CustomValidatorError {
	if err := cv.Validator.Struct(i); err != nil {
		var validationErrors validator.ValidationErrors
		errors.As(err, &validationErrors)
		customErrors := make([]CustomValidatorError, 0)

		for _, e := range validationErrors {
			customErrors = append(customErrors, CustomValidatorError{
				FailedField: e.Field(),
				Tag:         e.Tag(),
				Value:       e.Value(),
				Message:     e.Error(),
			})
		}

		return customErrors
	}

	return nil
}

func (cv *CustomValidator) ValidateRequestBody(c *fiber.Ctx, i interface{}) error {
	if err := c.BodyParser(i); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validationErrors := cv.Validate(i)

	if len(validationErrors) > 0 {
		return fiber.NewError(fiber.StatusBadRequest, validationErrors[0].Message)
	}

	return nil
}

func (cv *CustomValidator) ValidateQueryParams(c *fiber.Ctx, i interface{}) error {
	if err := c.QueryParser(i); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validationErrors := cv.Validate(i)

	if len(validationErrors) > 0 {
		return fiber.NewError(fiber.StatusBadRequest, validationErrors[0].Message)
	}

	return nil
}
