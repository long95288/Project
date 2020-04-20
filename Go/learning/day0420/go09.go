package main

func test() []func() {
    var s []func()
    
    for i:=0;i<2;i++{
        x := i
        s = append(s, func() {
            println(&x,x)
        })
    }
    return s
}
func main() {
    for _,f := range test(){
        f()
    }
}
