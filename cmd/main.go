package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"x-monitor/internal/api"
	"x-monitor/middleware"
)

func setupRoutes(app *fiber.App) {

	apiRoute := app.Group("/api/v1")

	// 设置路由以处理不同的 API 请求
	apiRoute.Get("/cpu", api.GetCPU)
	apiRoute.Get("/memory", api.GetMemory)
	apiRoute.Get("/disk", api.GetDisk)
	apiRoute.Get("/network", api.GetNetwork)
	apiRoute.Get("/system", api.GetSystem)
	apiRoute.Get("/all", api.GetAll)
}

func main() {
	app := fiber.New() // 初始化 GoFiber 应用
	app.Use(middleware.ResponseMiddleware)
	setupRoutes(app) // 配置路由
	log.Fatal(app.Listen(":8080"))
}
