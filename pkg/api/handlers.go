package api

import (
	"github.com/gin-gonic/gin"
	"github.com/vladislavlad/goroutines"
	"gorm.io/gorm"
	"net/http"
	"runtime"
	"simp-service/pkg/model"
	"strconv"
	"time"
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

	goroutines.Launch(
		func() { h.DB.Find(&comments) },
	)

	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

func (h Handler) CommentGet(c *gin.Context) {
	var comment model.Comment

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.MakeError("Cannot parse 'id' param"))
		return
	}

	goroutines.Launch(
		func() { h.DB.First(&comment, id) },
	)

	c.JSON(http.StatusOK, comment)
}

func (h Handler) CommentCreate(c *gin.Context) {
	var comment model.Comment

	err := c.ShouldBindJSON(&comment)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.MakeError("Cannot process JSON object"))
		return
	}

	now := time.Now()
	comment.CreatedAt = now
	comment.UpdatedAt = now
	comment.DeletedAt = nil
	h.DB.Save(&comment)
	c.Status(http.StatusCreated)
}

func (h Handler) CommentUpdate(c *gin.Context) {
	var comment model.Comment

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.MakeError("Cannot parse 'id' param"))
		return
	}

	goroutines.Launch(
		func() { h.DB.First(&comment, id) },
	)

	var commentUpdate CommentUpdate
	err = c.ShouldBindJSON(&commentUpdate)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.MakeError("Cannot process JSON object"))
		return
	}

	comment.Text = commentUpdate.Text
	comment.UpdatedAt = time.Now()
	h.DB.Save(&comment)

	c.JSON(http.StatusOK, comment)
}

func (h Handler) CommentDelete(c *gin.Context) {
	var comment model.Comment

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.MakeError("Cannot parse 'id' param"))
		return
	}

	goroutines.Launch(
		func() { h.DB.First(&comment, id) },
	)

	now := time.Now()
	comment.DeletedAt = &now
	h.DB.Save(&comment)

	c.Status(http.StatusNoContent)
}
