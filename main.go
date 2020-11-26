package main

import (
	"log"
	"restfulGo/entity"
	"restfulGo/repository"
	"restfulGo/usecase"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type req struct {
	Content string `json:"content"`
	Title   string `json:"title" ShouldBindJSONnding:"required"`
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
		result, _ := uc.Add(&entity.Todo{
			Title:    body.Title,
			Content:  body.Content,
			IsDone:   true,
			CreateAt: time.Now(),
		})
		c.JSON(201, gin.H{
			"Result": result,
		})
		log.Print(body)
	})

	r.GET("/", func(c *gin.Context) {
		result := uc.Get()
		c.JSON(200, gin.H{
			"data": result,
		})
	})

	r.DELETE("/:id", func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{
				"data": err,
			})
			return
		}
		result, err := uc.Del(id)
		if err != nil {
			c.JSON(400, gin.H{
				"Error": err,
			})
			return
		}
		c.JSON(200, gin.H{
			"data": result,
		})
	})

	r.PUT("/:id",func (c *gin.Context){
		var body req
		if err := c.ShouldBindJSON(&body); err != nil {
			return
		}
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(400, gin.H{
				"data": err,
			})
			return
		}
		result,err := uc.Update(id,&entity.Todo{
			Title:    body.Title,
			Content:  body.Content,
			IsDone:   true,
			CreateAt: time.Now(),
		})
		if err != nil {
			c.JSON(400, gin.H{
				"data": err,
			})
			return
		}
		c.JSON(202, gin.H{
			"Result": result,
		})
	})

	r.Run()
}
