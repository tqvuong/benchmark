package main

import (
    "fmt"
    "sync"
    "time"
)

func task(n int, wg *sync.WaitGroup) {
    defer wg.Done()
    result := 0
    for i := 0; i < 1000*n; i++ {
        result += i
    }
}

func concurrencyTest(numTasks int) float64 {
    var wg sync.WaitGroup
    wg.Add(numTasks)

    startTime := time.Now()

    for i := 0; i < numTasks; i++ {
        go task(i, &wg)
    }

    wg.Wait()
    endTime := time.Now()

    return endTime.Sub(startTime).Seconds()
}

func main() {
    numTasks := 10000
    duration := concurrencyTest(numTasks)
    fmt.Printf("Time taken to complete %d tasks: %.2f seconds\n", numTasks, duration)
}
