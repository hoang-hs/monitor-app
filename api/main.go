package main

import (
	"api/src/present/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	middleware.RecordMetrics()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/prometheus", gin.WrapH(promhttp.Handler()))
	if err := r.Run(":1234"); err != nil {
		panic(err)
	}
}
