package monitor

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/process"
	"sort"
)

func GetCPUUsage() (fiber.Map, error) {
	processes, _ := process.Processes()
	var cpuProcesses []fiber.Map
	for _, p := range processes {
		cpuPercent, err := p.CPUPercent()
		if err == nil && cpuPercent >= 0.1 {
			name, err := p.Name()
			memory, _ := p.MemoryInfo()
			numThreads, _ := p.NumThreads()
			cpuTime, _ := p.Times()
			if err == nil {
				cpuProcesses = append(cpuProcesses, fiber.Map{
					"pid":        p.Pid,
					"name":       name,
					"cpuPercent": cpuPercent,
					"memory":     memory,
					"numThreads": numThreads,
					"cpuTime":    cpuTime,
				})
			}
		}
	}

	// 降序排序进程列表
	sort.Slice(cpuProcesses, func(i, j int) bool {
		return cpuProcesses[i]["cpuPercent"].(float64) > cpuProcesses[j]["cpuPercent"].(float64)
	})
	info, _ := cpu.Info()
	avg, _ := load.Avg()
	time, _ := cpu.Times(true)
	return fiber.Map{
		"cpuInfo":      info,
		"cpuAvg":       avg,
		"cpuTime":      time,
		"cpuProcesses": cpuProcesses,
	}, nil
}
