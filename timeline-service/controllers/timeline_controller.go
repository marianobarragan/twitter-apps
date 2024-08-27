package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"timeline-service/domain"
)

func Ping(c *gin.Context) {
	c.String(200, "pong")
}

func GetUserTimeline(s domain.Service) func(c *gin.Context) {
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

		from := c.Query("from")
		if from == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "must provide a valid from"})
			return
		}

		to := c.Query("to")
		if to == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "must provide a valid to"})
			return
		}

		fromSec, err := strconv.Atoi(from)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "must provide a valid to timestamp"})
			return
		}
		toSec, err := strconv.Atoi(to)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "must provide a valid from timestamp"})
			return
		}

		timeline, err := s.GetUserTimeline(id, time.Unix(int64(fromSec), 0), time.Unix(int64(toSec), 0))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, timeline)
	}
}

func AddUserTimeline(s domain.Service) func(c *gin.Context) {
	return func(c *gin.Context) {
		var tweet domain.Tweet
		if err := c.ShouldBindJSON(&tweet); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "unable to parse JSON body"})
			return
		}

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

		err = s.AddTweetToUserTimeline(id, tweet)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "OK"})
	}
}
