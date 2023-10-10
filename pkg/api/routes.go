package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	ginEngine := gin.Default()

	handler := Handler{DB: db}

	ginEngine.GET("/hello", handler.Hello)
	ginEngine.GET("/os", handler.GetOs)

	ginEngine.GET("/comments", handler.CommentList)
	ginEngine.GET("/comments/:id", handler.CommentGet)
	ginEngine.POST("/comments", handler.CommentCreate)
	ginEngine.PUT("/comments/:id", handler.CommentUpdate)
	ginEngine.DELETE("/comments/:id", handler.CommentDelete)

	return ginEngine
}
