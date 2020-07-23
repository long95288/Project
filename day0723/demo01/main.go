package main

import (
    "errors"
    "fmt"
    "reflect"
)
// 需求:根据一个函数字符串，调用函数，提供函数参数,接收返回值
// 反射调用无参函数
// 反射调用有参无返回值函数
// 反射调用有入参和返回的函数
// 动态调用函数，无参数
type T struct{}
func (t *T) Do(){
    fmt.Println("Hello Do")
}
// 有参函数
func (t *T) Do2(a int,b string){
    fmt.Println("hello" + b,a)
}
func (t *T) Do3() (string,error){
    return "Hello",errors.New("return error")
}
func main() {
    // 函数名称
    name := "Do"
    t := &T{}
    // 反射调用函数
    reflect.ValueOf(t).MethodByName(name).Call(nil)
    // 反射调用有参函数
    a := reflect.ValueOf(111)
    b := reflect.ValueOf("world")
    in := []reflect.Value{a,b}
    reflect.ValueOf(t).MethodByName("Do2").Call(in)
    // 反射含有返回值的函数
    rets := reflect.ValueOf(t).MethodByName("Do3").Call(nil)
    fmt.Printf(" strValue:%[1]v \n errValue:%[2]v\n strType:%[1]T \n errType:%[2]T\n",
        rets[0],rets[1].Interface().(error))
    
}
