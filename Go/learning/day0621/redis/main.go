package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/go-redis/redis"
    "time"
)


// 全局redis连接变量
var rdb *redis.Client
var ctx = context.Background()
/**
初始化redis连接
 */
func initClient() error {
    rdb = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
        Password: "long?redis.2037",
        DB: 0,
    })
    
    _,err := rdb.Ping(ctx).Result()
    if err != nil {
        return err
    }
    return nil
}
func setDemo() {
    val := struct {
        Name string `json:"name"`
        Age int `json:"age" `
    }{"Tom",12}
    
    str,err := json.Marshal(val)
    if err != nil {
        return
    }
    fmt.Println("待序列化数据:",str)
    err = rdb.Set(ctx,"mykey",str ,0).Err()
    if err != nil {
        fmt.Println("设置错误",err)
    }
}
func setDemo2()  {
    // 设置3秒过期
    err := rdb.Set(ctx,"key2","hello world",3*time.Second).Err()
    if err != nil {
        fmt.Println("设置值错误,err = ",err)
    }
}
func getDemo2()  {
    val,err := rdb.Get(ctx,"key2").Result()
    if err != nil {
        fmt.Println("获得数据失败",err)
        return
    }
    fmt.Println("key2 = ",val)
}
func getDemo() {
    val,err := rdb.Get(ctx,"mykey").Result()
    if err != nil {
        fmt.Println("获得数据失败")
        return
    }
    fmt.Println("mykey = ",val)
}

func main() {
    // 初始化
    initClient()
    // 第一组测试
    setDemo()
    getDemo()
    // 第二组测试
    setDemo2()
    getDemo2()
    // 等待4秒后再次获得
    time.Sleep(4*time.Second)
    getDemo2()
}
