package core

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type app struct {
	router *gin.Engine
}

func NewApp() app {
	router := gin.Default()
	loadEnviroment()

	return app{
		router,
	}
}

func (a app) AplyMiddlewares(middlewares ...func(c *gin.Context)) {
	for _, middleware := range middlewares {
		a.router.Use(middleware)
	}
}

func (a app) Group(endPoint string, fn func(g *gin.RouterGroup)) {
	group := a.router.Group(endPoint)
	{
		fn(group)
	}
}

func (a app) Post(endpoint string, handler func(Request) (Success, Fail)) {
	a.router.POST(endpoint, func(ctx *gin.Context) {
		a.ProcessCtx(ctx, handler)
	})
}

func (a app) Get(endpoint string, handler func(Request) (Success, Fail)) {
	a.router.GET(endpoint, func(ctx *gin.Context) {
		a.ProcessCtx(ctx, handler)
	})
}

func (a app) Delete(endpoint string, handler func(Request) (Success, Fail)) {
	a.router.DELETE(endpoint, func(ctx *gin.Context) {
		a.ProcessCtx(ctx, handler)
	})
}

func (a app) Head(endpoint string, handler func(Request) (Success, Fail)) {
	a.router.HEAD(endpoint, func(ctx *gin.Context) {
		a.ProcessCtx(ctx, handler)
	})
}

func (a app) Patch(endpoint string, handler func(Request) (Success, Fail)) {
	a.router.PATCH(endpoint, func(ctx *gin.Context) {
		a.ProcessCtx(ctx, handler)
	})
}

func (a app) Options(endpoint string, handler func(Request) (Success, Fail)) {
	a.router.OPTIONS(endpoint, func(ctx *gin.Context) {
		a.ProcessCtx(ctx, handler)
	})
}

func (a app) Put(endpoint string, handler func(Request) (Success, Fail)) {
	a.router.PUT(endpoint, func(ctx *gin.Context) {
		a.ProcessCtx(ctx, handler)
	})
}

func (a app) ProcessCtx(ctx *gin.Context, handler func(Request) (Success, Fail)) {
	request := NewRequest(ctx)
	success, fail := handler(request)

	if fail.ErrorCode != 0 {
		ctx.JSON(fail.ErrorCode, fail.Message)
	} else {
		ctx.JSON(success.SuccessCode, success.Data)
	}
}

func (a app) GroupMethod(group *gin.RouterGroup, method string, endPoint string, handler func(Request) (Success, Fail)) {
	switch method {
	case "GET":
		group.GET(endPoint, func(ctx *gin.Context) {
			a.ProcessCtx(ctx, handler)
		})
	case "POST":
		group.POST(endPoint, func(ctx *gin.Context) {
			a.ProcessCtx(ctx, handler)
		})
	case "PUT":
		group.PUT(endPoint, func(ctx *gin.Context) {
			a.ProcessCtx(ctx, handler)
		})
	case "DELETE":
		group.DELETE(endPoint, func(ctx *gin.Context) {
			a.ProcessCtx(ctx, handler)
		})
	default:
		panic("Método HTTP não suportado")
	}
}

func (a app) Run() {
	a.router.Run(":3333")
}

func loadEnviroment() {
	_, b, _, _ := runtime.Caller(0)
	var ProjectRootPath = filepath.Join(filepath.Dir(b), "../../")
	godotenv.Load(ProjectRootPath + "/.env")
}

func getUrlDatabase() string {
	dbPort := os.Getenv("DB_PORT")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)

}
