package main

import "fmt"

func appendInt(x []int,y int) []int{
    var z []int
    zlen := len(x) + 1
    if zlen <= cap(x){
        // 还有剩余的容量
        z = x[:zlen]
    }else{
        // 没有剩余的容量,需要扩容
        zcap := zlen
        if zcap < 2*len(x) {
            zcap = 2*len(x)
        }
        z = make([]int,zlen,zcap)
        copy(z,x)
    }
    // y赋值
    z[len(x)] = y
    return z
}
// 可变长的参数添加
func appendInt2(x []int, y ...int) []int {
    var z []int
    zlen := len(x) + len(y)
    if zlen <= cap(x){
        // 还有剩余的容量
        z = x[:zlen]
    }else{
        // 没有剩余的容量,需要扩容
        zcap := zlen
        if zcap < 2*len(x) {
            zcap = 2*len(x)
        }
        z = make([]int,zlen,zcap)
        copy(z,x)
    }
    // y赋值
    copy(z[len(x):],y)
    return z
}
func main() {
    var x,y []int
    for i:=0;i<10;i++{
        y = appendInt2(x,i)
        fmt.Printf("%d cap=%d\t%v\n",i,cap(y),y)
        x = y
    }
    
}
