//
// 简单的窗口界面，加载应用图标
//
package main

import (
    "github.com/therecipe/qt/gui"
    "github.com/therecipe/qt/widgets"
    "os"
)

func InitUi() *widgets.QMainWindow{
    app := widgets.NewQMainWindow(nil,0)
    app.SetWindowTitle("页面标题")
    app.SetGeometry2(300,300,300,300)
    app.SetWindowIcon(gui.NewQIcon5("app.png"))
    return app
}

func main() {
    widgets.NewQApplication(len(os.Args),os.Args)
    app := InitUi()
    app.Show()
    widgets.QApplication_Exec()
}
