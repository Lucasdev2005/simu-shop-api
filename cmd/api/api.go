package main

import (
	"fmt"
	"net/http"
	"os"
	"simushop/internal/core"
	"simushop/internal/database"
	"simushop/internal/entity"
	"simushop/internal/handler"
	"simushop/internal/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
)

func main() {
	app := core.NewApp()

	database := database.NewDatabase(
		postgres.Open(fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)),
		entity.User{},
		entity.Topic{},
		entity.Product{},
		entity.ShoppingCart{},
	)

	repository := repository.NewRepository(database.Db)

	handler := handler.NewHandler(repository)

	app.Group("api", func(g *gin.RouterGroup) {
		userGroup := g.Group("user")
		{
			app.GroupMethod(userGroup, http.MethodPost, "/", handler.CreateUser)
			app.GroupMethod(userGroup, http.MethodPut, "/:id", handler.UpdateUser)
		}

		productGroup := g.Group("product")
		{
			app.GroupMethod(productGroup, http.MethodPost, "/", handler.CreateProduct)
			app.GroupMethod(productGroup, http.MethodPut, "/:id", handler.UpdateProduct)
			app.GroupMethod(productGroup, http.MethodGet, "/", handler.ListProducts)
		}

		topicGroup := g.Group("topic")
		{
			app.GroupMethod(topicGroup, http.MethodPost, "/", handler.CreateTopic)
			app.GroupMethod(topicGroup, http.MethodPut, "/:id", handler.UpdateTopic)
			app.GroupMethod(topicGroup, http.MethodGet, "/", handler.ListTopics)
		}

		shoppingCartGroup := g.Group("shopping-cart")
		{
			app.GroupMethod(shoppingCartGroup, http.MethodPost, "/", handler.AddItemOnCart)
		}
	})

	app.Run()
}
