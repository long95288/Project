package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)
func output(s string){
    padding := 8
    for _,v := range s {
        fmt.Printf("%c", v)
        padding -= 1
        if padding == 0 {
            fmt.Println()
            padding = 8
        }
    }
    for padding > 0 && padding < 8 {
        fmt.Print("0")
        padding --
        if padding == 0 {
            fmt.Println()
        }
    }
}
func main() {
    in := bufio.NewReader(os.Stdin)
    str1,_ := in.ReadString('\n')
    str2,_ := in.ReadString('\n')
    
    str1 = strings.Replace(str1,"\r\n", "", -1)
    str1 = strings.Replace(str1,"\n", "", -1)
    
    str2 = strings.Replace(str2,"\r\n", "", -1)
    str2 = strings.Replace(str2,"\n", "", -1)
    
    //fmt.Println(str1)
    //fmt.Println(str2)
    
    output(str1)
    output(str2)
}
