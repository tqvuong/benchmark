from http.server import BaseHTTPRequestHandler, HTTPServer

class SimpleHTTPRequestHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header('Content-type', 'text/plain')
        self.end_headers()
        self.wfile.write(b'Hello, World!\n')

def run(server_class=HTTPServer, handler_class=SimpleHTTPRequestHandler):
    server_address = ('', 3000)
    httpd = server_class(server_address, handler_class)
    print('Server running at port 3000')
    httpd.serve_forever()

if __name__ == '__main__':
    run()

# ❯ wrk -t12 -c400 -d30s http://localhost:3000/

# Running 30s test @ http://localhost:3000/
#   12 threads and 400 connections
#   Thread Stats   Avg      Stdev     Max   +/- Stdev
#     Latency   255.76us  251.39us   5.55ms   94.93%
#     Req/Sec   348.62    391.32     1.40k    82.11%
#   7124 requests in 30.06s, 0.90MB read
#   Socket errors: connect 0, read 16584, write 22, timeout 0
# Requests/sec:    236.97
# Transfer/sec:     30.55KB