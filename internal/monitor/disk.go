package monitor

import "github.com/shirou/gopsutil/v3/disk"

func GetDiskUsage() ([]disk.UsageStat, error) {
	var usages []disk.UsageStat
	partitions, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}

	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err == nil {
			usages = append(usages, *usage)
		}
	}
	return usages, nil
}
