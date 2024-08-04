package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server running at port 3000")
    http.ListenAndServe(":3000", nil)
}

/*
❯ wrk -t12 -c400 -d30s http://localhost:3000/

Running 30s test @ http://localhost:3000/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.74ms    1.44ms  39.14ms   86.93%
    Req/Sec     6.90k     1.73k   59.76k    95.33%
  2475846 requests in 30.04s, 309.31MB read
  Socket errors: connect 0, read 390, write 0, timeout 0
Requests/sec:  82419.68
Transfer/sec:     10.30MB
*/