package users

import (
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Create(c *gin.Context) error
	Get(c *gin.Context) (error)
	GetAll(c *gin.Context)  (error)
}
