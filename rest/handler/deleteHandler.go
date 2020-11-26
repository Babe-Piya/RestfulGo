package handler

import (
	"strconv"
	"github.com/gin-gonic/gin"
)

func (rh *Handler) Del(c *gin.Context){
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{
				"data": err,
			})
			return
		}
		result, err := rh.uc.Del(id)
		if err != nil {
			c.JSON(400, gin.H{
				"Error": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"data": result,
		})
}