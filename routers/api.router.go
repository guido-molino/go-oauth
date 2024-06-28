package routers

import (
	v1 "main/routers/v1"

	"github.com/gin-gonic/gin"
)

func ApiRouterGroup(r *gin.Engine) {
	// HealthCheck
	r.GET("/nanobastardo", func(c *gin.Context) {
		c.Header("Content-Type", "image/jpeg")
		c.File("./assets/nanobastardo.jpeg")
	})

	api := r.Group("/api")
	// Grouping endpoints under /api/v1 route for versioning
	v1.SetV1Routes(api)
}
