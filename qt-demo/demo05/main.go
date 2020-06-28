//
// 简单的窗口界面，加载应用图标
// tag 窗口居中
//
package main

import (
    "github.com/therecipe/qt/core"
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
    //center(app)
    
    // 布局的载体
    layoutWidget := widgets.NewQWidget(app,core.Qt__Widget)
    layoutWidget.SetGeometry2(0,0,300,200)
    app.SetCentralWidget(layoutWidget)
    
    okBtn := widgets.NewQPushButton2("Ok",layoutWidget)
    cancelBtn := widgets.NewQPushButton2("Cancel",layoutWidget)
    
    hBox := widgets.NewQHBoxLayout2(layoutWidget)
    hBox.SetContentsMargins(0,0,0,0)
    hBox.AddStretch(1)
    hBox.AddWidget(okBtn,0,0)
    hBox.AddWidget(cancelBtn,0,0)
    //
    //vbox := widgets.NewQVBoxLayout2(layoutWidget)
    //vbox.SetContentsMargins(0, 0, 0, 0)
    //vbox.AddStretch(1)
    //vbox.AddLayout(hBox, 0)
    
    return app
}

func main() {
    widgets.NewQApplication(len(os.Args),os.Args)
    app := InitUi()
    app.Show()
    widgets.QApplication_Exec()
}