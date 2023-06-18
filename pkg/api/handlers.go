package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"runtime"
	"simp-service/pkg/model"
)

type Handler struct {
	DB *gorm.DB
}

func (h Handler) GetOs(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"os": runtime.GOOS,
		},
	)
}

func (h Handler) Hello(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Hello World!",
		},
	)
}

func (h Handler) CommentList(c *gin.Context) {
	var comments []model.Comment
	h.DB.Find(&comments)

	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

func (h Handler) CommentCreate(c *gin.Context) {
	var comment model.Comment

	err := c.BindJSON(&comment)
	if err != nil {
		// add common error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot process JSON object"})
		return
	}

	h.DB.Save(&comment)
	c.Status(http.StatusCreated)
}
