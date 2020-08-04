package main

import "fmt"

type ListNode struct{
    Val int
    Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    
    p,q,flag := l1,l2,&ListNode{Val: 0}
    current := flag
    carry := 0
    for p != nil || q != nil {
        x := 0
        y := 0
        if p != nil {
            x = p.Val
        }
        if q != nil {
            y = q.Val
        }
        sum := carry + x + y
        carry = sum/10
        current.Next = &ListNode{Val: sum%10}
        current = current.Next
        if p != nil{
            p = p.Next
        }
        if q != nil {
            q = q.Next
        }
        if carry >0  {
            current.Next = &ListNode{Val: carry}
        }
    }
    return flag.Next
}

func main() {
    l1 := &ListNode{
        Val: 2,
        Next: &ListNode{
            Val: 3,
            Next: &ListNode{
                Val: 3,
            },
        },
    }
    l2 := &ListNode{
        Val: 5,
        Next: &ListNode{
            Val: 6,
            Next: &ListNode{
                Val: 4,
            },
        },
    }
    fmt.Println(AddTwoNumbers(l1,l2))
}
