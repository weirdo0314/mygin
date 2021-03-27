package main

import (
	"log"
	"mygin/gin"
	"net/http"
	"time"
)

func onlyForV2() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred

		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatueCode, c.R.RequestURI, time.Since(t))
	}
}
func main() {
	r := gin.New()
	r.Use(gin.Logger()) // global midlleware
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *gin.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Pattern)
		})
	}
	r.Run(":8080")
}
