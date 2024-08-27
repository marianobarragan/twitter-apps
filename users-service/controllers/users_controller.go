package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"users-service/domain"
)

func Ping(c *gin.Context) {
	c.String(200, "pong")
}

func GetUser(s domain.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		if idParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "must provide a user id"})
			return
		}

		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "must provide a valid user id"})
			return
		}

		user, found, err := s.GetUser(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if !found {
			errMsg := fmt.Sprintf("user %d not found", id)
			c.JSON(http.StatusNotFound, gin.H{"error": errMsg})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func PostUser(s domain.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		var user domain.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "unable to parse JSON body"})
			return
		}

		user, err := s.SaveUser(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}

func AddUserSubscription(s domain.Service) func(c *gin.Context) {
	return func(c *gin.Context) {

		// TODO implement method

		err := errors.New("unimplemented method")
		c.JSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
		return
	}
}

func RemoveUserSubscription(s domain.Service) func(c *gin.Context) {
	return func(c *gin.Context) {

		// TODO implement method

		err := errors.New("unimplemented method")
		c.JSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
		return
	}
}
