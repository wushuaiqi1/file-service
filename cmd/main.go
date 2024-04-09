package main

import (
	"file-service/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

func main() {
	
}

func test() {
	fmt.Println("main start....")

	//hash, err := utils.FileUpload()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(hash)

	r := gin.Default()
	//限制数据最大内存
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(context *gin.Context) {
		file, err := context.FormFile("file")
		if err != nil {
			fmt.Println("fail", err)
			return
		}
		log.Println(file.Filename)
		fmt.Println(file.Size) //423395b->423kb->0.4mb

		open, err := file.Open()
		if err != nil {
			return
		}
		all, err := io.ReadAll(open)
		upload, err := utils.FileUpload(all)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(upload)
		fmt.Println(file.Header["Content-Type"])
	})
	err := r.Run()
	if err != nil {
		return
	}
}
