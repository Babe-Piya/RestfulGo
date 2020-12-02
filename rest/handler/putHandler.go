package handler

import (
	"time"
	"restfulGo/entity"
	"log"
	"strconv"
	"github.com/gin-gonic/gin"
)

func (rh *Handler) Put(c *gin.Context) {
	var body ReqPut
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
		result,err := rh.uc.Update(id,&entity.Todo{
			Title:    body.Title,
			Content:  body.Content,
			IsDone:   body.IsDone,
			CreateAt: time.Now(),
		})
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