package main

import (
    "fmt"
)

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
        // 创建新的节点,节点里面放的是相加之后的个位数
        // current.Next里面可能已经有值了，但是不管.直接覆盖
        current.Next = &ListNode{Val: sum%10}
        current = current.Next
        if p != nil{
            p = p.Next
        }
        if q != nil {
            q = q.Next
        }
        // 有进位的话会往前创建一个node,如果后面还有数据的话这个节点会被丢弃
        if carry >0  {
            current.Next = &ListNode{Val: carry}
        }
    }
    return flag.Next
}
func deleteDuplicates(head *ListNode) *ListNode {
    cur := head
    for cur != nil && cur.Next != nil {
        if cur.Val == cur.Next.Val {
            // 删除重复节点
            // 跳过下一个节点
            cur.Next = cur.Next.Next
        }else{
            // 不相同到下一个
            cur = cur.Next
        }
    }
    return head
}
// 递归的方法
func deleteDuplicates2(head *ListNode) *ListNode{
    // 递归出口
    if head == nil || head.Next == nil {
        return head
    }
    // 递归条件
    head.Next = deleteDuplicates2(head.Next)
    if head.Val == head.Next.Val {
        head = head.Next
    }
    return head
}
func main() {
    l1 := &ListNode{
        Val: 9,
        Next: &ListNode{
            Val: 9,
            Next: &ListNode{
                Val: 9,
            },
        },
    }
    l2 := &ListNode{
        Val: 9,
        Next: &ListNode{
            Val: 9,
            Next: &ListNode{
                Val: 9,
            },
        },
    }
    _ = l2
    //fmt.Println(AddTwoNumbers(l1,l2))
    fmt.Println(deleteDuplicates(l1))
    fmt.Println(deleteDuplicates2(l1))
}
