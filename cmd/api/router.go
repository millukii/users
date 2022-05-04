package main

import (
	"github.com/gin-gonic/gin"

)

// middleware setup
func setupRouter(r *gin.Engine, handler UserHandler) *gin.Engine {

	r.GET("/users", func(c *gin.Context) {
		handler.GetAll(c)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		handler.Get(c)
	})

	r.POST("/users", func(c *gin.Context) {
		handler.Create(c)
	})

	return r
}
