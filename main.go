package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	r := gin.Default()

	c := cors.New(cors.Options{
		Debug:            true,
		AllowedHeaders:   []string{"Origin", "Authorization", "Content-Type"},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	})
	r.Use(func(ctx *gin.Context) {
		c.HandlerFunc(ctx.Writer, ctx.Request)
	})

	r.GET("/*any", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":6000")
}
