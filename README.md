# stacktraces

Example of how to get stacktraces for your Go server (with Gin, but you can do this with most routers).

## The trace

```
2021/10/07 09:51:23 Recovered from panic: goroutine 20 [running]:
runtime/debug.Stack()
        /usr/local/go/src/runtime/debug/stack.go:24 +0x65
main.main.func1.1()
        /src/stacktraces/cmd/server/server.go:17 +0x39
panic({0x14d1440, 0x162e760})
        /usr/local/go/src/runtime/panic.go:1038 +0x215
github.com/jakecoffman/stacktraces/service2.DoStuff2(...)
        /src/stacktraces/service2/service.go:4
github.com/jakecoffman/stacktraces/service1.DoStuff1(...)
        /src/stacktraces/service1/service.go:6
main.main.func2(0x2)
        /src/stacktraces/cmd/server/server.go:25 +0x28
github.com/gin-gonic/gin.(*Context).Next(...)
        /Users/jake/go/pkg/mod/github.com/gin-gonic/gin@v1.7.4/context.go:165
main.main.func1(0xc0001d2f00)
        /src/stacktraces/cmd/server/server.go:21 +0x6b
github.com/gin-gonic/gin.(*Context).Next(...)
        /Users/jake/go/pkg/mod/github.com/gin-gonic/gin@v1.7.4/context.go:165
github.com/gin-gonic/gin.(*Engine).handleHTTPRequest(0xc0003911e0, 0xc0001d2f00)
        /Users/jake/go/pkg/mod/github.com/gin-gonic/gin@v1.7.4/gin.go:489 +0x63e
github.com/gin-gonic/gin.(*Engine).ServeHTTP(0xc0003911e0, {0x163d310, 0xc0003b7260}, 0xc0001d2d00)
        /Users/jake/go/pkg/mod/github.com/gin-gonic/gin@v1.7.4/gin.go:445 +0x1c5
net/http.serverHandler.ServeHTTP({0xc000430570}, {0x163d310, 0xc0003b7260}, 0xc0001d2d00)
        /usr/local/go/src/net/http/server.go:2878 +0x43b
net/http.(*conn).serve(0xc0003edc20, {0x1640bc0, 0xc000430480})
        /usr/local/go/src/net/http/server.go:1929 +0xb08
created by net/http.(*Server).Serve
        /usr/local/go/src/net/http/server.go:3033 +0x4e8
```

## FAQ

- Is this idiomatic Go?

No, it's idiomatic to pass errors all the way back up to the handler where the error is written.

- Is there a downside to doing it this way for certain unexpected errors?

Other than being shunned by the Go community? I don't think so. If there is a reason let me know in an issue.

