//
// 简单的窗口界面，加载应用图标
// tag 窗口居中
// tag 添加按钮
//
package main

import (
    "fmt"
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
    app.SetWindowTitle("页面标题")
    app.SetGeometry2(300,300,300,300)
    app.SetWindowIcon(gui.NewQIcon5("F:\\GoPath\\src\\myproject\\go_learn\\qt-demo\\demo02\\app.png"))
    
    // 设置窗口居中
    center(app)
    
    // 按钮
    btn := widgets.NewQPushButton2("关闭按钮",app)
    btn.Resize(btn.SizeHint())
    btn.Move2(50,50)
    btn.ConnectClicked(func(check bool) {
        fmt.Println("关闭窗口")
        app.Close()
    })
    btn.Show()
    ////
    
    return app
}

func main() {
    widgets.NewQApplication(len(os.Args),os.Args)
    app := InitUi()
    app.Show()
    widgets.QApplication_Exec()
}