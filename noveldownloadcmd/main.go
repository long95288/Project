package main

import (
    "fmt"
    "log"
    "myproject/go_learn/noveldownloadcmd/util"
    "os"
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
    defer file.Close()
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    for i,url := range urls{
        fmt.Println("下载:",url)
        downloadCount := 10
        for downloadCount > 0{
            content,err := util.DownloadChapter(url)
            if err != nil {
                fmt.Println("err", err)
                downloadCount --
                time.Sleep(2*time.Second)
                continue
            }else{
                _, err = file.WriteString(content)
                if err != nil {
                    fmt.Println("write err", err)
                    os.Exit(1)
                }
                break
            }
        }
        fmt.Printf("下载:%d/%d url = %s 成功\n", i, len(urls), url)
        time.Sleep(1 * time.Second)
    }
}
