package main

import (
    "fmt"
    "log"
    "os"
    "myproject/go_learn/noveldownloadcmd/util"
    "time"
)

func main() {
    args := os.Args
    if len(args) != 2 {
        log.Fatal("参数错误")
    }
    name,urls,err := util.AnalyzeUrl(args[1])
    if err != nil{
        log.Fatal(err)
    }
    fmt.Printf("novalname:%s\nchapternumber :%d\n",name,len(urls))
    file,err := os.OpenFile(name+".txt",os.O_CREATE|os.O_APPEND,0666)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    for i,url := range urls{
        fmt.Println("下载:",url)
        content,err := util.DownloadChapter(url)
        if err != nil {
            fmt.Println("err",err)
            break
        }
        _,err = file.WriteString(content)
        if err != nil {
            fmt.Println("write err",err)
            break
        }
        fmt.Printf("下载:%d/%d url = %s 成功\n",i,len(urls),url)
        time.Sleep(500*time.Millisecond)
    }
}
