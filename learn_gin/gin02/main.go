package main

import (
    "github.com/gin-gonic/gin"
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
        Name string `json:"name" form:"name"`
        Age int `json:"age" form:"age"`
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
    
    r.Run(":8080")
}
