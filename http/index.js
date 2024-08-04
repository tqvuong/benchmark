const http = require("http");

const server = http.createServer((req, res) => {
  res.statusCode = 200;
  res.setHeader("Content-Type", "text/plain");
  res.end("Hello, World!\n");
});

const PORT = 3000;
server.listen(PORT, () => {
  console.log(`Server running at port ${PORT}`);
});

/*
❯ wrk -t12 -c400 -d30s http://localhost:3000/

Running 30s test @ http://localhost:3000/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.21ms  701.42us  34.47ms   86.95%
    Req/Sec     6.36k   462.67     7.16k    74.78%
  2278952 requests in 30.01s, 354.26MB read
  Socket errors: connect 0, read 388, write 0, timeout 0
Requests/sec:  75928.18
Transfer/sec:     11.80MB
*/
