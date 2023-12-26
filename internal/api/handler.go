package api

import (
	"github.com/gofiber/fiber/v2"
	"x-monitor/internal/monitor"
)

func GetCPU(c *fiber.Ctx) error {
	return HandlerError(c, monitor.GetCPUUsage)
}

func GetMemory(c *fiber.Ctx) error {
	return HandlerError(c, monitor.GetMemoryUsage)
}

func GetDisk(c *fiber.Ctx) error {
	return HandlerError(c, monitor.GetDiskUsage)
}

func GetNetwork(c *fiber.Ctx) error {
	return HandlerError(c, monitor.GetNetworkInfo)
}
func GetSystem(c *fiber.Ctx) error {
	return HandlerData(c, monitor.GetSystemInfo)
}

func GetAll(c *fiber.Ctx) error {
	return HandlerData(c, getAll)
}
func getAll() *fiber.Map {
	cpu, _ := monitor.GetCPUUsage()
	memory, _ := monitor.GetMemoryUsage()
	disk, _ := monitor.GetDiskUsage()
	network, _ := monitor.GetNetworkInfo()
	system := monitor.GetSystemInfo()
	return &fiber.Map{
		"cpu":     cpu,
		"memory":  memory,
		"disk":    disk,
		"network": network,
		"system":  system,
	}
}
