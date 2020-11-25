package main

import (
	"log"
	"time"
	"restfulGo/entity"
	"github.com/gin-gonic/gin"
	"restfulGo/repository"
	"restfulGo/usecase"
)

type req struct {
	Content string `json:"content"`
	Title   string `json:"title" binding:"required"`
}

func main() {
	repo := repository.CreateRepo()
	uc := usecase.CreateUsecase(repo)
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		var body req
		if err := c.ShouldBindJSON(&body); err != nil {
			return
		}
			result,_ := uc.Add(&entity.Todo{
				Title : body.Title,
				Content : body.Content,
				IsDone : true,
				CreateAt : time.Time{},
			})
			c.JSON(200,gin.H{	
				"Result" : result,
			})
			log.Print(body)
	})

	r.GET("/", func(c *gin.Context) {
		result := uc.Get() 
		c.JSON(200,gin.H{
			"data" : result,
		})
	})
	r.Run()
}
