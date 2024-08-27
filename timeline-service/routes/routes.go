package routes

import (
	"github.com/gin-gonic/gin"
	"timeline-service/controllers"
	"timeline-service/domain"
)

func NewRouter(s domain.Service) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controllers.Ping)
	r.GET("/users/:id/timeline", controllers.GetUserTimeline(s))
	r.POST("/users/:id/timeline", controllers.AddUserTimeline(s))
	return r
}
