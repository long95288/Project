package main

import (
    "fmt"
    "strconv"
)

func isFlowers(n int) bool {
    ns := strconv.Itoa(n)
    ret := 0
    for _,i := range ns {
        ni,_ := strconv.Atoi(string(i))
        tmp := 1
        for j:=0; j < len(ns); j++{
            tmp *= ni
        }
        ret += tmp
    }
    return ret == n
}
func main() {
    // 1 数字位数
    // 2 第几个水仙花
    var n = 0
    var m = 0
    fmt.Scan(&n)
    fmt.Scan(&m)
    if n > 7 || n < 3 {
        fmt.Println("-1")
        return
    }
    //
    start := 1
    for i:=0; i < n - 1;i ++{
        start *= 10
    }
    res := []int{}
    end := 10 * start -1
    //fmt.Printf("%d %d\n", start, end)
    // 计算数据
    for i := start;i <= end;i++{
        if isFlowers(i) {
            res = append(res, i)
            if len(res) > m {
                fmt.Println(res[m])
                return
            }
        }
    }
    
    //
    if len(res) <= 0{
        fmt.Println("-1")
    }else if len(res) < m {
        fmt.Println(res[len(res) -1] * m)
    }else{
        fmt.Println(res[m])
    }
}
