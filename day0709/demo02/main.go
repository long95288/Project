package main

import (
    "bufio"
    "fmt"
    "os"
)
var m = make(map[string]int)
func k(list []string) string {
    return fmt.Sprintf("%q",list)
}
func Add(list []string) {
    m[k(list)]++
}
func Count(list []string) int {
    return m[k(list)]
}

func main() {
    //
    seen := make(map[string]bool)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        line := input.Text()
        if !seen[line] {
            // 设置已经出现过了
            seen[line] = true
            fmt.Println(line)
        }
    }
    if err := input.Err();err!=nil{
        fmt.Fprintf(os.Stderr,"dedup:%v\n",err)
        os.Exit(1)
    }
    
}
