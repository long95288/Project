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
    "strconv"
    "time"
)


var app *widgets.QMainWindow
var urlLineEdit *widgets.QLineEdit
var novelNameLineEdit *widgets.QLineEdit
var novelChapterNumLabel *widgets.QLabel
var novelStatusLabel *widgets.QLabel
var analyzeBtn *widgets.QPushButton
var startDownloadBtn *widgets.QPushButton
var processBar *widgets.QProgressBar
var layoutWidget *widgets.QWidget
var urlList []string
var novelName = ""
var chapterNum = 0
var savePath = "download/"
var bgPixMap *gui.QPixmap

func InitUi() *widgets.QMainWindow{
    
    app := widgets.NewQMainWindow(nil,0)
    app.SetWindowTitle("小说下载器")
    // app.SetGeometry2(300,300,400,180)
    app.SetWindowIcon(gui.NewQIcon5("image/app.png"))
    
    
    // 布局
    layoutWidget = widgets.NewQWidget(app,core.Qt__Widget)
    // layoutWidget.SetGeometry2(0,0,400,300)
    app.SetCentralWidget(layoutWidget)
    app.SetAutoFillBackground(true)
    //
    rightAlign := core.Qt__AlignRight
    grid := widgets.NewQGridLayout(layoutWidget)
    grid.SetContentsMargins(10,10,10,10)
    grid.SetSpacing(10)
    grid.AddWidget2(widgets.NewQLabel2("URL:",app,0),0,0, rightAlign)
    
    // URL输入框
    urlLineEdit = widgets.NewQLineEdit(app)
    grid.AddWidget2(urlLineEdit,0,1,0)
    
    // 小说保存名称
    grid.AddWidget2(widgets.NewQLabel2("小说名:",app,0),1,0, rightAlign)
    novelNameLineEdit = widgets.NewQLineEdit(app)
    grid.AddWidget2(novelNameLineEdit,1,1,0)
    
    // 章节数
    grid.AddWidget2(widgets.NewQLabel2("章节数:",app,0),2,0, rightAlign)
    novelChapterNumLabel = widgets.NewQLabel(app,0)
    grid.AddWidget2(novelChapterNumLabel,2,1,0)
    
    // 状态说明
    grid.AddWidget2(widgets.NewQLabel2("当前状态:",app,0),3,0, rightAlign)
    novelStatusLabel = widgets.NewQLabel(app,0)
    grid.AddWidget2(novelStatusLabel,3,1,0)
    
    // 分析和下载按钮
    analyzeBtn = widgets.NewQPushButton2("分析",app)
    startDownloadBtn = widgets.NewQPushButton2("下载",app)
    startDownloadBtn.SetEnabled(false)
    grid.AddWidget2(analyzeBtn,4,0, rightAlign)
    grid.AddWidget2(startDownloadBtn,4,1,0)
    
    // 进度条
    grid.AddWidget2(widgets.NewQLabel2("进度:",app,0),5,0, rightAlign)
    processBar = widgets.NewQProgressBar(app)
    processBar.SetValue(0)
    grid.AddWidget2(processBar,5,1,0)
    
    // 设置窗口居中
    center(app)
    // 设置样式
    setStyle(app)
    return app
}
func center(app *widgets.QMainWindow){
    qr := app.FrameGeometry()
    cp := widgets.NewQDesktopWidget().AvailableGeometry2(app).Center()
    qr.MoveCenter(cp)
    app.Move(qr.TopLeft())
}
func setStyle(app *widgets.QMainWindow) {
    style, err := ioutil.ReadFile("style/style.qss")
    if err != nil {
        return
    }
    app.SetStyleSheet(string(style))
    
}

// 下载状态时的ui
func setDownloadStatusUI()  {
    novelNameLineEdit.SetEnabled(false)
    analyzeBtn.SetEnabled(false)
    startDownloadBtn.SetEnabled(false)
}
// 分析成功时的UI设置
func setAnalyzeSuccessUI(){
    novelNameLineEdit.SetText(novelName)
    novelNameLineEdit.SetEnabled(true)
    analyzeBtn.SetEnabled(false)
    startDownloadBtn.SetEnabled(true)
    novelChapterNumLabel.SetText(strconv.Itoa(chapterNum))
    processBar.SetValue(0)
}
// 初始值UI
func setIndexStatusUI() {
    novelNameLineEdit.SetEnabled(false)
    novelNameLineEdit.SetText("")
    novelChapterNumLabel.SetText("")
    analyzeBtn.SetEnabled(true)
    startDownloadBtn.SetEnabled(false)
}

// 下载小说
func download()  {
    os.MkdirAll(savePath,0666)
    file,err := os.OpenFile(savePath + novelNameLineEdit.Text()+".txt",os.O_CREATE|os.O_APPEND,0666)
    if err != nil {
        fmt.Println(err)
    }
    defer file.Close()
    successFlag := true
    for index,url := range urlList{
        msg1 := fmt.Sprintf("下载:%d url:%s", index, url)
        fmt.Println(msg1)
        novelStatusLabel.SetText(msg1)
        content,err := utils.DownloadChapter(url)
        if err != nil {
            fmt.Println("err",err)
            successFlag = false
            break
        }
        _,err = file.WriteString(content)
        if err != nil {
            fmt.Println("write err",err)
            successFlag = false
            break
        }
        msg1 = fmt.Sprintf("下载:%d url:%s 成功",index,url)
        novelStatusLabel.SetText(msg1)
        process := float32(index)/float32(len(urlList)) * 100
        processBar.SetValue(int(process))
        time.Sleep(500*time.Millisecond)
    }
    msg := ""
    if successFlag {
        msg = fmt.Sprintf("下载《%s》完成",novelName)
    }else{
        msg = fmt.Sprintf("下载《%s》出错",novelName)
    }
    widgets.QMessageBox_Information(
        nil,
        "信息",
        msg,
        widgets.QMessageBox__Yes,
        widgets.QMessageBox__Yes)
    setIndexStatusUI()
}

// 事件处理
func initEvent() {
    analyzeBtn.ConnectClicked(func(checked bool) {
        fmt.Println("分析按钮按下")
        url:= urlLineEdit.Text()
        if url == "" {
            widgets.QMessageBox_Critical(nil,"错误","URL不能为空",widgets.QMessageBox__Yes,widgets.QMessageBox__Yes)
            return
        }
        novelName2,chapterUrl,err := utils.AnalyzeUrl(url)
        if err != nil {
            return
        }
        urlList = chapterUrl
        novelName = novelName2
        chapterNum = len(urlList)
        fmt.Printf("分析结果:\n小说名:%s\n章节数:%d\n",novelName,chapterNum)
        // 更新gui
        setAnalyzeSuccessUI()
    })
    startDownloadBtn.ConnectClicked(func(checked bool) {
        setDownloadStatusUI()
        go download()
    })
    app.ConnectPaintEvent(func(event *gui.QPaintEvent) {
        // 第一种方式
        // 重设图片宽高以适应应用大小
        bgPixMap = gui.NewQPixmap3("image/bg.jpg","",core.Qt__AutoColor)
        bgPixMap = bgPixMap.Scaled2(app.Width(),app.Height(),core.Qt__IgnoreAspectRatio,core.Qt__SmoothTransformation)
        bgPalette := gui.NewQPalette()
        brush := gui.NewQBrush7(bgPixMap)
        bgPalette.SetBrush(gui.QPalette__Background,brush)
        app.SetPalette(bgPalette)
        // 第二种方式
        //bgLabel.SetGeometry2(0,0,app.Width(),app.Height())
        event.Accept()
    })
}
func main() {
    widgets.NewQApplication(len(os.Args),os.Args)
    app = InitUi()
    initEvent()
    app.Show()
    widgets.QApplication_Exec()
}