package monitor

import "github.com/shirou/gopsutil/v3/cpu"

func GetCPUUsage() ([]cpu.InfoStat, error) {
	return cpu.Info()
}
