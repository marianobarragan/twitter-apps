package routes

import (
	"github.com/gin-gonic/gin"
	"tweets-service/controllers"
	"tweets-service/domain"
)

func NewRouter(s domain.Service) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controllers.Ping)
	r.GET("/tweets/:id", controllers.GetTweet(s))
	r.POST("/tweets", controllers.PostTweet(s))
	return r
}
