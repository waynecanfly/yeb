package router

import (
	_ "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"yeb/controllers"
	"yeb/middlewares"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middlewares.Cors())

	v1 := r.Group("v1")
	{
		v1.GET("/captcha", func(c *gin.Context) {
			controllers.Captcha(c, 4)
		})
	}

	err := r.Run(":8081")
	if err != nil {
		return 
	}
}