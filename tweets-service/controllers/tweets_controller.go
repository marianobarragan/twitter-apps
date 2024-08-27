package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tweets-service/domain"
)

func Ping(c *gin.Context) {
	c.String(200, "pong")
}

func GetTweet(s domain.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		if idParam == "" {
			err := errors.New("must provide a tweet id")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tweet, found, err := s.GetTweet(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if !found {
			errMsg := fmt.Sprintf("user %d not found", id)
			c.JSON(http.StatusNotFound, gin.H{"error": errMsg})
			return
		}

		c.JSON(http.StatusOK, tweet)
	}
}

func PostTweet(s domain.Service) func(c *gin.Context) {
	return func(c *gin.Context) {

		userToken := c.GetHeader("x-auth-token")
		if userToken == "" {
			err := errors.New("you must provide a auth token in your request ('x-auth-token')")
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// TODO implement auth service, get userID from token
		userID, err := strconv.Atoi(userToken)
		if err != nil {
			err2 := errors.New("unable to parse auth token ('x-auth-token')")
			c.JSON(http.StatusUnauthorized, gin.H{"error": err2.Error()})
			return
		}

		var tweet domain.Tweet
		if err := c.ShouldBindJSON(&tweet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(tweet.Text) > 280 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "tweet text exceeds 280 characters"})
			return
		}

		tweet, err = s.SaveTweet(userID, tweet)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, tweet)
	}
}
