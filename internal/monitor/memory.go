package monitor

import "github.com/shirou/gopsutil/v3/mem"

func GetMemoryUsage() (*mem.VirtualMemoryStat, error) {
	return mem.VirtualMemory()
}
