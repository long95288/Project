//
// 简单的窗口界面，加载应用图标
// tag 窗口居中
//
package main

import (
    "github.com/therecipe/qt/gui"
    "github.com/therecipe/qt/widgets"
    "os"
)

func center(app *widgets.QMainWindow){
    qr := app.FrameGeometry()
    cp := widgets.NewQDesktopWidget().AvailableGeometry2(app).Center()
    qr.MoveCenter(cp)
    app.Move(qr.TopLeft())
}

func InitUi() *widgets.QMainWindow{
    app := widgets.NewQMainWindow(nil,0)
    app.SetWindowTitle("下载器")
    app.SetGeometry2(300,300,400,300)
    app.SetWindowIcon(gui.NewQIcon5("F:\\GoPath\\src\\myproject\\go_learn\\qt-demo\\demo02\\app.png"))
    hPolicy := widgets.QSizePolicy__Policy(widgets.QSizePolicy__Fixed)
    vPolicy := widgets.QSizePolicy__Policy(widgets.QSizePolicy__Fixed)
    app.SetSizePolicy2(hPolicy,vPolicy)
    app.ConnectPaintEvent(func(event *gui.QPaintEvent) {
        event.Accept()
    })
    // URL title
    //urlLabel := widgets.NewQLabel2("URL:",app,0)
    //analyzeBtn := widgets.NewQPushButton2("分析",app)
    //processLabel := widgets.NewQLabel2("进度",app,0)
    //startBtn := widgets.NewQPushButton2("开始",app)
    //novelNameLabel := widgets.NewQLabel2("小说名",app,0)
    //chapterNumLabel := widgets.NewQLabel2("章节数",app,0)
    
    //
    
    // 设置窗口居中
    center(app)
    return app
}

func main() {
    widgets.NewQApplication(len(os.Args),os.Args)
    app := InitUi()
    app.Show()
    widgets.QApplication_Exec()
}