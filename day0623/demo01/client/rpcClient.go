package main

import (
    "fmt"
    "log"
    "net/rpc"
)

type Args struct {
    A,B int
}
type Quotient struct {
    Quo,Rem int
}

func main() {
   
    client,err := rpc.DialHTTP("tcp","localhost:1234")
    if err != nil {
        log.Fatal("Dialing: ",err)
        return
    }
    args := Args{17,8}
    var reply int
    // 远程过程调用
    err = client.Call("Arith.Multiply",args,&reply)
    if err != nil {
        log.Fatal("arith error :",err)
        return
    }
    // 输出回调信息
    fmt.Printf("Arith: %d * % d = %d \n",args.A,args.B,reply)
    var quot Quotient
    err = client.Call("Arith.Divide",args,&quot)
    if err != nil {
        log.Fatal("arith error :",err)
        return
    }
    fmt.Printf("Arith: %d / %d = %d remainder %d \n",args.A,args.B,quot.Quo,quot.Rem)
    
}
