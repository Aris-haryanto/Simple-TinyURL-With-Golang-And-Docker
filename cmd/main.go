package main

import (
	"tinyurl/api"
	"tinyurl/services"

	"github.com/gin-gonic/gin"
)

var (
	param api.TinyStore
	ts    services.TinyService
)

func main() {
	r := gin.Default()
	r.POST("/shorten", func(c *gin.Context) {
		c.BindJSON(&param)
		result := ts.Store(param)
		c.JSON(result.HttpCode, result)
	})
	r.GET("/:shortcode", func(c *gin.Context) {
		result := ts.Get(c.Param("shortcode"))
		c.JSON(result.HttpCode, result)
	})
	r.GET("/:shortcode/stats", func(c *gin.Context) {
		result := ts.Stats(c.Param("shortcode"))
		c.JSON(result.HttpCode, result)
	})

	r.Run(":3000")
}
