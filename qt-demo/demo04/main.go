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
func CloseEvent(event *gui.QCloseEvent){
    reply := widgets.QMessageBox_Question(nil,
        "title",
        "提示说明，是否确认关闭",
        widgets.QMessageBox__Yes|widgets.QMessageBox__No,
        widgets.QMessageBox__Yes)
    if reply == widgets.QMessageBox__Yes {
        fmt.Println("选择确认关闭")
        event.Accept()
    }else{
        //
        fmt.Println("选择取消")
        event.Ignore()
    }
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
    // 设置提示信息
    btn.SetToolTip("这是提示信息<b>QPushButton</b>widget")
    btn.ConnectClicked(func(check bool) {
        fmt.Println("关闭窗口")
        fmt.Printf("%t",check)
        //app.Close()
    })
    btn.Show()
    ////
    // 修改默认关闭按钮,加入确认框
    app.ConnectCloseEvent(CloseEvent)
    ///////////////////
    // 绝对定位
    lbl1 := widgets.NewQLabel2("lable1",app,0)
    lbl2 := widgets.NewQLabel2("lable2",app,0)
    // 设置定位
    lbl1.Move2(20,10)
    lbl2.Move2(20,30)
    
    return app
}

func main() {
    widgets.NewQApplication(len(os.Args),os.Args)
    app := InitUi()
    app.Show()
    widgets.QApplication_Exec()
}