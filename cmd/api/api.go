package main

import (
	"net/http"
	"os"
	"simushop/internal/core"
	"simushop/internal/database"
	"simushop/internal/entity"
	"simushop/internal/handler"
	"simushop/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	app := core.NewApp()

	database := database.NewDatabase(
		os.Getenv("DB_PORT"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		entity.User{},
	)

	repository := repository.NewRepository(database.Db)

	handler := handler.NewHandler(repository)

	app.Group("api", func(g *gin.RouterGroup) {
		userGroup := g.Group("user")
		{
			app.GroupMethod(userGroup, http.MethodPost, "/", handler.CreateUser)
		}
	})

	app.Run()
}
