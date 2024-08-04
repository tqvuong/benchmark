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

func cpuTest(durationSeconds int) []float64 {
    start := time.Now()
    duration := time.Duration(durationSeconds) * time.Second
    var cpuUsage []float64

    for time.Since(start) < duration {
        startCPU := runtime.NumCPU()
        startTime := time.Now()
        task()
        elapsed := time.Since(startTime)
        endCPU := runtime.NumCPU()

        cpuPercent := float64(endCPU-startCPU) / elapsed.Seconds() * 100
        cpuUsage = append(cpuUsage, cpuPercent)
    }

    return cpuUsage
}

func main() {
    duration := 10
    cpuUsage := cpuTest(duration)
    totalCpuUsage := 0.0
    for _, usage := range cpuUsage {
        totalCpuUsage += usage
    }
    avgCpuUsage := totalCpuUsage / float64(len(cpuUsage))
    fmt.Printf("Average CPU usage: %.10f%%\n", avgCpuUsage)
}
