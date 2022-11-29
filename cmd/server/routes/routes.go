package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vincentconace/app-fibonacci/cmd/server/handler"
	"github.com/vincentconace/app-fibonacci/internal/fibonacci"
)

type Router interface {
	MapRoutes()
}

type router struct {
	eng *gin.Engine
	rg  *gin.RouterGroup
}

func NewRouter(eng *gin.Engine) Router {
	return &router{eng: eng}
}

func (r *router) MapRoutes() {
	r.setGroup()

	r.buildFibonacciRoutes()
}

func (r *router) setGroup() {
	r.rg = r.eng.Group("/api/v1")
}

func (r *router) buildFibonacciRoutes() {

	service := fibonacci.NewService()
	handler := handler.NewHandler(service)
	fiboGroup := r.rg.Group("/fibonacci")
	{
		fiboGroup.POST("/", handler.Fibonacci())
	}
}
