package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
)

func main() {
    r := gin.Default()
    // 1、获得URL参数
    r.GET("/p1", func(c *gin.Context) {
        param1 := c.DefaultQuery("param1","no exit")
        param2 := c.Query("param2")
        c.JSON(http.StatusOK,gin.H{
            "param1":param1,
            "param2":param2,
        })
    })
    // 请求:http://localhost:8080/p1?param1=11&param2=22
    // 返回:{"param1":"11","param2":"22"}
    // 请求:http://localhost:8080/p1?param2=22
    // 返回:{"param1":"no exit","param2":"22"}
    
    // 2.form参数
    r.POST("/p2", func(c *gin.Context) {
        param1 := c.PostForm("param1")
        param2 := c.PostForm("param2")
        c.JSON(http.StatusOK,gin.H{
            "param1":param1,
            "param2":param2,
        })
    })
    
    // 3.获得path参数
    r.GET("/p3/:param1/:param2", func(c *gin.Context) {
        param1 := c.Param("param1")
        param2 := c.Param("param2")
        c.JSON(http.StatusOK,gin.H{
            "param1":param1,
            "param2":param2,
        })
    })
    // 请求:http://localhost:8080/p3/pp1/pp2
    // 返回:{"param1":"pp1","param2":"pp2"}
    //
    
    // shouldBind()会根据请求的content-type自动选择绑定器
    
    // 4.获得json数据
    type s1 struct {
        Name string `json:"name" form:"name" binding:"required"`
        Age int `json:"age" form:"age" binding:"required"`
    }
    r.POST("/p4", func(c *gin.Context) {
        var s s1
        if err := c.ShouldBind(&s);err == nil{
            c.JSON(http.StatusOK,s)
        }else{
            c.JSON(http.StatusBadRequest,gin.H{
                "err":err.Error(),
            })
        }
    })
    
    // 5.绑定form表单数据
    r.POST("/p5", func(c *gin.Context) {
        var s s1
        if err := c.ShouldBind(&s);err == nil{
            c.JSON(http.StatusOK,s)
        }else{
            c.JSON(http.StatusBadRequest,gin.H{
                "err":err.Error(),
            })
        }
    })
    // 6.url路径绑定
    // 请求:http://localhost:8080/p6?name=fff&age=12
    // 响应:{"name":"fff","age":12}
    r.GET("/p6", func(c *gin.Context) {
        var s s1
        if err := c.ShouldBind(&s);err == nil{
            c.JSON(http.StatusOK,s)
        }else{
            c.JSON(http.StatusBadRequest,gin.H{
                "err":err.Error(),
            })
        }
        
    })
    // 7.单文件上传
    r.POST("/upload", func(c *gin.Context) {
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
    // 8.多文件上传
    //r.MaxMultipartMemory =
    r.POST("/uploads", func(c *gin.Context) {
        form ,err := c.MultipartForm()
        if err != nil {
            c.JSON(http.StatusInternalServerError,gin.H{
                "err":err.Error(),
            })
            return
        }
        files := form.File["file"]
        for index,file := range files{
            log.Println(index," ",file.Filename)
            c.SaveUploadedFile(file,file.Filename)
        }
        c.JSON(http.StatusOK,gin.H{
            "message":fmt.Sprintf("上传%d个文件成功",len(files)),
        })
    })
    r.Run(":8080")
}
