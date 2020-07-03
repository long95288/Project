package main

import (
    "github.com/therecipe/qt/gui"
    "github.com/therecipe/qt/widgets"
    "os"
)

/*
创建了一个窗口
我们创建了一个QTextEdit,并把他设置为窗口的布局
*/

func InitUi() *widgets.QMainWindow {
    // 创建窗口
    app := widgets.NewQMainWindow(nil, 0)
    
    // 设置窗口的标题
    app.SetWindowTitle("Qt 教程")
    
    // 设置窗口的位置和大小
    app.SetGeometry2(300, 300, 300, 220)
    
    // 设置窗口的图标，引用当前目录下的web.png图片
    app.SetWindowIcon(gui.NewQIcon5("app.png"))
    
    //// 布局窗口组件
    //layoutWidget := widgets.NewQWidget(app, core.Qt__Widget)
    ////layoutWidget.SetGeometry(core.NewQRect4(300, 300, 300, 220))
    //layoutWidget.SetGeometry2(0, 0, 300, 220)
    //app.SetCentralWidget(layoutWidget)
    
    textEdit := widgets.NewQTextEdit(app)
    textEdit.Resize2(300, 200)
    app.SetCentralWidget(textEdit)
    
    // 子按钮
    exitAction := widgets.NewQAction3(gui.NewQIcon5("app.png"), "&Exit", app)
    // 快捷键，自定义
    exitAction.SetShortcut(gui.NewQKeySequence2("Ctrl+Q", gui.QKeySequence__NativeText))
    // 提示语
    exitAction.SetStatusTip("Exit application")
    // 事件触发
    exitAction.ConnectTriggered(func(checked bool) {
        app.Close()
    })
    
    var actions []*widgets.QAction
    actions = append(actions, exitAction)
    
    //创建一个菜单栏
    //menubar := widgets.NewQMenuBar(app)
    menubar := app.MenuBar()
    //添加菜单
    fileMenu := menubar.AddMenu2("&File")
    //添加按钮
    fileMenu.AddActions(actions)
    
    // 创建工具栏
    //toolbar := widgets.NewQToolBar("Exit", layoutWidget)
    toolbar := app.AddToolBar3("Exit")
    toolbar.AddActions(actions)
    
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
