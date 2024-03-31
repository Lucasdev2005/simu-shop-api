package main

import (
	"net/http"
	"simushop/internal/core"
	"simushop/internal/handler"
	"simushop/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	app := core.NewApp()
	repository := repository.NewRepository(app.Db)
	handler := handler.NewHandler(repository)

	app.Group("profile", func(g *gin.RouterGroup) {
		app.GroupMethod(g, http.MethodPost, "/", handler.CreateUser)
	})

	app.Run()
}
