package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func awaitDone(ctx context.Context, anotherChan <-chan any) {
	select {
	case <-anotherChan:
		fmt.Println("done")
	case <-ctx.Done():
		fmt.Println("canceled")
	}
	fmt.Println(ctx.Err()) // TOGGEL_B
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})

		ctx := c // TOGGLE_A
		// ctx := c.Request.Context() // TOGGLE_A

		anotherChan := make(chan any)
		go awaitDone(ctx, anotherChan)
		anotherChan <- nil
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run()
}
