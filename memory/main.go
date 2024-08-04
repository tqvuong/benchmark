package main

import (
    "fmt"
    "runtime"
    "time"
)

func task() {
    result := 0
    for i := 0; i < 1000; i++ {
        result += i
    }
}

func memoryTest() (uint64, uint64, uint64, uint64) {
    var memStats runtime.MemStats
    runtime.ReadMemStats(&memStats)
    beforeMem := memStats.Alloc / 1024 / 1024

    start := time.Now()
    duration := 10 * time.Second // 10 seconds
    memUsage := []uint64{}
    totalMem := uint64(0)

    for time.Since(start) < duration {
        task()
        runtime.ReadMemStats(&memStats)
        mem := memStats.Alloc / 1024 / 1024
        memUsage = append(memUsage, mem)
        totalMem += mem
    }

    runtime.ReadMemStats(&memStats)
    afterMem := memStats.Alloc / 1024 / 1024

    maxMem := uint64(0)
    for _, mem := range memUsage {
        if mem > maxMem {
            maxMem = mem
        }
    }

    return beforeMem, afterMem, maxMem, totalMem
}

func main() {
    beforeMem, afterMem, maxMem, totalMem := memoryTest()
    fmt.Printf("Memory usage before tasks (MB): %d\n", beforeMem)
    fmt.Printf("Memory usage after tasks (MB): %d\n", afterMem)
    fmt.Printf("Max memory usage (MB): %d\n", maxMem)
    fmt.Printf("Total memory usage (MB): %d\n", totalMem)
}
// Memory usage before tasks (MB): 0
// Memory usage after tasks (MB): 5
// Max memory usage (MB): 5
// Total memory usage (MB): 1410998