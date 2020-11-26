package handler

import (
	"github.com/gin-gonic/gin"
)

func (rh *Handler) Get(c *gin.Context) {
	result := rh.uc.Get()
		c.JSON(200, gin.H{
			"data": result,
		})
}