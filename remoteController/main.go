package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "os/exec"
)

func main() {
    e:= gin.Default()
    e.GET("/cmd",HandleController)
    e.GET("/", func(context *gin.Context) {
        context.JSON(200, gin.H{
            "command_list":[]string{
                "http://localhost:9999/cmd?id=1",
                "http://localhost:9999/cmd?id=2",
            },
        })
    })
    err := e.Run(":9999")
    if err != nil {
        log.Fatal(err)
    }
    log.Println("service start at http://localhost:9999")
}

func shutdown(c *gin.Context){
    args :=[]string{"-s","-t","30"}
    cmd := exec.Command("shutdown",args...)
    err := cmd.Run()
    if err != nil {
        c.JSON(http.StatusInternalServerError,gin.H{
            "message":fmt.Sprintf("关机失败,err : %v",err),
        })
    }else{
        c.JSON(http.StatusOK,gin.H{
            "message": "电脑将在30s后关机",
        })
    }
}
func cancelShutdown(c *gin.Context){
    args := []string{"-a"}
    cmd := exec.Command("shutdown",args...)
    err := cmd.Run()
    if err != nil {
        c.JSON(http.StatusInternalServerError,gin.H{
            "message":fmt.Sprintf("取消关机失败,err : %v",err),
        })
    }else{
        c.JSON(http.StatusOK,gin.H{
            "message":"取消关机成功",
        })
    }
}

func HandleController(c *gin.Context)  {
    id := c.Query("id")
    switch id {
    case "1":
        shutdown(c)
    case "2":
        cancelShutdown(c)
    default:
        c.JSON(
            http.StatusBadRequest,
            gin.H{
                "message":"bad Request",
        })
    }
}
