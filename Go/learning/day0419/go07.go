package main

func count() int {
    println("Count.")
    return 3
}

func main() {
    for i,c := 0,count(); i < c; i++ {
        println("a",i)
    }
    c := 0
    for c < count(){
        println("b",c)
        c ++
    }
    
    data :=[3]string{"a","b","c"}
    for i,s := range data{
        println(i,s)
    }
    for _,s := range data{
        println(s)
    }
    for range data{
        // 清空channel等操作
    }
}
