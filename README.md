## Bug report for fasthttp
This app contains two APIs.
Both of them will serve a static file from disk.
`http://localhost:8080/1` will do it by openning the file, reading it and setting it as body.
`http://localhost:8080/2` will do it by using `fasthttp.ServeFile` function.


There are also two images in this repo. `image1.jpg` and `image2.jpg`.
When servering `image1.jpg`, `fasthttp.ServeFile` is much slower.

```
> go-wrk -d 3 "http://localhost:8080/1"

Running 3s test @ http://localhost:8080/1
  10 goroutine(s) running concurrently
31200 requests in 2.923608791s, 1.46GB read
Requests/sec:       10671.74
Transfer/sec:       512.34MB
Avg Req Time:       937.054µs
Fastest Request:    134.841µs
Slowest Request:    7.879186ms
Number of Errors:   0


==========================================

> go-wrk -d 3 "http://localhost:8080/2"

Running 3s test @ http://localhost:8080/2
  10 goroutine(s) running concurrently
248 requests in 3.029993142s, 11.92MB read
Requests/sec:       81.85
Transfer/sec:       3.93MB
Avg Req Time:       122.177142ms
Fastest Request:    118.257µs
Slowest Request:    216.498781ms
Number of Errors:   0

```

But amazingly when you change `image_path` variable in code to `image2.jpg`
Results are relatively close. I cannot understand what is happening.
I've tried to use pprof but I don't know why it does not generate cpu profile
for fasthttp api.
