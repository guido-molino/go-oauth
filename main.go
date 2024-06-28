package main

import (
	"main/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Load App configs
	routers.ApiRouterGroup(r)
	r.Run()
}
