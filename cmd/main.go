package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"image"
	"image/jpeg"
	"log"
	"os"
	"strings"
	"time"
	"x-monitor/internal/api"
	"x-monitor/middleware"
)

func setupRoutes(app *fiber.App) {

	app.Static("/", "./static")

	app.Post("/images", func(ctx *fiber.Ctx) error {

		var maps map[string]string
		err := json.Unmarshal(ctx.Body(), &maps)
		if err != nil {
			return err
		}
		imageData := maps["image"]

		dataURLPrefix := "data:image/jpeg;base64,"
		if strings.HasPrefix(imageData, dataURLPrefix) {
			imageData = strings.TrimPrefix(imageData, dataURLPrefix)
		}

		reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imageData))
		img, _, err := image.Decode(reader)
		if err != nil {
			panic(err)
		}

		// 如果需要，可以对图像进行缩放或其他处理
		// img = resize.Resize(100, 0, img, resize.Lanczos3) // 例子：缩小图像的宽度到100像素

		imageDir := "images"
		if _, err := os.Stat(imageDir); os.IsNotExist(err) {
			_ = os.Mkdir(imageDir, 0755)

		}
		// 创建目标文件
		outFile, err := os.Create(fmt.Sprintf("./images/%d.jpeg", time.Now().Unix())) // 根据需要选择图像格式，这里是JPEG
		if err != nil {
			panic(err)
		}
		defer func(outFile *os.File) {
			_ = outFile.Close()
		}(outFile)

		// 将图像写入文件
		err = jpeg.Encode(outFile, img, nil) // 或者 png.Encode(outFile, img, nil)，根据图像格式选择
		if err != nil {
			panic(err)
		}

		println("图像已保存为 output_image.jpg")

		return ctx.Status(200).JSON(fiber.Map{
			"code": 200,
		})

	})
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
