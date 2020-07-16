package main

import (
    "fmt"
    "golang.org/x/net/html"
    "os"
)
// 遍历访问节点
func visit(links []string, n *html.Node) []string {
    if n.Type == html.ElementNode && n.Data == "a"{
        for _,a := range n.Attr{
            if a.Key == "href"{
                links = append(links,a.Val)
            }
        }
    }
    for c := n.FirstChild;c != nil;c = c.NextSibling{
        links = visit(links,c)
    }
    return links
}

func main() {
    file,err := os.Open("test.html")
    doc,err := html.Parse(file)
    if err != nil {
        fmt.Fprintf(os.Stdout,"findLinks1:%v\n",err)
        os.Exit(1)
    }
    for _,link := range visit(nil,doc){
        fmt.Println(link)
    }
}
