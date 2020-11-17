package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/Asuha-a/URLShortener/api/controllers"
)

func urlRouters(router *gin.RouterGroup) {
	u := router.Group("/urls")
	u.GET("", controllers.GetAllURL)
	// u.GET(":uuid", controllers.GetURL)
	u.POST("", controllers.PostURL)
	// u.PUT(":uuid", controllers.PutURL)
	// u.DELETE(":uuid", controllers.DeleteURL)
}
