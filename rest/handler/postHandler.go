package handler

import (
	"time"
	"restfulGo/entity"
	"github.com/gin-gonic/gin"
	"log"
)

func (rh *Handler) Add(c *gin.Context) {
	var body Req
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"data": err.Error(),
		})
		return
	}
	result, _ := rh.uc.Add(&entity.Todo{
		Title:    body.Title,
		Content:  body.Content,
		IsDone:   false,
		CreateAt: time.Now(),
	})
	c.JSON(201, gin.H{
		"Result": result,
	})
	log.Print(body)
}