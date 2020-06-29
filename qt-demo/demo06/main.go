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
    // center(app)
    layoutWidget := widgets.NewQWidget(app,core.Qt__Widget)
    layoutWidget.SetGeometry2(0,0,300,200)
    
    grid := widgets.NewQGridLayout(layoutWidget)
    grid.SetContentsMargins(0,0,0,0)
    
    names := []string{
        "Cls","Bck","","Close",
        "7","8","9","/",
        "数据","5","6","*",
        "1","2","3","4",
        "0",".","=","+",
    }
    
    var positions [20]interface{}
    k := 0
    for i:=0;i<5;i++{
        for j:=0;j<4;j++{
            item := [2]int{i,j}
            positions[k] = item
            k += 1
        }
    }
    for index,_:=range positions{
        button := widgets.NewQPushButton2(names[index],layoutWidget)
        value := positions[index]
        valueObj := value.([2]int)
        grid.AddWidget2(button,valueObj[0],valueObj[1],0)
    }
    return app
}

func main() {
    widgets.NewQApplication(len(os.Args),os.Args)
    app := InitUi()
    app.Show()
    widgets.QApplication_Exec()
}