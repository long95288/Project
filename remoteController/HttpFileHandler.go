package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "log"
)

func HttpFileHandler(context *gin.Context) {
    
    fmt.Println("文件上传请求....")
    multi, err := context.MultipartForm()
    if err != nil {
        fmt.Println(err)
        return
    }
    
    files := multi.File["file"]
    for index,file := range files{
        log.Println(index," ",file.Filename)
        err =  context.SaveUploadedFile(file,"G:\\win下载\\"+file.Filename)
        if err != nil {
            fmt.Println(err)
        }
        log.Printf("上传文件:%s, 完成", file.Filename)
    }
    context.JSON(200, gin.H{
        "resp":"",
    })
}