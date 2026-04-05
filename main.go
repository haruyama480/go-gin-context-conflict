package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func delayedFunc(ctx context.Context) {
	select {
	case <-time.After(1 * time.Second):
		fmt.Println("1s passed without cancel")
		return

	case <-ctx.Done():
		fmt.Println("canceled")
		return
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})

		ctx := c // TOGGLE_A
		// ctx := c.Request.Context() // TOGGLE_A

		// go delayedFunc(ctx) // TOGGLE_B
		delayedFunc(ctx) // TOGGLE_B
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
