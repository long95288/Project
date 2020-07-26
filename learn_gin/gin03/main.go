package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)
func main() {
    r := gin.Default()
    r.Any("/test", func(c *gin.Context) {
        c.JSON(http.StatusOK,gin.H{
            "msg":"success",
        })
    })
    r.NoRoute(func(c *gin.Context) {
        c.JSON(http.StatusNotFound,gin.H{
            "msg":"404 not found",
        })
    })
    r.Run(":8080")
    
}
