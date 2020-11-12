package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/Asuha-a/URLShortener/api/controllers"
)

func userRouters(router *gin.RouterGroup) {
	u := router.Group("/users")
	u.GET("login", controllers.Test)
	u.POST("login", controllers.Login)
	u.POST("signup", controllers.Signup)
}
