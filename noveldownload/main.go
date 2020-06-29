//
// 简单的窗口界面，加载应用图标
// tag 窗口居中
//
package main

import (
    "fmt"
    "github.com/therecipe/qt/core"
    "github.com/therecipe/qt/gui"
    "github.com/therecipe/qt/widgets"
    "io/ioutil"
    "myproject/go_learn/noveldownload/utils"
    "os"
    "strings"
    "sync"
)

func center(app *widgets.QMainWindow){
    qr := app.FrameGeometry()
    cp := widgets.NewQDesktopWidget().AvailableGeometry2(app).Center()
    qr.MoveCenter(cp)
    app.Move(qr.TopLeft())
}
var urlLineEdit *widgets.QLineEdit
var novelNameLineEdit *widgets.QLineEdit
var novelChapterNumLabel *widgets.QLabel
var novelStatusLabel *widgets.QLabel
var analyzeBtn *widgets.QPushButton
var startDownloadBtn *widgets.QPushButton
var processBar *widgets.QProgressBar
type Chapter struct {
    index int
    url string
    content string
}
//
var urlList []string
// 章节url管道
var chapterUrlChan chan Chapter
// 章节内容管道
var chapterContentChan chan Chapter
// 下载线程数
var downloadThread = 5
// lock
var lock sync.Mutex

func InitUi() *widgets.QMainWindow{
    
    chapterUrlChan = make(chan Chapter,5)
    chapterContentChan = make(chan Chapter,5)
    
    app := widgets.NewQMainWindow(nil,0)
    app.SetWindowTitle("下载器")
    app.SetGeometry2(300,300,400,300)
    app.SetWindowIcon(gui.NewQIcon5("F:\\GoPath\\src\\myproject\\go_learn\\qt-demo\\demo02\\app.png"))
    
    // 布局
    layoutWidget := widgets.NewQWidget(app,core.Qt__Widget)
    layoutWidget.SetGeometry2(0,0,400,300)
    app.SetCentralWidget(layoutWidget)
    
    grid := widgets.NewQGridLayout(layoutWidget)
    grid.SetContentsMargins(10,10,10,10)
    grid.SetSpacing(10)
    grid.AddWidget2(widgets.NewQLabel2("URL:",app,0),0,0,0)
    
    // URL输入框
    urlLineEdit = widgets.NewQLineEdit(app)
    grid.AddWidget2(urlLineEdit,0,1,0)
    
    // 小说保存名称
    grid.AddWidget2(widgets.NewQLabel2("小说名",app,0),1,0,0)
    novelNameLineEdit = widgets.NewQLineEdit(app)
    grid.AddWidget2(novelNameLineEdit,1,1,0)
    
    // 章节数
    grid.AddWidget2(widgets.NewQLabel2("章节数",app,0),2,0,0)
    novelChapterNumLabel = widgets.NewQLabel(app,0)
    grid.AddWidget2(novelChapterNumLabel,2,1,0)
    
    // 状态说明
    grid.AddWidget2(widgets.NewQLabel2("当前状态",app,0),3,0,0)
    novelStatusLabel = widgets.NewQLabel(app,0)
    grid.AddWidget2(novelStatusLabel,3,1,0)
    
    // 分析和下载按钮
    analyzeBtn = widgets.NewQPushButton2("分析",app)
    startDownloadBtn = widgets.NewQPushButton2("下载",app)
    startDownloadBtn.SetEnabled(false)
    grid.AddWidget2(analyzeBtn,4,0,0)
    grid.AddWidget2(startDownloadBtn,4,1,0)
    
    // 进度条
    grid.AddWidget2(widgets.NewQLabel2("进度:",app,0),5,0,0)
    processBar = widgets.NewQProgressBar(app)
    processBar.SetValue(0)
    grid.AddWidget2(processBar,5,1,0)
    
    // 设置窗口居中
    center(app)
    return app
}
// 产生数据
func productChapter(){
    for index,value := range urlList{
        // 写入管道中
        chapterUrlChan<- Chapter{index: index,url: value}
    }
    // 写入完成之后关闭管道
    close(chapterUrlChan)
}
// 下载数据
// 消费者
func consumeChapter()  {
    for {
        oneChapter,ok := <- chapterUrlChan
        if !ok  {
            lock.Lock()
            downloadThread --
            if downloadThread == 0 {
                close(chapterContentChan)
            }
            lock.Unlock()
            break
        }
        content,err := utils.DownloadChapter(oneChapter.url)
        if err != nil {
            lock.Lock()
            downloadThread --
            if downloadThread == 0 {
                close(chapterContentChan)
            }
            lock.Unlock()
            break
        }
        chapterContentChan <- Chapter{index: oneChapter.index,content: content}
    }
    
}
// 生成小说
func generateNovel(){
    novelMap := make(map[int]string)
    for {
       oneChapter,ok :=<- chapterContentChan
       if !ok {
           break
       }
       novelMap[oneChapter.index] = oneChapter.content
    }
    novelName := novelNameLineEdit.Text()
    novelName += ".txt"
    sb := strings.Builder{}
    // 获得所有数据之后写入本地文件
    for i:= 0;i < len(novelMap);i++{
        sb.WriteString(novelMap[i])
    }
    err := ioutil.WriteFile(novelName,[]byte(sb.String()),0666)
    if err != nil {
        widgets.QMessageBox_Critical(
            nil,
            "写入错误",
            "写入小说文件错误",
            widgets.QMessageBox__Yes,
            widgets.QMessageBox__Yes)
    }
}
// 下载小说
func download()  {
    // 插入url列表
    go productChapter()
    // 获得数据
    for i:=0; i<downloadThread; i++{
        go consumeChapter()
    }
    // 写入文件
    go generateNovel()
}

// 事件处理
func initEvent() {
    analyzeBtn.ConnectClicked(func(checked bool) {
        fmt.Println("分析按钮按下")
        url:= urlLineEdit.Text()
        novelName,chapterUrl,err := utils.AnalyzeUrl(url)
        if err != nil {
            return
        }
        novelNameLineEdit.SetText(novelName)
        urlList = chapterUrl
        analyzeBtn.SetEnabled(false)
    })
    startDownloadBtn.ConnectClicked(func(checked bool) {
        fmt.Println("开始按钮按下")
        startDownloadBtn.SetEnabled(false)
        novelNameLineEdit.SetEnabled(false)
        go download()
    })
}
func main() {
    widgets.NewQApplication(len(os.Args),os.Args)
    app := InitUi()
    initEvent()
    app.Show()
    widgets.QApplication_Exec()
}