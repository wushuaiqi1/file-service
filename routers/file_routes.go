package routers

import (
	"file-service/controller"
	"github.com/gin-gonic/gin"
)

//初始化文件路由

func InitFileUploadRouter(r *gin.RouterGroup) gin.IRouter {
	handler := controller.NewFileUploadController()
	router := r.Group("/file")
	router.POST("/upload", handler.FileUpload)
	return r
}
