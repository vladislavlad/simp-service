package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"runtime"
	"simp-service/pkg/model"
	"strconv"
	"sync"
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
	var wg sync.WaitGroup
	var comments []model.Comment

	wg.Add(1)
	go dbFind(&wg, h, &comments)()

	wg.Wait()
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

func (h Handler) UpdateCreate(c *gin.Context) {
	var comment model.Comment

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot parse 'id' param"})
		return
	}

	h.DB.First(&comment, id)
	var commentUpdate CommentUpdate

	err = c.BindJSON(&commentUpdate)
	if err != nil {
		// add common error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot process JSON object"})
		return
	}

	comment.Comment = commentUpdate.Comment
	h.DB.Save(&comment)

	c.JSON(http.StatusOK, comment)
}

func dbFind(wg *sync.WaitGroup, h Handler, comments *[]model.Comment) func() {
	return func() {
		defer wg.Done()
		h.DB.Find(&comments)
	}
}
