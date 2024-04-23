package handlers

import "github.com/gofiber/fiber/v2"

type Validator interface {
	ValidateRequestBody(c *fiber.Ctx, i interface{}) error
	ValidateQueryParams(c *fiber.Ctx, i interface{}) error
}
