package main

import (
    "fmt"
    "reflect"
)

// 类型转换和赋值
// 将T1中的值赋值给tag为等于T2.
// T1.A -> T2.AA
// T1.B -> T2.BB
// 怎么知道T1.A对应T2.AA？
// T1中有一个tag为T2,值为AA
type T1 struct {
    A int `T2:"AA"`
    B string `T2:"BB"`
}
type T2 struct {
    AA int
    BB string
}
func main() {
    t := T1{
        A: 123,
        B: "hello",
    }
    rT1Type := reflect.TypeOf(t)
    rT1Value := reflect.ValueOf(t)
    
    t2 := &T2{}
    rT2Value := reflect.ValueOf(t2)
    
    // 赋值
    for i:=0;i< rT1Type.NumField();i++{
        field := rT1Type.Field(i)
        // 获得T1类型中的Tag的值为T2的Tag值
        newTTag := field.Tag.Get("T2")
        // 获得对应T1中对应Tag的值,为域值
        tValue := rT1Value.Field(i)
        // 给T2设置值
        rT2Value.Elem().FieldByName(newTTag).Set(tValue)
    }
    fmt.Println(t2)
}
