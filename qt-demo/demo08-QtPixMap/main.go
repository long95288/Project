package main

import (
    "github.com/therecipe/qt/core"
    "github.com/therecipe/qt/gui"
    "github.com/therecipe/qt/widgets"
    "os"
)

func InitUi() *widgets.QMainWindow {
// 创建窗口
app := widgets.NewQMainWindow(nil, 0)

// 设置窗口的标题
app.SetWindowTitle("Qt 教程")

// 设置窗口的位置和大小
app.SetGeometry2(300, 300, 300, 220)

// 设置窗口的图标，引用当前目录下的web.png图片
app.SetWindowIcon(gui.NewQIcon5("images/app.ico"))

// 布局窗口组件载体
widget := widgets.NewQWidget(app, core.Qt__Widget)
//widget.SetGeometry(core.NewQRect4(300, 300, 300, 220))
widget.SetGeometry2(0, 0, 300, 220)
app.SetCentralWidget(widget)
// 状态栏
app.StatusBar()

// 水平布局
hbox := widgets.NewQHBoxLayout2(widget)
pixmap := gui.NewQPixmap3("images/icons8-youtube.png", "", core.Qt__AutoColor)

// 标签组件demo.go
label := widgets.NewQLabel(widget, 0)
label.SetScaledContents(true)
label.SetPixmap(pixmap)

hbox.AddWidget(label, 0, 0)

return app
}

func main() {
// 创建一个应用程序对象
// sys.argv参数是一个列表，从命令行输入参数
widgets.NewQApplication(len(os.Args), os.Args)

// 初始化窗口
app := InitUi()

// 显示组件
app.Show()

// 确保应用程序干净的退出
widgets.QApplication_Exec()
}
