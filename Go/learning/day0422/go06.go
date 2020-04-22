package main

import "unsafe"

func toString(bs []byte) string {
    return *(*string)(unsafe.Pointer(&bs))
}

func main() {
    bs := []byte("Hello world")
    // 添加字符串
    bs = append(bs,"abc"...)
    s := toString(bs)
    println(s)
}
