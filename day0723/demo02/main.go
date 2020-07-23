package main

import (
    "fmt"
    "reflect"
)

// 反射解析struct 中的tag
//
//
type T struct {
    A int `json:"a" test:"testaa"`
    B string `json:"b" test:"testbb"`
}
func main() {
   t := T{
       A: 124,
       B: "Hello",
   }
   tt := reflect.TypeOf(t)
   for i:=0;i<tt.NumField();i++{
       field := tt.Field(i)
       
       // 输出json 的tag内容
       if json,ok := field.Tag.Lookup("json");ok{
           fmt.Printf("field: %s,testTag tag:%s\n",field.Name,json)
       }
       
       // 输出test的tag内容
       if testTag,ok := field.Tag.Lookup("test");ok{
           fmt.Printf("field: %s,testTag tag:%s \n",field.Name, testTag)
       }
   }
}
