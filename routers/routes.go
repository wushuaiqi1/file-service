package routers

import "github.com/gin-gonic/gin"

func InitRouters() *gin.Engine {

	r := gin.New()
	apiGroup := r.Group("/api")
	InitFileUploadRouter(apiGroup)
	return r
}
