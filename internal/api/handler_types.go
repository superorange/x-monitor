package api

import "github.com/gofiber/fiber/v2"

type HandlerWithErrorCall[T any] func() (T, error)
type HandlerWithDataCall[T any] func() T

func HandlerError[T any](ctx *fiber.Ctx, handler HandlerWithErrorCall[T]) error {
	result, err := handler()
	if err != nil {
		return NewApiError(-1, err)
	}
	return ctx.JSON(NewApiResponse(0, result))
}

func HandlerData[T any](ctx *fiber.Ctx, handler HandlerWithDataCall[T]) error {
	return ctx.JSON(NewApiResponse(0, handler()))
}
