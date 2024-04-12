package routers

import (
	"file-service/controller"
	"github.com/gin-gonic/gin"
)

//初始化文件路由

func InitFileUploadRouter(r *gin.RouterGroup) gin.IRouter {
	handler := controller.NewFileUploadController()
	router := r.Group("/file")
	//error on parse multipart form array: multipart: NextPart: EOF
	//  /file/upload/lock和/api/file/upload/lock/有区别 后者虽然调用到指定接口但数据传输失败
	router.POST("/upload", handler.FileUpload)
	router.POST("/upload/lock", handler.FileUploadLock)
	return r
}
