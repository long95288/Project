package main

import (
    "fmt"
    "reflect"
)

// 动物接口
type Animal interface {
    speak()
}
type Cry interface {
    cry()
}
// 汽车实现speak()方法
type Car struct {

}

func (c *Car) speak() {
    fmt.Println("wu wu wu ...")
}

func main() {
    car := &Car{}
    rCarType := reflect.TypeOf(car)
    
    animalType := reflect.TypeOf((*Animal)(nil)).Elem()
    cryType := reflect.TypeOf((*Cry)(nil)).Elem()
    
    // 判断car是否实现了Animal接口
    re := rCarType.Implements(animalType)
    fmt.Println(re)
    // car 是否实现了Cry接口
    re = rCarType.Implements(cryType)
    fmt.Println(re)
}
