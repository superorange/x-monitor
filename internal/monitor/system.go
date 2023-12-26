package monitor

import (
	"runtime"
	"time"
)

type SystemInfo struct {
	GoVersion    string
	NumGoroutine int
	NumCPU       int
	StartTime    time.Time
}

var startTime = time.Now()

func GetSystemInfo() *SystemInfo {
	return &SystemInfo{
		GoVersion:    runtime.Version(),
		NumGoroutine: runtime.NumGoroutine(),
		NumCPU:       runtime.NumCPU(),
		StartTime:    startTime,
	}
}
