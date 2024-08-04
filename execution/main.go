// Go Example
package main

import (
    "fmt"
    "time"
)

func task() {}

func main() {
    start := time.Now()
    for i := 0; i < 1000; i++ {
        task()
    }
    end := time.Now()
    fmt.Printf("Execution Time: %v seconds\n", end.Sub(start).Seconds())
}