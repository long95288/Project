package main

import (
    "fmt"
    "myproject/go_learn/noveldownload/utils"
    "os"
    "time"
)

func main() {
    url := "https://www.jupindai.com/book/2911.html"
    novelName,urlList,err := utils.AnalyzeUrl(url)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(novelName)
    fmt.Println(len(urlList))
    fmt.Printf("%q\n",urlList[0])
    //content,err := utils.DownloadChapter(urlList[0])
    //fmt.Println(content)
    //fmt.Println(err)
    file,err := os.OpenFile(novelName+".txt",os.O_CREATE|os.O_APPEND,0666)
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    for _,url := range urlList{
        fmt.Println("下载:",url)
        content,err := utils.DownloadChapter(url)
        if err != nil {
            fmt.Println("err",err)
            break
        }
        _,err = file.WriteString(content)
        if err != nil {
            fmt.Println("write err",err)
            break
        }
        fmt.Printf("下载:%s 成功\n",url)
        time.Sleep(500*time.Millisecond)
    }
}
