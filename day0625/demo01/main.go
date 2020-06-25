package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
    "strings"
)

// 创建一级文件夹
func createOneLevelDir(){
    // 在所在盘的根目录创建文件夹,如果运行在F盘,则在F盘创建 F://oneLevelDir
    err := os.Mkdir("/oneLevelDir",0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    // 在特定盘符创建文件夹
    err = os.Mkdir("D://oneLevelDir",0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    // 在当前目录下创建文件夹
    err = os.Mkdir("oneLevelDir",0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("创建文件夹成功")
}
// 创建多级文件夹
func createMultiLevelDir(){
    // 创建多级文件夹
    err := os.MkdirAll("/multiDir1/second/third",0666)
    if err != nil {
        fmt.Println(err)
    }
}
func testReaddirnames(file *os.File){
    filestr,err := file.Readdirnames(0)
    if err != nil {
        fmt.Println(err)
        return
    }
    for _,name := range filestr{
        fmt.Println("name :",name)
    }
}
func testReaddir(file *os.File)  {
    // 读取文件的个数,0为全部读取
    var n int = 0
    fmt.Println("file.Name ",file.Name())
    fmt.Println("Name\t\tIsDir\t\tSize\t\tMode\t\tModTime")
    fileList,err := file.Readdir(n)
    if err != nil {
        return
        fmt.Println(err)
    }
    for _,f := range fileList {
        fmt.Printf("%v\t\t%v\t\t%v\t\t%v\t\t%v\n",f.Name(),f.IsDir(),f.Size(),f.Mode(),f.ModTime())
    }
}
// 读取文件夹内容
func readDir(){
    //
    file,err := os.Open("./")
    if err != nil {
        fmt.Println(err)
        return
    }
    // testReaddirnames(file)
    testReaddir(file)
    //file.Close()
}
// 删除文件夹
func deleteDir(){
    err :=os.Remove("test2.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    err = os.RemoveAll("oneLevelDir2")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("删除成功")
}
func renameDir()  {
    //err := os.Rename("oneLevelDir","oneLevelDir2")
    err := os.Rename("test2.txt","oneLevelDir2/test2.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("重命名成功")
}
func testCopy() {
    // 源文件
    srcFile,err := os.Open("oneLevelDir2/test2.txt")
    defer srcFile.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
    
    // 目标文件
    desFile,err := os.OpenFile("test2.txt",os.O_CREATE|os.O_TRUNC|os.O_RDWR,0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    n,err := io.Copy(desFile,srcFile)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(n)
    
}
func createFile() {
    _,err := os.OpenFile("test2.txt",os.O_CREATE,0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("创建文件成功")
}

func readFile1(){
    file,err := os.Open("test2.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    data := make([]byte,0)
    buffer := make([]byte,1024)
    for {
        n,err := file.Read(buffer)
        if err == io.EOF {
            // 读取完成
            break
        }
        // 循环赋值
        for i:=0;i<n;i++{
            data = append(data,buffer[i])
        }
    }
    fmt.Printf("%q",string(data))
}
func readFile2() {
    file,err := os.Open("test2.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    // 带缓冲的
    r:= bufio.NewReader(file)
    sb := strings.Builder{}
    for true {
       tmp,err := r.ReadString('\n')
       if err == io.EOF {
           // 如果最后一行数据没有\n符也要加进去
           sb.WriteString(tmp)
           break
       }
       sb.WriteString(tmp)
    }
    fmt.Printf("%q",sb.String())
}
func readFile3() {
    data,err := ioutil.ReadFile("test2.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("%q",string(data))
}
func writeFile1() {
    // 覆盖写入新数据
    file,err := os.OpenFile("test3.txt",os.O_RDWR|os.O_CREATE|os.O_TRUNC,0666)
    if err != nil {
        return
    }
    defer file.Close()
    n,err := file.Write([]byte("Welcome"))
    n,err = file.WriteString("Hello World \r\n")
    if err != nil {
        return
    }
    fmt.Println(n)
    //// 偏移插入
    //_,err = file.WriteAt([]byte("Append"),6)
    //if err != nil {
    //    return
    //}
}

func writeFile2(){
    file,err := os.OpenFile("test3.txt",os.O_RDWR|os.O_CREATE|os.O_TRUNC,0666)
    if err != nil {
        return
    }
    defer file.Close()
    r := bufio.NewWriter(file)
    r.WriteString("Hello,World")
    r.Flush()
}
func writeFile3() {
    err := ioutil.WriteFile("test3.txt",[]byte("Hello,World3"),0666)
    if err != nil {
        return
    }
}
func main() {
    // createOneLevelDir()
    // createMultiLevelDir()
    // readDir()
    // renameDir()
    // testCopy()
    // deleteDir()
    // createFile()
    // readFile2()
    // readFile3()
    // writeFile1()
    // writeFile2()
    // writeFile3()
    r1 := strings.NewReader("Hello World")
    br := bufio.NewReader(r1)
    data := ""
    for true {
        s,err := br.ReadString('\n')
        if err == io.EOF {
            data += s
            break
        }
        data += s
    }
    fmt.Printf("%q",data)
}
