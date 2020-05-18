package main

import (
    "fmt"
    "strconv"
    "time"
)

func main() {
    now := time.Now()
    fmt.Printf("now = %v now Type=%T",now,now)
    
    fmt.Println("year = ",now.Year())
    fmt.Println("mon = ",now.Month())
    fmt.Println("mon = ",int(now.Month()))
    fmt.Println("day = ",now.Day())
    fmt.Println("hour = ",now.Hour())
    fmt.Println("min = ",now.Minute())
    fmt.Println("second = ",now.Second())
    
    fmt.Printf("当前年月日:%d %d %d %d:%d:%d\n",
        now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
    
    datestr := fmt.Sprintf("当前年月日:%d %d %d %d:%d:%d",
        now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
    fmt.Println(datestr)
    fmt.Println(now.Format("2006-01-02 15:04:05"))
    now = time.Now()
    fmt.Println(now.Format("2006-01-02 03:04:05"))
    fmt.Printf("unix = %v unixnano = %v\n",now.Unix(),now.UnixNano())
    
    start := time.Now().Unix()
    strf := func() func()  {
        return func() {
            str := ""
            for i := 0; i < 100000; i++ {
                str += "Hello" + strconv.Itoa(i)
            }
        }
        }()
    strf()
    end := time.Now().Unix()
    fmt.Printf("执行%v秒\n",end -start)
    
    i := 0
    for {
        i++
        fmt.Println(i)
        time.Sleep(time.Millisecond*100)
        if i == 100 {
            break
        }
    }
    
}
