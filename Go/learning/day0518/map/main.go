package main

import "fmt"

func main() {
    var a map[string]string
    a = make(map[string]string,10)
    a["1"] = "v1"
    a["2"] = "v2"
    fmt.Println(a)
    
    m1 := make(map[string]string)
    m1["1"] = "c1"
    m1["2"] = "c2"
    m1["3"] = "c3"
    fmt.Println(m1)
    
    m2 := map[string]string {
        "1":"d1",
        "2":"d2",
    }
    fmt.Println(m2)
    
    m3 := map[string]map[string]string{
        "s1": {"name":"n1","class":"N1",
        },
        "s2":{
            "name":"n2","class":"N2",
        },
    }
    fmt.Println(m3)
    m3["s3"] = map[string]string{
        "name":"n3",
        "class":"N3",
    }
    fmt.Println(m3)
    delete(m3["s2"],"class")
    fmt.Println(m3)
    delete(m3,"s2")
    fmt.Println(m3)
    
    val,ok := m3["s1"]
    if ok{
        fmt.Println("value",val)
    }else{
        fmt.Println("no key")
    }
    for k,v := range m3{
        fmt.Printf("key:%v value:%v\n",k,v)
        for k,v := range v{
            fmt.Printf("key:%v value:%v\n",k,v)
        }
    }
    
    var monsters = []map[string]string{
        1:{
            "name":"牛魔王",
            "age":"500",
        },
        2:{
            "name":"玉兔精",
            "age":"400",
        },
    }
    fmt.Println(monsters)
    fmt.Printf("%T",monsters)
    monsters = append(monsters, map[string]string{
        "name":"新怪",
        "age":"400",
    })
    fmt.Println(monsters)
}
