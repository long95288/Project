package main

import "fmt"

func main() {
    s := "雨痕笔记"
    
    for i := 0; i <len(s) ; i++ {
        fmt.Printf("%d: [%c]\n",i,s[i])
    }
    for i,c := range s{
        fmt.Printf("%d: [%c] \n",i,c)
        
    }
}
