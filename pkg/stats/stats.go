package stats

import (
	"fmt"
	"runtime"
)

// needs change - return value only
func GetAllocatedMemory() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return fmt.Sprintf("%d KB\n", m.Alloc/1024)
}
