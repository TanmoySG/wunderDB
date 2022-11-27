package stats

import (
	"fmt"
	"runtime"
)

// needs change - return value only
func GetAllocatedMemory() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return fmt.Sprintf("%v KB", m.Alloc/1024)
}

func GetCpuUsage() int {
	return runtime.NumCPU()
}
