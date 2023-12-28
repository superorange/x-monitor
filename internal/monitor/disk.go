package monitor

import (
	"github.com/shirou/gopsutil/v3/disk"
)

// DiskPartitionInfo 包含分区信息
type DiskPartitionInfo struct {
	Device      string  `json:"device"`
	Mountpoint  string  `json:"mountpoint"`
	Fstype      string  `json:"fstype"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

// GetDiskUsage 获取物理硬盘及其挂载点信息
func GetDiskUsage() ([]DiskPartitionInfo, error) {
	partitions, err := disk.Partitions(false) // false 表示只列出物理磁盘的分区
	if err != nil {
		return nil, err
	}

	var partitionsInfo []DiskPartitionInfo
	for _, partition := range partitions {
		usageStat, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			continue // 如果无法获取分区使用情况，则跳过此分区
		}

		partitionsInfo = append(partitionsInfo, DiskPartitionInfo{
			Device:      partition.Device,
			Mountpoint:  partition.Mountpoint,
			Fstype:      partition.Fstype,
			Total:       usageStat.Total,
			Free:        usageStat.Free,
			Used:        usageStat.Used,
			UsedPercent: usageStat.UsedPercent,
		})
	}

	return partitionsInfo, nil
}
