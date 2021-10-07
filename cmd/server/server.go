package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jakecoffman/stacktraces/service1"
	"log"
	"runtime/debug"
)

func main() {
	e := gin.New()

	e.Use(func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Println("Recovered from panic:", string(debug.Stack()))
				// Try to log a response, if it doesn't work because the connection
				// is closed then that's ok.
				context.JSON(500, "Internal Server Error")
			}
		}()
		context.Next()
	})

	e.GET("/", func(context *gin.Context) {
		service1.DoStuff1()
		context.JSON(200, "OK!")
	})

	if err := e.Run("127.0.0.1:8080"); err != nil {
		log.Fatalln(err)
	}
}
