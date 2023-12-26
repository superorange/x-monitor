package monitor

import "github.com/shirou/gopsutil/v3/net"

func GetNetworkInfo() ([]net.IOCountersStat, error) {
	return net.IOCounters(true)
}
