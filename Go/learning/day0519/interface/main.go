package main

import "fmt"

type Usb interface {
    Start()
    Stop()
}

type Phone struct{
    name string
}

func (p Phone) Start() {
    fmt.Println("手机开始工作。。。。")
}
func (p Phone) Stop() {
    fmt.Println("手机停止工作....")
}
func (p Phone) Call() {
    fmt.Println("手机打电话...")
}

type Camera struct{
    name string
}
func (c Camera) Start() {
    fmt.Println("相机开始工作。。。。")
}
func (c Camera) Stop() {
    fmt.Println("相机停止工作")
}
func (c Camera) Photography(){
    fmt.Println("拍照....")
}
type Computer struct{
}

func (c Computer) Working(usb Usb) {
    usb.Start()
    // 使用类型断言
    if v,flag := usb.(Phone);flag{
        v.Call()
    }
    if v, flag := usb.(Camera); flag {
        v.Photography()
    }
    usb.Stop()
}

func TypeJudge(item... interface{}){
    for index,v := range item{
        switch v.(type) {
        case bool:
            fmt.Printf("第%v参数 bool类型,值为:%v\n",index,v)
        case float32:
            fmt.Printf("第%v参数 float32类型,值为:%v\n",index,v)
        case float64:
            fmt.Printf("第%v参数 float64类型,值为:%v\n",index,v)
        case int, int32, int64:
            fmt.Printf("第%v参数 int 类型,值为:%v\n",index,v)
        case string:
            fmt.Printf("第%v参数 string类型,值为:%v\n",index,v)
        case Phone:
            fmt.Printf("第%v参数 Phone类型,值为:%v\n",index,v)
        case Camera:
            fmt.Printf("第%v参数 Camera类型,值为:%v\n",index,v)
        default:
            fmt.Printf("第%v参数 未知类型,值为:%v\n",index,v)
        }
    }
}

func main() {
    var usbArr [3]Usb = [3]Usb{
        0:Phone{
            name: "vivo",
        },
        1:Camera{
            name: "尼康",
        },
        2:Phone{
            name: "meizu",
        },
    }
    var computer Computer
    for _,v := range usbArr{
        computer.Working(v)
        fmt.Println()
    }
    var usb Usb = Phone{name: "apple"}
    fmt.Println(usb)
    usb = Camera{name: "尼康"}
    fmt.Println(usb)
    
    TypeJudge(1,"stringvalue",1.1,Phone{"phone"},Camera{"camera"},computer)
}
