package routers

import "github.com/gin-gonic/gin"

func InitRouters() *gin.Engine {
	r := gin.New()
	//限制单次最大文件上传内存
	r.MaxMultipartMemory = 8 << 20
	apiGroup := r.Group("/api")
	InitFileUploadRouter(apiGroup)
	return r
}
