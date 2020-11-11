package routers

import (
	"github.com/gin-gonic/gin"
)

// Init user routers
func Init() {
	r := gin.Default()
	api := r.Group("/api")
	v1 := api.Group("/v1")
	userRouters(v1)
	r.Run()
}
