package leetcode

import (
    "fmt"
)

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
/**
锯齿状数据,只能递减,每次减1。步数最少。
奇数位先减,操作的步数。
偶数位再减,操作的步数。
比较那个走得少。
 */
func MovesToMakeZigzag(nums []int) int {
    /**
    锯齿状数据,只能递减,每次减1。步数最少。
    奇数位先减,操作的步数。
    偶数位再减,操作的步数。
    比较那个走得少。
    */
    oddStep := 0
    eventStep := 0
    length := len(nums)
    // 奇位数先减 1,3,5,7
    for i:=0;i < length; i = i +2{
        tmp := nums[i]
        // 处在头部,只比较右边的数据
        if i == 0 {
            if nums[i + 1] > tmp {
                // A < B 高过，不用管
                continue
            }else {
                // A >= B,下拉
                oddStep = oddStep + (tmp - nums[i+1] + 1)
                continue
            }
        }
        // 处在尾部
        if i == length -1 {
            // B > C
            if nums[i-1] > tmp {
                continue
            }else {
                // B < C
                oddStep = oddStep + (tmp - nums[i-1] + 1)
                continue
            }
        }
        
        // 处在中间,比两边。低过最低的便可
        if i > 0 && i < length-1 {
            // 比较两边,
            if nums[i-1] > tmp && tmp < nums[i+1] {
                continue
            }
            if nums[i-1] > nums[i+1] {
                // A > C 低过右边便可
                if nums[i + 1] > tmp {
                    continue
                }else {
                    oddStep = oddStep + (tmp - nums[i+1] + 1)
                    continue
                }
            }else{
                // A < C 低过左边便可
                if nums[i - 1] > tmp {
                    continue
                }else {
                    oddStep = oddStep + (tmp - nums[i-1] + 1)
                    continue
                }
            }
        }
    }
    
    // 偶数位先减 2,4,6,8
    for i:=1;i < length;i= i + 2{
        tmp := nums[i]
        // 尾部
        if i == length -1 {
            // B > C
            if nums[i-1] > tmp {
                continue
            }else {
                // B < C
                eventStep = eventStep + (tmp - nums[i-1] + 1)
                continue
            }
        }
        // 中间
        if i > 0 && i < length-1 {
            // 比较两边,
            if nums[i-1] > tmp && tmp < nums[i+1] {
                continue
            }
            if nums[i-1] > nums[i+1] {
                // A > C 低过右边便可
                if nums[i + 1] > tmp {
                    continue
                }else {
                    eventStep = eventStep + (tmp - nums[i+1] + 1)
                    continue
                }
            }else{
                // A < C 低过左边便可
                if nums[i - 1] > tmp {
                    continue
                }else {
                    eventStep = eventStep + (tmp - nums[i-1] + 1)
                    continue
                }
            }
        }
    }
    
    if oddStep > eventStep {
        return eventStep
    }
    return oddStep
}
func ReverseList(head *ListNode) *ListNode {
   var tmp *ListNode = nil
   cur := head
   for cur != nil {
       tmpNode := &ListNode{
           Val: cur.Val,
       }
       tmpNode.Next = tmp
       tmp = tmpNode
       cur = cur.Next
   }
   return tmp
}
// 解法2
func ReverseList2(head *ListNode) *ListNode {
    var tmp *ListNode = nil
    cur := head
    for cur != nil {
        tmpNode := cur
        cur = cur.Next
        // 断开链表
        tmpNode.Next = nil
        // 重新连接
        tmpNode.Next = tmp
        tmp = tmpNode
    }
    return tmp
}
// 判断链表是否是回文数
func isPalindrome(head *ListNode) bool {
    // 复制链表
    cp := []int{}
    cur := head
    for cur != nil {
        cp = append(cp,cur.Val)
        cur = cur.Next
    }
    // 逆序比较
    cur = head
    index := len(cp) - 1
    index2 := 0
    for index != index2 && index2 < index {
        if cp[index] != cp[index2] {
            return false
        }
        index --
        index2 ++
    }
    //for cur != nil{
    //    if cur.Val != cp[index] {
    //        return false
    //    }
    //    index --
    //    cur = cur.Next
    //}
    return true
}
/**
二进制转10进制
 */
func getDecimalValue(head *ListNode) int {
    // 读出来
    //arr := []int{}
    cur := head
    length := 0
    for cur != nil {
        //arr = append(arr,cur.Val)
        cur = cur.Next
        length ++
    }
    // 按权重进行相加
    //length := len(arr)
    length --
    base := 1 << length
    result := 0
    for cur = head;cur!=nil;cur=cur.Next {
        result = result + cur.Val * base
        fmt.Println(result," ",base)
        base /= 2
    }
    return result
}
func getKthFromEnd(head *ListNode, k int) *ListNode {
    // 第一遍算长度
    length := 0
    cur := head
    for ;cur != nil;cur= cur.Next {
        length ++
    }
    // 算前进步数
    length = length - k
    cur = head
    for i:=1; i <=length; i++{
        cur = cur.Next
    }
    return cur
}
func getKthFromEnd2(head *ListNode, k int) *ListNode {
    fast,slow := head,head
    for i:=k;i>0;i--{
        fast = fast.Next
    }
    for fast != nil {
        slow = slow.Next
        fast = fast.Next
    }
    return slow
}
func removeDuplicateNodes(head *ListNode) *ListNode {
    // 向后读取,如果后面的数据和现在的一样,删除节点。不一样,指针前移
    if head == nil {
        return head
    }
    pre,cur := head,head.Next
    // 临时缓冲区
    buf := make(map[int]struct{})
    buf[pre.Val] = struct{}{}
    for cur != nil {
        if _,ok := buf[cur.Val];ok {
            // 相同,删除节点
            cur = cur.Next
            pre.Next = cur
        }else{
            // 不相同,继续
            buf[cur.Val] = struct{}{}
            pre = pre.Next
            cur = cur.Next
        }
    }
    return head
}
func removeDuplicateNodes2(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    
    pre,cur := head,head.Next
    for pre != nil {
        preTmp := pre
        for cur != nil {
            if cur.Val == pre.Val {
                cur = cur.Next
                preTmp.Next = cur
            }else{
                preTmp = preTmp.Next
                cur = cur.Next
            }
        }
        pre = pre.Next
        if pre == nil {
            break
        }
        cur = pre.Next
    }
    return head
}
func removeDuplicateNodes3(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    index := head
    for index != nil {
        cur := index
        for cur.Next != nil {
            if cur.Next.Val == index.Val {
                // 删除节点
                cur.Next = cur.Next.Next
            }else{
                cur = cur.Next
            }
        }
        index = index.Next
    }
    return head
}
func reversePrint(head *ListNode) []int {
    arr := []int{}
    for cur := head;cur != nil;cur = cur.Next{
        arr = append([]int{cur.Val},arr...)
    }
    return arr
}
func reversePrint2(head *ListNode) []int {
    length := 0
    for cur:=head;cur != nil;cur=cur.Next{
        length ++
    }
    arr := make([]int,length)
    for cur:=head;cur!=nil;cur=cur.Next{
        arr[length-1] = cur.Val
        length --
    }
    return arr
}
func deleteNode(node *ListNode) {
    // 快慢指针
    // 快指针比慢指针快1步，慢指针需要后退一步,再继续
    step := 2
    fast,slow := node,node
    for fast!=nil && fast.Next != nil {
        if step > 0 {
            // 停留一步
            step --
        }else{
            slow = slow.Next
        }
        
        fast = fast.Next.Next
    }
    // 删除中间节点
    slow.Next = slow.Next.Next
}
