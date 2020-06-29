package utils

import (
    "bufio"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "golang.org/x/text/encoding/simplifiedchinese"
    "io"
    "net/http"
    "strings"
)

func ConvertGBK2Str(gbkStr string) (string,error) {
    // gbk -> utf-8
    ret,err := simplifiedchinese.GBK.NewDecoder().String(gbkStr)
    return ret,err
}
// 分析首页url
// 返回小说的名以及各个章节url列表
//
func AnalyzeUrl(url string)(novelName string,chapterUrlList []string,err error)  {
    // 分析小说需要得到一个urlList和小说名
    response,err := http.Get(url)
    rootUrl := strings.Split(url,"/")[2]
    fmt.Println(rootUrl)
    if err != nil {
        return "",nil,err
    }
    body := response.Body
    defer response.Body.Close()
    reader := bufio.NewReader(body)
    sum := ""
    sb := strings.Builder{}
    for {
        s, err := reader.ReadString('\n')
        if err == io.EOF {
            sb.WriteString(s)
            break
        }
        sb.WriteString(s)
    }
    sum,err = ConvertGBK2Str(sb.String())
    if err != nil {
        return "",nil,err
    }
    doc,err := goquery.NewDocumentFromReader(strings.NewReader(sum))
    if err != nil {
        return "",nil,err
    }
    novelName = doc.Find(".bookTitle").Text()
    chapterUrlList = doc.Find("#list-chapterAll > dl > dd > a").Map(func(i int, selection *goquery.Selection) string {
        href,_ := selection.Attr("href")
        return "https://"+rootUrl + href
    })
    return novelName,chapterUrlList,nil
}

// 下载一个章节的内容
// 章节标题和章节内容封装好了
func DownloadChapter(url string) (string,error) {
    resp,err := http.Get(url)
    if err != nil {
        fmt.Println(err)
        return "",err
    }
    body := resp.Body
    defer resp.Body.Close()
    r := bufio.NewReader(body)
    sb := strings.Builder{}
    sum := ""
    for {
        s, err := r.ReadString('\n')
        if err == io.EOF {
            sb.WriteString(s)
            break
        }
        sb.WriteString(s)
    }
    sum,err = ConvertGBK2Str(sb.String())
    doc,err := goquery.NewDocumentFromReader(strings.NewReader(sum))
    if err != nil {
        return "",err
    }
    
    title := doc.Find(".readTitle").Text()
    context := doc.Find("#htmlContent").Text()
    return title + "\n" + context,nil
}