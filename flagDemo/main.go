package main

import (
    "flag"
    "fmt"
)

func main() {
    url := flag.String("url","","")
    enable := flag.Bool("enable",false,"")
    
    fmt.Printf("url = %s enable = %t",*url,*enable)
}
