package main

import (
	"api-users/internal/users"
	"api-users/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Create(c *gin.Context) error
	Get(c *gin.Context) (error)
	GetAll(c *gin.Context)  (error)
}

type user struct {
		svc users.UserService
}
func NewUserHandler(svc users.UserService) UserHandler {
	return &user{svc: svc}
}

func (u *user) Create(c *gin.Context) error {
validUser := &users.User{}	
 if err := c.ShouldBindJSON(&validUser); err != nil {
	c.AbortWithError(http.StatusBadRequest, err)
	return nil
 }

	err := u.svc.Save(c, validUser)
	if err != nil {
			respondError(c, err)
	}
	c.JSON(http.StatusOK, "User created")
	return nil
}
func (u *user) Get(c *gin.Context) (error) {

	id, _ :=	c.Params.Get("id")
	user, err := u.svc.Get(c, id)
	if err != nil {
		respondError(c, err)
	}
	c.JSON(http.StatusOK, user)
	return nil
}
func (u *user) GetAll(c *gin.Context) (error) {
	users, err := u.svc.GetAll(c)
	if err != nil {
		respondError(c, err)
	}
	c.JSON(http.StatusOK, users)
	return nil
}

func respondError(c *gin.Context,err error)  error {
	switch err {
		case errors.ErrResourceNotFound:
			c.JSON(http.StatusNotFound, err)
			return err
		case errors.ErrInvalidPayload:
				c.AbortWithError(http.StatusBadRequest, err)
				return err
		case errors.ErrUnauthenticated:
				c.AbortWithError(http.StatusForbidden, err)
				return err
		case errors.ErrUnauthorized:
				c.AbortWithError(http.StatusForbidden, err)
				return err
		default:
				c.AbortWithError(http.StatusInternalServerError, err)
	}
	return nil
}	