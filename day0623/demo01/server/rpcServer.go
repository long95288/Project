package main

import (
    "errors"
    "fmt"
    "net/http"
    "net/rpc"
)

type Args struct {
    A,B int
}
type Quotient struct{
    Quo, Rem int
}

// rpc 的结构对象
type Arith int
/*
两数相乘函数，被调用
*/
func (t *Arith) Multiply(args *Args, reply *int) error{
    *reply = args.A + args.B
    return nil
}
/*
两数相除函数，被调用
*/
func (t *Arith) Divide(args *Args, quo *Quotient) error{
    if args.B == 0{
        return errors.New("Divide by zero")
    }
    quo.Quo = args.A / args.B
    quo.Rem = args.A % args.B
    return nil
}

func main() {
    arith := new(Arith)
    rpc.Register(arith)
    rpc.HandleHTTP()
    
    err := http.ListenAndServe(":1234",nil)
    if err != nil {
        fmt.Println(err.Error())
    }
}
