package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetCPU(t *testing.T) {
	app := fiber.New()
	app.Get("/api/v1/cpu", GetCPU)

	req := httptest.NewRequest("GET", "/api/v1/cpu", nil)
	resp, err := app.Test(req, -1)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestGetMemory(t *testing.T) {
	app := fiber.New()
	app.Get("/api/v1/memory", GetMemory)

	req := httptest.NewRequest("GET", "/api/v1/memory", nil)
	resp, err := app.Test(req, -1)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

// ... 类似的测试用例用于磁盘、网络和系统信息接口
