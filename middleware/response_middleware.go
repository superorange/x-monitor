package middleware

import (
	"github.com/gofiber/fiber/v2"
	"x-monitor/internal/api"
)

func ErrorHandler(c *fiber.Ctx, err error) error {
	rsp := api.Response{
		Code:    fiber.StatusInternalServerError,
		Message: "数据获取失败",
	}

	if e, ok := err.(*api.Error); ok {
		rsp.Code = e.StatusCode
		rsp.Message = e.Message
		rsp.Error = e.Err
	}
	// 可在此处添加错误日志记录
	return c.Status(200).JSON(&rsp)
}

func ResponseMiddleware(c *fiber.Ctx) error {
	err := c.Next()
	if err != nil {
		return ErrorHandler(c, err)
	}
	return nil
}
