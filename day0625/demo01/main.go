package main

import (
    "fmt"
    "os"
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
// 读取文件夹内容
func readDir(){

}
// 删除文件夹
func deleteDir(){

}

func main() {
    // createOneLevelDir()
    createMultiLevelDir()
}
