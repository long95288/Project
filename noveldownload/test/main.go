package main

import (
    "fmt"
    "myproject/go_learn/noveldownload/utils"
)

func main() {
    url := "https://www.jupindai.com/book/87.html"
    novelName,urlList,err := utils.AnalyzeUrl(url)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(novelName)
    fmt.Println(len(urlList))
    fmt.Printf("%q\n",urlList[0])
    content,err := utils.DownloadChapter(urlList[0])
    fmt.Println(content)
    fmt.Println(err)
    
}
