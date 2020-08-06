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