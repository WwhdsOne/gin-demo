package initialize

import (
	"gin-demo/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	router := r.Group("/api")

	{
		privateRouter := router.Group("/system")

		controller.InitStudentRouter(privateRouter)

	}

}
