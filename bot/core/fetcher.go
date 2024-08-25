package core

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type systemInfo struct {
	CPU           string
	Cores         string
	Memory        string
	Frequency     string
	SystemType    string
	KernelVersion string
	Platform      string
}

var (
	once    sync.Once
	sysInfo systemInfo
)

// GetFetch 获取基本系统信息
func GetFetch() systemInfo {
	once.Do(func() {
		cpuInfo, err := cpu.Info()
		if err != nil || len(cpuInfo) == 0 {
			return
		}

		physicalCores, _ := cpu.Counts(false)
		logicalCores, _ := cpu.Counts(true)
		totalMemory, _ := mem.VirtualMemory()
		hostInfo, _ := host.Info()

		sysInfo = systemInfo{
			CPU:           cpuInfo[0].ModelName,
			Cores:         fmt.Sprintf("%dc%dt", physicalCores, logicalCores),
			Memory:        fmt.Sprintf("%.2f GB", float64(totalMemory.Total)/1e9),
			Frequency:     fmt.Sprintf("%.2f GHz", cpuInfo[0].Mhz/1000),
			SystemType:    runtime.GOOS,
			KernelVersion: hostInfo.KernelVersion,
			Platform:      fmt.Sprintf("%s %s", hostInfo.Platform, hostInfo.PlatformVersion),
		}
	})

	return sysInfo
}
