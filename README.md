# go-gin-context-conflict

```
% go test -race ./...
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /hello                    --> github.com/haruyama480/go-gin-context-conflict.setupRouter.func1 (3 handlers)
[GIN] 2026/04/05 - 20:48:09 | 200 | 262.208µs |                 | GET      "/hello"
[GIN] 2026/04/05 - 20:48:09 | 200 | 342.25µs |                 | GET      "/hello"
[GIN] 2026/04/05 - 20:48:09 | 200 | 382.375µs |                 | GET      "/hello"
[GIN] 2026/04/05 - 20:48:09 | 200 | 450.042µs |                 | GET      "/hello"
[GIN] 2026/04/05 - 20:48:09 | 200 | 132.167µs |                 | GET      "/hello"
[GIN] 2026/04/05 - 20:48:09 | 200 | 152.458µs |                 | GET      "/hello"
[GIN] 2026/04/05 - 20:48:09 | 200 | 66.084µs |                 | GET      "/hello"
[GIN] 2026/04/05 - 20:48:09 | 200 | 52.708µs |                 | GET      "/hello"
[GIN] 2026/04/05 - 20:48:09 | 200 | 184.708µs |                 | GET      "/hello"
[GIN] 2026/04/05 - 20:48:09 | 200 | 35.167µs |                 | GET      "/hello"
[GIN] 2026/04/05 - 20:48:09 | 200 | 268.959µs |                 | GET      "/hello"
[GIN] 2026/04/05 - 20:48:09 | 200 | 26.084µs |                 | GET      "/hello"
==================
WARNING: DATA RACE
Write at 0x00c00020e020 by goroutine 23:
  github.com/gin-gonic/gin.(*Engine).ServeHTTP()
      /Users/kazusaku/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/gin.go:669 +0x100
  github.com/haruyama480/go-gin-context-conflict.TestHelloRouteParallel.func1()
      /Users/kazusaku/ghq/github.com/haruyama480/go-gin-context-conflict/main_test.go:22 +0x164
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1934 +0x164
  testing.(*T).Run.gowrap1()
      /usr/local/go/src/testing/testing.go:1997 +0x3c

Previous read at 0x00c00020e020 by goroutine 29:
  github.com/gin-gonic/gin.(*Context).hasRequestContext()
      /Users/kazusaku/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/context.go:1436 +0x88
  github.com/gin-gonic/gin.(*Context).Done()
      /Users/kazusaku/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/context.go:1450 +0x104
  github.com/haruyama480/go-gin-context-conflict.delayedFunc()
      /Users/kazusaku/ghq/github.com/haruyama480/go-gin-context-conflict/main.go:18 +0x5c
  github.com/haruyama480/go-gin-context-conflict.setupRouter.func1.gowrap1()
      /Users/kazusaku/ghq/github.com/haruyama480/go-gin-context-conflict/main.go:34 +0x3c

Goroutine 23 (running) created at:
  testing.(*T).Run()
      /usr/local/go/src/testing/testing.go:1997 +0x6e0
  github.com/haruyama480/go-gin-context-conflict.TestHelloRouteParallel()
      /Users/kazusaku/ghq/github.com/haruyama480/go-gin-context-conflict/main_test.go:16 +0x4c
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1934 +0x164
  testing.(*T).Run.gowrap1()
      /usr/local/go/src/testing/testing.go:1997 +0x3c

Goroutine 29 (running) created at:
  github.com/haruyama480/go-gin-context-conflict.setupRouter.func1()
      /Users/kazusaku/ghq/github.com/haruyama480/go-gin-context-conflict/main.go:34 +0x124
  github.com/gin-gonic/gin.(*Context).Next()
      /Users/kazusaku/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/context.go:192 +0x17c
  github.com/gin-gonic/gin.CustomRecoveryWithWriter.func1()
      /Users/kazusaku/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/recovery.go:90 +0x90
  github.com/gin-gonic/gin.(*Context).Next()
      /Users/kazusaku/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/context.go:192 +0x17c
  github.com/gin-gonic/gin.LoggerWithConfig.func1()
      /Users/kazusaku/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/logger.go:282 +0x104
  github.com/gin-gonic/gin.(*Context).Next()
      /Users/kazusaku/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/context.go:192 +0x17c
  github.com/gin-gonic/gin.(*Engine).handleHTTPRequest()
      /Users/kazusaku/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/gin.go:722 +0x574
  github.com/gin-gonic/gin.(*Engine).ServeHTTP()
      /Users/kazusaku/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/gin.go:672 +0x368
  github.com/haruyama480/go-gin-context-conflict.TestHelloRouteParallel.func1()
      /Users/kazusaku/ghq/github.com/haruyama480/go-gin-context-conflict/main_test.go:22 +0x164
  testing.tRunner()
      /usr/local/go/src/testing/testing.go:1934 +0x164
  testing.(*T).Run.gowrap1()
      /usr/local/go/src/testing/testing.go:1997 +0x3c
==================
[GIN] 2026/04/05 - 20:48:09 | 200 | 53.917µs |                 | GET      "/hello"
[GIN] 2026/04/05 - 20:48:09 | 200 | 29.584µs |                 | GET      "/hello"
[GIN] 2026/04/05 - 20:48:09 | 200 | 44.958µs |                 | GET      "/hello"
[GIN] 2026/04/05 - 20:48:09 | 200 | 28.584µs |                 | GET      "/hello"
[GIN] 2026/04/05 - 20:48:09 | 200 | 254.666µs |                 | GET      "/hello"
[GIN] 2026/04/05 - 20:48:09 | 200 | 705.834µs |                 | GET      "/hello"
[GIN] 2026/04/05 - 20:48:09 | 200 | 55.958µs |                 | GET      "/hello"
[GIN] 2026/04/05 - 20:48:09 | 200 |   1.88ms |                 | GET      "/hello"
--- FAIL: TestHelloRouteParallel (0.00s)
    --- FAIL: TestHelloRouteParallel/ParallelRequest#14 (0.00s)
        testing.go:1617: race detected during execution of test
    --- FAIL: TestHelloRouteParallel/ParallelRequest#07 (0.00s)
        testing.go:1617: race detected during execution of test
    --- FAIL: TestHelloRouteParallel/ParallelRequest#10 (0.00s)
        testing.go:1617: race detected during execution of test
    --- FAIL: TestHelloRouteParallel/ParallelRequest#13 (0.00s)
        testing.go:1617: race detected during execution of test
    --- FAIL: TestHelloRouteParallel/ParallelRequest#01 (0.00s)
        testing.go:1617: race detected during execution of test
    --- FAIL: TestHelloRouteParallel/ParallelRequest#11 (0.00s)
        testing.go:1617: race detected during execution of test
    --- FAIL: TestHelloRouteParallel/ParallelRequest#03 (0.00s)
        testing.go:1617: race detected during execution of test
    --- FAIL: TestHelloRouteParallel/ParallelRequest#17 (0.00s)
        testing.go:1617: race detected during execution of test
    --- FAIL: TestHelloRouteParallel/ParallelRequest#09 (0.00s)
        testing.go:1617: race detected during execution of test
    --- FAIL: TestHelloRouteParallel/ParallelRequest#04 (0.00s)
        testing.go:1617: race detected during execution of test
    --- FAIL: TestHelloRouteParallel/ParallelRequest (0.00s)
        testing.go:1617: race detected during execution of test
    --- FAIL: TestHelloRouteParallel/ParallelRequest#19 (0.00s)
        testing.go:1617: race detected during execution of test
FAIL
FAIL    github.com/haruyama480/go-gin-context-conflict  0.301s
FAIL
```

## the reason
gin.Context is not goroutine-safe. It is reused for a new another request.

see:
- https://github.com/gin-gonic/gin/issues/4117
- https://engineering.nifty.co.jp/blog/35119

## how to fix

There are two options.

```
% perl -i -pe 's{^(\s*)(//\s*)?(.*TOGGLE_A)}{$2 ? "$1$3" : "$1// $3"}e' main.go
% go test -race ./...
ok      github.com/haruyama480/go-gin-context-conflict  1.621s
```

```
% perl -i -pe 's{^(\s*)(//\s*)?(.*TOGGLE_B)}{$2 ? "$1$3" : "$1// $3"}e' main.go
% go test -race ./...
ok      github.com/haruyama480/go-gin-context-conflict  2.601s
```
