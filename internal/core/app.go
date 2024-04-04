package core

import (
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

func (a app) ProcessCtx(ctx *gin.Context, handler func(Request) Response) {
	request := NewRequest(ctx)
	result := handler(request)

	ctx.JSON(result.Code, result.Data)
}

func (a app) GroupMethod(group *gin.RouterGroup, method string, endPoint string, handler func(Request) Response) {
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
