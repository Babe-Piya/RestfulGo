package handler

import (
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
)

func (rh *Handler) Put(c *gin.Context) {
	var body map[string]interface{}
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(400, gin.H{
				"data": err,
			})
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			log.Print(err)
			c.JSON(400, gin.H{
				"data": err.Error(),
			})
			return
		}
		result,err := rh.uc.Update(id,body)
		if err != nil {
			log.Println(err)
			c.JSON(400, gin.H{
				"data": err.Error(),
			})
			return
		}
		c.JSON(202, gin.H{
			"Result": result,
		})
}