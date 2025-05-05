package router

import (
	"github.com/clapat-bb/memo/controller"
	"github.com/clapat-bb/memo/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)

	auth := r.Group("/user")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/profile", controller.Profile)

		auth.POST("/memos", controller.CreateMemo)
		auth.GET("/memos", controller.ListMemos)
		auth.PUT("/memos/:id", controller.UpdateMemo)
		auth.DELETE("/memos/:id", controller.DeleteMemo)
	}
}
