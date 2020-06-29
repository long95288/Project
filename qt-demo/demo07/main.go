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
    center(app)
    // 布局窗口组件载体
    layoutWidget := widgets.NewQWidget(app, core.Qt__Widget)
    //layoutWidget.SetGeometry(core.NewQRect4(300, 300, 300, 220))
    layoutWidget.SetGeometry2(0, 0, 300, 220)
    app.SetCentralWidget(layoutWidget)
    
    // 表格布局
    grid := widgets.NewQGridLayout(layoutWidget)
    grid.SetContentsMargins(5, 5, 5, 0)
    // 设置组件之间的间距。
    grid.SetSpacing(10)
    
    title := widgets.NewQLabel2("Title", layoutWidget, 0)
    author := widgets.NewQLabel2("Author", layoutWidget, 0)
    review := widgets.NewQLabel2("Review", layoutWidget, 0)
    
    titleEdit := widgets.NewQLineEdit(layoutWidget)
    authorEdit := widgets.NewQLineEdit(layoutWidget)
    reviewEdit := widgets.NewQLineEdit(layoutWidget)
    
    grid.AddWidget2(title, 1, 0, 0)
    grid.AddWidget2(titleEdit, 1, 1, 0)
    
    grid.AddWidget2(author, 2, 0, 0)
    grid.AddWidget2(authorEdit, 2, 1, 0)
    
    grid.AddWidget2(review, 3, 0, 0)
    grid.AddWidget3(reviewEdit, 3, 1, 5, 1, 0)
    return app
}

func main() {
    widgets.NewQApplication(len(os.Args),os.Args)
    app := InitUi()
    app.Show()
    widgets.QApplication_Exec()
}