package rest

import (
	"restfulGo/repository"
	"restfulGo/usecase"
	"restfulGo/rest/handler"
	"github.com/gin-gonic/gin"
)

func LetGo () {
	repo := repository.CreateRepo()
	uc := usecase.CreateUsecase(repo)
	rest := handler.CreateHandler(uc)
	runGin := gin.Default()

	runGin.POST("/",rest.Add )
	runGin.GET("/",rest.Get )
	runGin.DELETE("/:id",rest.Del )
	runGin.PUT("/:id",rest.Put )

	runGin.Run()
}