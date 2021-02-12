package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "html/template"
    "log"
    "net/http"
    "os/exec"
    "strconv"
    "time"
)

type status struct {
    Status string
}
var globalStatus = status{Status: "正常"}
const TIME_LAYOUT = "2006-01-02T15:04"

func main() {
    e:= gin.Default()
    e.GET("/cmd",HandleController)
    e.GET("/", func(context *gin.Context) {
        t,err := template.ParseFiles("index.html")
        if err != nil {
            log.Println("err : ", err)
            return
        }
        globalStatus.Status = "正常"
        log.Println(t.Execute(context.Writer, globalStatus))
    })
    
    e.POST("/SetShutdownTime", func(context *gin.Context) {
        t,err := template.ParseFiles("index.html")
        if err != nil {
            log.Println("err : ", err)
            return
        }
        r := context.Request
        err = r.ParseForm()
        if err != nil {
            t.Execute(context.Writer, status{Status: err.Error()})
            return
        }
        arrTime := r.Form["time"]
        if len(arrTime) < 1 {
            t.Execute(context.Writer, status{Status: "未设置关机时间"})
            return
        }
        shutdownTime := arrTime[0]
        if "" == shutdownTime {
            t.Execute(context.Writer, status{Status: "错误时间格式"})
            return
        }
        times , err := time.Parse(TIME_LAYOUT,shutdownTime)
        if nil != err {
            t.Execute(context.Writer, status{Status: "错误时间格式"})
            return
        }
        d := times.Unix() - time.Now().Unix()
        if d > 0 {
            args :=[]string{"-s","-t"}
            args = append(args, strconv.FormatInt(d, 10))
            cmd := exec.Command("shutdown",args...)
            err := cmd.Run()
            if err != nil {
                t.Execute(context.Writer, status{"设置关机任务失败"})
            }else{
                globalStatus.Status = "电脑将于:"+ shutdownTime + "关闭"
                t.Execute(context.Writer, globalStatus)
            }
            return
        }
        t.Execute(context.Writer, status{Status: "错误时间格式"})
        return
    })
    err := e.Run(":9999")
    if err != nil {
        log.Fatal(err)
    }
    log.Println("service start at http://localhost:9999")
}

func shutdown(c *gin.Context){
    args :=[]string{"-s","-t","30"}
    t,err := template.ParseFiles("index.html")
    if err != nil {
        log.Println("err : ", err)
        return
    }
    cmd := exec.Command("shutdown",args...)
    err = cmd.Run()
    if err != nil {
        t.Execute(c.Writer, status{"设置关机任务失败"})
    }else{
        globalStatus.Status = "电脑将在30s后关机"
        t.Execute(c.Writer,globalStatus )
    }
}
func shutdown2(c *gin.Context)  {
    args :=[]string{"-s","-t","30"}
    cmd := exec.Command("shutdown",args...)
    err := cmd.Run()
    if err != nil {
       c.JSON(http.StatusInternalServerError,
           gin.H{"message":"服务器内部错误"})
    }else{
        globalStatus.Status = "电脑将在30s后关机"
        c.JSON(http.StatusOK, gin.H{"message":globalStatus.Status})
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
    case "3":
        shutdown2(c)
    default:
        c.JSON(
            http.StatusBadRequest,
            gin.H{
                "message":"bad Request",
        })
    }
}
