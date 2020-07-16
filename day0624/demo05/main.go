package main

import (
    "bufio"
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "golang.org/x/text/encoding/simplifiedchinese"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
)

// 第一种方式
func readFileDemo01()  {
    buf := make([]byte,1024)
    f,_ := os.Open("findlinks1.go")
    defer f.Close()
    for true {
        n,_ := f.Read(buf)
        if n == 0{
            // 读出来为字节数为0说明已经读完了,
            break
        }
        os.Stdout.Write(buf[:n])
    }
}
func writeFileDemo01() {
    // Open 为O_ReadOnly模式
    f,_ := os.Open("test.txt")
    f.Chmod(0666)
    defer f.Close()
    //w := bufio.NewWriter(f)
    //n,err := w.WriteString("Hello World")
    //if err != nil {
    //    fmt.Println(err)
    //    return
    //}
    //fmt.Println(n)
    //err = w.Flush()
    //if err != nil {
    //    fmt.Println(err)
    //}
    n,err := f.Write([]byte("Write Demo1"))
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(len("Write Demo1"),n)
}
func writeFileDemo02() {
    err := ioutil.WriteFile("test.txt",[]byte("HelloWorld"),0666)
    if err != nil {
        fmt.Println("err ",err)
        return
    }
    fmt.Println("write success")
}
func writeFileDemo03(){
    file,err := os.OpenFile("test1.txt",os.O_APPEND|os.O_CREATE,0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    n, err :=file.WriteString("Write Demo03")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(n,len("Write Demo03"))
}
// 第二种方式读取
func readFileDemo02() {
    buf := make([]byte,2)
    f,_ := os.Open("findlinks1.go")
    defer f.Close()
    r := bufio.NewReader(f)
    w := bufio.NewWriter(os.Stdout)
    defer w.Flush()
    for {
        n,_ := r.Read(buf)
        if n == 0 {break}
        w.Write(buf[:n])
    }
}
// 读取字符串
func readFileDemo03(){
    f,_ := os.Open("findlinks1.go")
    defer f.Close()
    r := bufio.NewReader(f)
    var sum string
    for {
        s,err := r.ReadString('\n')
        if err != nil {
            fmt.Println(err)
            break
        }
        sum += s
    }
    fmt.Println(sum)
    
}
func readFileDemo04() {
    data,err := ioutil.ReadFile("chapter1.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(data))
}
func readHttpBodyDemo()  {
    resp,err := http.Get("https://www.jupindai.com/book/87/28526030.html")
    if err != nil {
        fmt.Println(err)
        return
    }
    body := resp.Body
    r := bufio.NewReader(body)
    var sum string
    for {
        s, err := r.ReadString('\n')
        if err != nil {
            fmt.Println(err)
            break
        }
        sum += s
    }
    sum,err = ConvertGBK2Str(sum)
    file,err := os.OpenFile("chapter1.txt",os.O_CREATE|os.O_APPEND,0666)
    n,err := file.WriteString(sum)
    if err != nil {
        return
    }
    fmt.Println(len(sum),n)
}
func ConvertStr2GBK(str string) (string,error) {
    // utf-8 -> GBK
    ret,err := simplifiedchinese.GBK.NewEncoder().String(str)
    return ret,err
}

func ConvertGBK2Str(gbkStr string) (string,error) {
    // gbk -> utf-8
    ret,err := simplifiedchinese.GBK.NewDecoder().String(gbkStr)
    return ret,err
}
func ConvertGBKBit2Str(str []byte) (string,error)  {
    // byte转str
    strByte,err := simplifiedchinese.GBK.NewDecoder().Bytes(str)
    if err != nil {
        return "", err
    }
    return string(strByte),nil
}

func parseHtml() {
    resp,err := http.Get("https://www.jupindai.com/book/87/28526030.html")
    if err != nil {
        fmt.Println(err)
        return
    }
    body := resp.Body
    defer resp.Body.Close()
    r := bufio.NewReader(body)
    var sum string
    for {
        s, err := r.ReadString('\n')
        if err == io.EOF {
            fmt.Println(err)
            break
        }
        sum += s
    }
    sum,err = ConvertGBK2Str(sum)
    doc,err := goquery.NewDocumentFromReader(strings.NewReader(sum))
    if err != nil {
        fmt.Println(err)
        return
    }
    
    title := doc.Find(".readTitle").Text()
    fmt.Println(title)
    context := doc.Find("#htmlContent").Text()
    file,err := os.OpenFile("chapter1.txt",os.O_APPEND|os.O_CREATE,0666)
    _,err =file.WriteString(title+"\n"+context + "\n")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("写入成功")
}

func main() {
    //fmt.Println("tcpServer")
    //readFileDemo01()
    //fmt.Println("demo02")
    //readFileDemo02()
    //fmt.Println("demo03")
    //readFileDemo03()
    // readFileDemo04()
    //readHttpBodyDemo()
    // writeFileDemo01()
    // writeFileDemo02()
    // writeFileDemo03()
    parseHtml()
}
