package rest

import (
	"restfulGo/pkg/database"
	"restfulGo/repository"
	"restfulGo/usecase"
	"restfulGo/rest/handler"
	"github.com/gin-gonic/gin"
)

func LetGo () {
	database.InitDB()
	repo := repository.CreateRepo(database.DB)
	uc := usecase.CreateUsecase(repo)
	rest := handler.CreateHandler(uc)
	runGin := gin.Default()

	runGin.POST("/",rest.Add )
	runGin.GET("/",rest.Get )
	runGin.DELETE("/:id",rest.Del )
	runGin.PUT("/:id",rest.Put )

	runGin.Run(":8081")
}