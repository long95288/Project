package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

func helloHandler(c *gin.Context)  {
    c.JSON(http.StatusOK,gin.H{
        "message":"Hello world",
    })
    
}

func main() {
    r := gin.Default()
    r.GET("/hello",helloHandler)
    if err := r.Run(":8080"); err!= nil{
        fmt.Println("startup service failed,err:%v\n",err)
    }
}
