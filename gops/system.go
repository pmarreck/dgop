package gops

import (
	"fmt"
	"sync"
	"time"

	"github.com/AvengeMedia/dgop/models"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/process"
)

type systemTracker struct {
	mu              sync.Mutex
	threadsCachedAt time.Time
	threadCount     int
}

var sysTracker = &systemTracker{}

func (self *GopsUtil) GetSystemInfo() (*models.SystemInfo, error) {
	// System info
	loadAvg, _ := load.Avg()
	procs, _ := process.Pids()
	bootTime, _ := host.BootTime()

	threadCount := self.getThreadCountCached(procs)

	return &models.SystemInfo{
		LoadAvg:   fmt.Sprintf("%.2f %.2f %.2f", loadAvg.Load1, loadAvg.Load5, loadAvg.Load15),
		Processes: len(procs),
		Threads:   threadCount,
		BootTime:  time.Unix(int64(bootTime), 0).Format("2006-01-02 15:04:05"),
	}, nil
}

func (self *GopsUtil) getThreadCountCached(procs []int32) int {
	now := time.Now()

	sysTracker.mu.Lock()
	defer sysTracker.mu.Unlock()

	if now.Sub(sysTracker.threadsCachedAt) < 10*time.Second {
		return sysTracker.threadCount
	}

	threadCount := 0
	for _, pid := range procs {
		proc, err := self.procProvider.NewProcess(pid)
		if err == nil {
			threads, _ := proc.NumThreads()
			threadCount += int(threads)
		}
	}

	sysTracker.threadCount = threadCount
	sysTracker.threadsCachedAt = now
	return threadCount
}
