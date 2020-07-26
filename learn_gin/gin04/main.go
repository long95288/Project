package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "time"
)
//  计时中间件
func StatCost() gin.HandlerFunc{
    return func(c *gin.Context) {
        start := time.Now()
        log.Println("开始计时")
        c.Set("name","transName")
        c.Next()
        cost := time.Since(start)
        log.Printf("请求花费时间:%v\n",cost)
    }
}

func main() {
    r := gin.Default()
    // 7.单文件上传
    r.POST("/upload",StatCost(), func(c *gin.Context) {
        file,err := c.FormFile("file1")
        if err != nil{
            c.JSON(http.StatusInternalServerError, gin.H{
                "err":err.Error(),
            })
            return
        }
        log.Println("上传文件名:",file.Filename)
        //  保存文件
        err = c.SaveUploadedFile(file,file.Filename)
        if err != nil {
            log.Println(err)
            return
        }else{
            c.JSON(http.StatusOK,gin.H{
                "message":fmt.Sprintf("%s 上传成功",file.Filename),
            })
        }
    })
    r.Run(":8080")
}
