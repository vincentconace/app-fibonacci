package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vincentconace/app-fibonacci/cmd/server/routes"
)

func main() {
	r := gin.Default()
	router := routes.NewRouter(r)
	router.MapRoutes()

	r.Run()
}
