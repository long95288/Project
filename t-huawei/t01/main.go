package main

import (
    "fmt"
)

// 寻找由两个素数相乘得到的数据
func isP(n int, pMap map[int]struct{}) bool {
    if _,ok := pMap[n];ok{
        return true
    }
    if n < 2 {
        return false
    }
    if n == 2 {
        return true
    }
    for j := 2; j < n; j ++ {
        if n % j == 0 {
            return false
        }
    }
    pMap[n] = struct{}{}
    return true
}
func main() {
    // 计算所有的素数
    var inputN = 0
    fmt.Scanf("%d", &inputN)
    //fmt.Printf("%d", inputN)
    pMap := make(map[int]struct{})
    max := inputN
    for i := 2; i <= max; i++{
        if isP(i, pMap) && max % i == 0 {
            sub :=  max / i
            if isP(sub, pMap) {
                //fmt.Println(pMap)
                if i < sub {
                    fmt.Printf("%d %d", i, sub)
                    return
                }else{
                    fmt.Printf("%d %d", sub, i)
                    return
                }
            }
        }
    }
    fmt.Println("-1 -1")
}
