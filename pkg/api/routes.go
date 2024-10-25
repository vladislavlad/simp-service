package api

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *Handler) *gin.Engine {
	ginEngine := gin.Default()

	ginEngine.GET("/hello", handler.Hello)
	ginEngine.GET("/os", handler.GetOs)

	ginEngine.GET("/comments", handler.CommentList)
	ginEngine.GET("/comments/:id", handler.CommentGet)
	ginEngine.POST("/comments", handler.CommentCreate)
	ginEngine.PUT("/comments/:id", handler.CommentUpdate)
	ginEngine.DELETE("/comments/:id", handler.CommentDelete)

	return ginEngine
}
