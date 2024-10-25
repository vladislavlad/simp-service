package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"runtime"
	"simp-service/pkg/model"
	"strconv"
	"sync"
	"time"
)

type Handler struct {
	DB *gorm.DB
}

func (h Handler) dbFindList(wg *sync.WaitGroup, comments *[]model.Comment) func() {
	return func() {
		defer wg.Done()
		h.DB.Find(&comments)
	}
}

func (h Handler) dbFindById(wg *sync.WaitGroup, comment *model.Comment, id uint64) func() {
	return func() {
		defer wg.Done()
		h.DB.First(&comment, id)
	}
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
	go h.dbFindList(&wg, &comments)()

	wg.Wait()
	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

func (h Handler) CommentGet(c *gin.Context) {
	var wg sync.WaitGroup
	var comment model.Comment

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.MakeError("Cannot parse 'id' param"))
		return
	}

	wg.Add(1)
	go h.dbFindById(&wg, &comment, id)()

	wg.Wait()
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
	var wg sync.WaitGroup
	var comment model.Comment

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.MakeError("Cannot parse 'id' param"))
		return
	}

	wg.Add(1)
	go h.dbFindById(&wg, &comment, id)()
	wg.Wait()

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
	var wg sync.WaitGroup
	var comment model.Comment

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.MakeError("Cannot parse 'id' param"))
		return
	}

	wg.Add(1)
	go h.dbFindById(&wg, &comment, id)()
	wg.Wait()

	now := time.Now()
	comment.DeletedAt = &now
	h.DB.Save(&comment)

	c.Status(http.StatusNoContent)
}
