package main

import "fmt"

func main() {
    var a byte = 0x7
    var b byte = 0xb
    fmt.Printf("%04b &^ %04b = %04b \n",a,b,a&^b)
}