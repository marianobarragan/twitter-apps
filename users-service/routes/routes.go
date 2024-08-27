package routes

import (
	"github.com/gin-gonic/gin"
	"users-service/controllers"
	"users-service/domain"
)

func NewRouter(s domain.Service) *gin.Engine {
	r := gin.Default()
	r.GET("/ping", controllers.Ping)
	r.GET("/users/:id", controllers.GetUser(s))
	r.POST("/users", controllers.PostUser(s))
	return r
}
