package main

import (
    "fmt"
    "golang.org/x/net/html"
    "os"
)

func main() {
    file,err := os.Open("test.html")
    doc,err := html.Parse(file)
    if err != nil {
        fmt.Fprint(os.Stderr,"outline:%v\n",err)
        os.Exit(1)
    }
    outline(nil,doc)
}
func outline(stack []string, n *html.Node) {
    if n.Type == html.ElementNode {
        stack = append(stack,n.Data)
        fmt.Println(stack)
    }
    for c := n.FirstChild;c != nil;c = c.NextSibling{
        outline(stack,c)
    }
}