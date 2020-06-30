package main

import (
    "fmt"
    "strconv"
    "sync"
    "time"
)

var inChan chan string
var outChan chan string
var lock sync.Mutex
var exit = 1
func product() {
    for i:=0;i<10;i++{
        s := strconv.FormatInt(int64(i),10)
        fmt.Println("写入阻塞:",time.Now().UnixNano())
        inChan <- s + "product"
        fmt.Println("写入成功:",time.Now().UnixNano())
        // 写入快慢
        // time.Sleep(2*time.Second)
    }
    close(inChan)
}
func consume()  {
    for {
        fmt.Println("读取阻塞:",time.Now().UnixNano())
        v,ok := <- inChan
        if !ok {
            fmt.Println("获得失败")
            break
        }
        fmt.Println("读取成功:",time.Now().UnixNano(),v)
        time.Sleep(time.Second*2)
    }
    close(outChan)
    lock.Lock()
    exit --
    lock.Unlock()
}
func consume2() {
    for v := range outChan {
        fmt.Println("consumer2 读出数据", v)
        // 读取慢
        time.Sleep(2*time.Second)
    }
    lock.Lock()
    exit --
    lock.Unlock()
}

func main() {
    inChan = make(chan string,2)
    outChan = make(chan string,2)
    go product()
    go consume()
    // go consume2()
    for {
        if exit == 0 {
            fmt.Println("所有goroutine完成运行")
            break
        }
    }
}
