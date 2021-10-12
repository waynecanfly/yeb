package main

import (
	_ "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"yeb/controllers"
)


func CollectRouter(r *gin.Engine) *gin.Engine {
	//r.Use(middleware.CORSMiddleware())
	r.GET("v1/captcha", func(c *gin.Context) {
		controllers.Captcha(c, 4)
	})
	r.POST("/v1/register", controllers.Register)
	//r.POST("/v1/login", controller.Login)
	return r
}
