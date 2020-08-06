package leetcode

type ListNode struct {
    Val int
    Next *ListNode
}


/**

 */
func hasCycle(head *ListNode) bool {
    cur := head
    m := make(map[*ListNode]interface{})
    for cur != nil{
        if _,ok := m[cur];ok {
            return true
        }
        m[cur] = nil
        cur = cur.Next
    }
    return false
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    m := make(map[*ListNode]interface{})
    if headA == nil || headB == nil {
        return nil
    }
    cur := headA
    for cur != nil {
        m[cur] = true
        cur = cur.Next
    }
    // 遍历B的时候,判断
    cur = headB
    for cur != nil {
        if _,ok := m[cur];ok{
            return cur
        }
        cur = cur.Next
    }
    return nil
}
// 876. 链表的中间结点
func middleNode(head *ListNode) *ListNode {
    length := 0
    cur := head
    if head == nil {
        return head
    }
    for cur != nil {
        length ++
        cur = cur.Next
    }
    cur = head
    if length % 2 == 0{
        // 偶数 = (length+2)/2
        length = (length + 2)/2
    }else {
        // 奇数 = (length+1)/2
        length = (length + 1)/2
    }
    for ;length>1;length-- {
        cur = cur.Next
    }
    return cur
}