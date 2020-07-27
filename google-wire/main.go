package main

import "fmt"

type Message struct {
    msg string
}
type Greeter struct {
    Message Message
}
type Event struct {
    Greeter Greeter
}

func NewMessage(msg string) Message {
    return Message{msg: msg}
}
// NewGreeter Greeter构造函数
func NewGreeter(m Message) Greeter {
    return Greeter{Message: m}
}
//
func NewEvent(g Greeter) Event{
    return Event{g}
}
func (e Event) Start(){
    msg := e.Greeter.Greet()
    fmt.Println(msg)
}
func (g Greeter) Greet() Message {
    return g.Message
}
//// 不使用wire
//func main() {
//    message := NewMessage("Hello World")
//    greeter := NewGreeter(message)
//    event := NewEvent(greeter)
//
//    event.Start()
//}
func main() {
    event := InitializeEvent("Hello_World")
    event.Start()
}
