package main

import (
    "fmt"
    "time"
)

func task() {
    result := 0
    for i := 0; i < 1000; i++ {
        result += i
    }
}

func throughputTest(durationSeconds int) int {
    start := time.Now()
    end := start.Add(time.Duration(durationSeconds) * time.Second)
    count := 0

    for time.Now().Before(end) {
        task()
        count++
    }

    return count
}

func main() {
    duration := 10
    operations := throughputTest(duration)
    fmt.Printf("Operations completed in %d seconds: %d\n", duration, operations)
}
