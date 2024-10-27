package initialize

import (
	"fmt"
	"gin-demo/controller"
	"gin-demo/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	router := r.Group("/api")
	{
		privateRouter := router.Group("/system")
		privateRouter.Use(middleware.JWTAuth())
		controller.InitStudentRouter(privateRouter)
	}
	{
		publicRouter := router.Group("/")
		controller.InitUserRouter(publicRouter)
	}
	rs := r.Routes()
	for _, v := range rs {
		fmt.Println(v.Path, v.Method)
	}

}
