package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "net/http"
    "golang.org/x/sync/errgroup"
    "time"
)
var (
    g errgroup.Group
)
func router01() http.Handler {
   e := gin.New()
   e.Use(gin.Recovery())
   e.GET("/", func(c *gin.Context) {
       c.JSON(http.StatusOK,gin.H{
           "code":http.StatusOK,
           "msg":"server01",
       })
   })
   return e
}
func router02() http.Handler {
    e := gin.New()
    e.Use(gin.Recovery())
    e.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK,gin.H{
            "code":http.StatusOK,
            "msg":"server02",
        })
    })
    return e
}

func main() {
    // 第一个服务
    server01 := &http.Server{
        Addr: ":8080",
        Handler: router01(),
        ReadTimeout: 5*time.Second,
        WriteTimeout: 10*time.Second,
    }
    server02 := &http.Server{
        Addr: ":8081",
        Handler: router02(),
        ReadTimeout: 5*time.Second,
        WriteTimeout: 10*time.Second,
    }
    // 启动协程1
    g.Go(func() error {
        return server01.ListenAndServe()
    })
    g.Go(func() error {
        return server02.ListenAndServe()
    })
    
    
    
    // 等待
    if err := g.Wait();err != nil{
        log.Fatal(err)
    }
}
