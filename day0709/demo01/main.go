package main

import "fmt"

// 去除数组中的空串
//
func noneEmpty(strings []string) []string {
    i := 0
    for _,s := range strings{
        if s != ""{
            strings[i]=s
            i++
        }
    }
    return strings[:i]
}
// 共享内存的数组
func noneEmpty2(strings []string) []string {
    out := strings[:0]
    for _,s := range strings{
        if s != "" {
            out = append(out,s)
        }
    }
    return out
}

func main() {
    data := []string{"one","","three"}
    fmt.Printf("%q\n",noneEmpty2(data))
    // 在noneEmpty函数中更改了data的内容
    fmt.Printf("%q\n",data)
}
