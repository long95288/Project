package leetcode

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "sort"
    "strconv"
    "strings"
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
func deleteDuplicates(head *ListNode) *ListNode {
    // map
    mp := make(map[int]int)
    for cur:=head;cur!=nil;cur=cur.Next{
        mp[cur.Val] += 1
    }
    dump := &ListNode{}
    dump.Next = head
    dup := false
    for cur := dump;cur!= nil&&cur.Next != nil;{
        // 判断重复 值大于2
        if v := mp[cur.Next.Val]; !dup && v > 1 {
            dup = true
        }
        // 删除重复的节点
        if dup {
            if v := mp[cur.Next.Val];v > 0{
                mp[cur.Next.Val] -= 1
                if v := mp[cur.Next.Val];v == 0 {
                    dup = false
                }
                cur.Next = cur.Next.Next
            }
        }else{
            cur = cur.Next
        }
    }
    return dump.Next
}
func deleteDuplicates2(head *ListNode) *ListNode {
    //if head == nil || head.Next == nil {
    //    return head
    //}
    //// 小于三个个元素的
    //if head.Next.Next == nil {
    //    if head.Val == head.Next.Val{
    //        return nil
    //    }else{
    //        return head
    //    }
    //}
    //// 三个元素以上的
    //// 新链表
    //newLink := &ListNode{}
    //// 新链表的尾巴指针
    //newLinkTail := newLink
    //// 前一个值,后一个值
    //beforeVal := head.Val
    //for cur := head.Next;cur != nil;{
    //    // 中间节点
    //    if cur.Next == nil {
    //        // 最后一个了,只要比较前面就行
    //        if beforeVal != cur.Val {
    //            // 不同,加入新链表
    //            newLinkTail.Next = cur
    //            break
    //        }else {
    //            //相同 掠过
    //            cur = cur.Next
    //        }
    //    }else{
    //        // 不是最后一个
    //        if beforeVal != cur.Val && cur.Val != cur.Next.Val {
    //            // 都不相同,加入节点
    //            newLinkTail.Next = cur
    //            // 更新值
    //            beforeVal = cur.Val
    //            cur = cur.Next
    //
    //            newLinkTail= newLinkTail.Next
    //            newLinkTail.Next = nil
    //        }else if beforeVal == cur.Val{
    //            cur = cur.Next
    //        }
    //    }
    //}
    //return newLink.Next
    return nil
}
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    
    // 先将待翻转的数据存下来,存完之后就翻转
    // 使用array存储
    // 哑节点
    if head.Next == nil {
        return head
    }
    dummy := &ListNode{}
    dummy.Next = head
    pre,cur := dummy,head
    arr := []*ListNode{}
    sub := n - m
    step := 2
    i := 0
    for {
        if step == 2 {
            // 分析阶段
            if i < m -1 {
                // 定位
                pre = cur
                cur = cur.Next
                i ++
            }else if i >= m -1 && i <= n -1{
                // 赋值
                arr = append(arr,cur)
                cur = cur.Next
                i ++
            }else{
                step --
                continue
            }
        }else if step == 1{
            // 反转阶段
            if sub >= 0 {
                tmp := arr[sub]
                tmp.Next = nil
                pre.Next = tmp
                pre = tmp
                sub --
            }else{
                // 翻转完成,链表连接
                pre.Next = cur
                step --
            }
        }else{
            break
        }
    }
    return dummy.Next
}
func reverseBetween2(head *ListNode, m int, n int) *ListNode {
    // 使用多指针，边走边翻转
    dummy := &ListNode{}
    dummy.Next = head
    // 新生成链表的尾巴
    dTail := dummy
    // 旧的数据
    var sHead *ListNode = nil
    var sTail *ListNode = nil
    i := 0
    cur := head
    for {
        if i < m -1 {
            dTail = cur
            cur = cur.Next
            i ++
        }else if i == m-1{
            // 切割点
            sTail = cur
            sHead = cur
            cur = cur.Next
            i ++
        }else if i > m -1 && i <= n -1 {
            // 翻转点
            tmp := cur
            cur = cur.Next
            tmp.Next = sHead
            sHead = tmp
            i ++
        }else if i == n {
            // 重新连接点
            dTail.Next = sHead
            sTail.Next = cur
            break
        }else{
            break
        }
    }
    
    return dummy.Next
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    stack1 := []int{}
    stack2 := []int{}
    // 数据入栈
    for cur:=l1;cur!= nil;cur=cur.Next{
        stack1 = append(stack1,cur.Val)
    }
    for cur:=l2;cur!=nil;cur=cur.Next{
        stack2 = append(stack2,cur.Val)
    }
    // 出栈相加
    dummy := &ListNode{}
    top1,top2 := len(stack1),len(stack2)
    carray := 0
    for top1 > 0 || top2 > 0{
        a,b := 0,0
        if top1 == 0 {
            a = 0
        }else{
            a = stack1[top1 - 1]
            top1 --
        }
        if top2 == 0 {
            b = 0
        }else{
            b = stack2[top2 -1]
            top2 --
        }
        sum := a + b + carray
        carray = sum / 10
        val := sum % 10
        node := &ListNode{
            Val:val,
            Next:nil,
        }
        node.Next = dummy.Next
        dummy.Next = node
    }
    if carray > 0 {
        cNode := &ListNode{
            Val:1,
            Next:nil,
        }
        cNode.Next = dummy.Next
        dummy.Next = cNode
    }
    return dummy.Next
}
func searchInsert(nums []int, target int) int {
    i := 0
    for ;i < len(nums);i++{
        if nums[i] == target {
            return i
        }
        if i + 1 < len(nums)  {
            if target > nums[i] && target < nums[i + 1]{
                return i + 1
            }
        }
    }
    return i
}
func merge(A []int, m int, B []int, n int)  {
    // 从后向前移动,每次寻找最大值填入A
    m,n = m-1,n-1
    index := m + n -1
    for m >=0 || n >= 0 {
        if m < 0 {
            // A已经选取结束了,把B放入便可
            A[index] = B[n]
            index --
            n --
            continue
        }
        if n < 0 {
            A[index] = A[m]
            index --
            m --
            continue
        }
        if A[m] >= B[n] {
            A[index] = A[m]
            m --
            index --
        }else {
            A[index] = B[n]
            n--
            index --
        }
    }
}
func generate(numRows int) [][]int {
   reNums := make([][]int,numRows)
   for i:=0;i<numRows;i++{
       if i == 0 {
           reNums[i] = []int{1}
       }else{
           subNum := make([]int,i+1)
           for j:=0;j < i + 1;j++ {
                pre1,pre2 := 0,0
                if j-1 <0 {
                    pre1 = 0
                }else{
                    pre1 = reNums[i-1][j-1]
                }
                
                if j >= i {
                    pre2 = 0
                }else{
                    pre2 = reNums[i-1][j]
                }
                subNum[j] = pre1 + pre2
           }
           reNums[i] = subNum
       }
   }
   return reNums
}
func reverse(nums []int, start int, end int) {
    for start < end{
        tmp := nums[start]
        nums[start] = nums[end]
        nums[end] = tmp
        start ++
        end --
    }
}
func rotate(nums []int, k int)  {
    // 先整体翻转,然后前k个数据翻转，然后再length - K个数据翻转
    length := len(nums)
    k = k % length
    // 1.全部翻转
    reverse(nums,0,length-1)
    // 2.翻转前k个
    reverse(nums,0,k-1)
    // 3.翻转后n-k个
    reverse(nums,k,length-1)
}

func factorial(n int64) int64{
    if n == 0 {
        return 1
    }
    return factorial(n-1)*n
}
func combination(n1, m1 int) int64 {
    n := int64(n1)
    m := int64(m1)
    if m == 0 {
        return 1
    }
    fn := factorial(n)
    fm := factorial(m)
    fnm := factorial(n-m)
    return (fn/fm)*(fn/fnm)
}
func getRow(rowIndex int) []int {
    // 只要计算一半便可
    half := rowIndex/2
    nums := make([]int,rowIndex + 1)
    for i:=0;i<=half;i++{
        nums[i] = int(combination(rowIndex,i))
    }
    for i:= half+1;i < rowIndex + 1;i++{
        nums[i] = nums[rowIndex - i]
    }
    return nums
}
// node插入到head的对应的位置
func insertNode(head *ListNode,node *ListNode){
    for cur := head;cur != nil;cur = cur.Next {
        if cur.Next != nil {
            if node.Val >= cur.Val && node.Val < cur.Next.Val {
                node.Next = cur.Next
                cur.Next = node
                break
            }
        }else{
            node.Next = nil
            cur.Next = node
            break
        }
    }
}
func partitionOld(head *ListNode, x int) *ListNode {
    // 构建两条链表,分别存放符合要求的。最后连接起来
    first,second := &ListNode{},&ListNode{}
    firstTail := first
    for cur:=head;cur!=nil;{
        if cur.Val <= x {
            // 搜索插入first表
            tmp := cur
            cur = cur.Next
            insertNode(first,tmp)
        }else{
            // 搜索插入second表
            tmp := cur
            cur = cur.Next
            insertNode(second,tmp)
        }
    }
    for cur:=first;cur!=nil;cur=cur.Next{
        firstTail = cur
    }
    // 重新连结
    firstTail.Next= second.Next
    return first.Next
}
func partition(head *ListNode, x int) *ListNode {
    first,second := &ListNode{},&ListNode{}
    firstTail,secondTail := first,second
    // 进行分割
    for cur := head;cur!= nil;{
        if cur.Val < x {
            // 添加到first链表
            tmp := cur
            cur = cur.Next
            tmp.Next = nil
            firstTail.Next = tmp
            firstTail = tmp
        } else {
            // 添加到second链表
            tmp := cur
            cur = cur.Next
            tmp.Next = nil
            secondTail.Next = tmp
            secondTail = tmp
        }
    }
    
    // 两个链表拼接
    firstTail.Next = second.Next
    return first.Next
}
func arrayPairSum(nums []int) int {
    // 排序
    sort.Ints(nums)
    // 奇数位相加
    sum := 0
    for i:=0;i<len(nums);i++{
        if i % 2 == 0 {
            sum += nums[i]
        }
    }
    return sum
}
func findPairs(nums []int, k int) int {
    if k < 0 {
        return 0
    }
    numsHas := make(map[int]bool)
    diffHas := make(map[int]bool)
    
    for _, num := range nums {
        if numsHas[num - k] {
            diffHas[num - k] = true
        }
        if numsHas[num + k] {
            diffHas[num] = true
        }
        numsHas[num] = true
    }
    return len(diffHas)
}
func matrixReshape(nums [][]int, r int, c int) [][]int {
    row := len(nums)
    column := len(nums[0])
    if r * c != row * column {
        return nums
    }
    list := []int{}
    for _,rowItem := range nums {
        for _,columnItem := range rowItem {
            list = append(list,columnItem)
        }
    }
    count := 0
    reNums := make([][]int,r)
    for i:=0;i<r;i++ {
        colNum := make([]int,c)
        for j:=0;j<c;j++{
            colNum[j] = list[count]
            count ++
        }
        reNums[i] = colNum
    }
    return reNums
}

func findUnsortedSubarray(nums []int) int {
    newNum := make([]int,len(nums))
    for i,num := range nums {
        newNum[i] = num
    }
    sort.Ints(newNum)
    fmt.Println(nums)
    fmt.Println(newNum)
    head := 0
    for i,num := range nums {
        if num != newNum[i] {
            head = i
            break
        }
    }
    tail := len(nums) -1
    for i:=tail;i>=0;i--{
        if nums[i] != newNum[i] {
            tail = i
            break
        }
    }
    return tail - head + 1
}
func canPlaceFlowers(flowerbed []int, n int) bool {
    f := append([]int{0},append(flowerbed,0)...)
    count := 0
    for i := 1;i <len(f)-1;i++ {
        if count == n {
            return true
        }
        if f[i-1] == 0 && f[i] == 0 && f[i+1] == 0 {
            count ++
            f[i] = 1
        }
    }
    return count == n
}
/**
给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reorder-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func reorderList(head *ListNode)  {
    // 分割 倒序 重连
    // 寻找中间节点
    slow,fast := head,head
    for ;fast != nil && fast.Next != nil;{
        fast = fast.Next.Next
        slow =slow.Next
    }
    // 从中间分割
    cur := slow.Next
    slow.Next = nil
    // 后面的链表倒置
    var newListHead *ListNode = nil
    for ; cur != nil;{
        tmp := cur
        cur = cur.Next
        tmp.Next = nil
        tmp.Next = newListHead
        newListHead = tmp
    }
    // 连接
    cur = head
    for cur2 := newListHead;cur2 != nil;{
        temp := cur2
        cur2 = cur2.Next
        temp.Next = cur.Next
        cur.Next = temp
        cur = cur.Next.Next
    }
}
func length(head *ListNode) int  {
    length :=0
    for cur:=head;cur!=nil;cur = cur.Next{
        length++
    }
    return length
}

func findMaxAverage(nums []int, k int) float64 {
    sum := 0
    max := math.MinInt64
    for i:=0;i<k;i++{
        sum += nums[i]
    }
    if sum > max{
        max = sum
    }
    for i := k;i < len(nums);i++{
        sum += nums[i]
        sum = sum - nums[i-k]
        if sum > max {
            max = sum
        }
    }
    return float64(max) / float64(k)
}
func isValid(s string) bool {
    stack := make([]int32,len(s))
    top := -1
    for _,c := range s {
       if top == -1 {
           top ++
           stack[top] = c
       }else{
            // '(' ')' '{' '}' '[' ']'
           if (stack[top] == '(' && c == ')') || (stack[top] == '{' && c == '}') || (stack[top] == '[' && c == ']') {
               top --
           }else{
               top ++
               stack[top] = c
           }
       }
    }
    return top < 0
}
type MinStack struct {
    Val []int
    Min []int
    TopVal int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    stack := MinStack{
        Val : []int{},
        Min : []int{},
        TopVal : 0,
    }
    stack.Val = append(stack.Val,-1)
    stack.Min = append(stack.Min,math.MinInt32)
    return stack
}


func (this *MinStack) Push(x int)  {
    if this.TopVal == 0{
        this.Min = append(this.Min,x)
    }else if x < this.Min[this.TopVal] {
        this.Min = append(this.Min,x)
    }else{
        this.Min = append(this.Min,math.MinInt32)
        this.Min[this.TopVal + 1] = this.Min[this.TopVal]
    }
    this.TopVal += 1
    this.Val = append(this.Val,x)
}

func (this *MinStack) Pop()  {
    this.TopVal -= 1
    this.Val = this.Val[:this.TopVal + 1]
    this.Min = this.Min[:this.TopVal + 1]
}


func (this *MinStack) Top() int {
    return this.Val[this.TopVal]
}


func (this *MinStack) GetMin() int {
    return this.Min[this.TopVal]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
func getNum(c int32) int {
    switch c {
    case 'I':
        return 1
    case 'V':
        return 5
    case 'X':
        return 10
    case 'L':
        return 50
    case 'C':
        return 100
    case 'D':
        return 500
    case 'M':
        return 1000
    default:
        return 0
    }
}
func romanToInt(s string) int {
    sum := 0
    pre := 0
    for i,v := range s{
        if i == 0{
            pre = getNum(v)
        }else{
            num := getNum(v)
            if pre < num {
                sum -= pre
            }else{
                sum += pre
            }
            pre = num
        }
    }
    sum += pre
    return sum
}
/**
图片平滑器
 */
func imageSmoother(M [][]int) [][]int {
    row := len(M)
    column := len(M[0])
    re := make([][]int,row)
    for i:=0;i<row;i++{
        re[i] = make([]int,column)
        for j := 0; j < column; j++ {
            count := 0
            sum := 0
            // 获得周围8格的数据
            for i1:=-1; i1 < 2; i1 ++{
                for j1:=-1;j1 < 2;j1++{
                    if i1 + i >= 0 && i1 + i < row  && j1 + j >= 0 && j1+ j < column{
                        sum += M[i1+i][j1+j]
                        count ++
                    }
                }
            }
            // 除以平均数
             re[i][j] = int(math.Floor(float64(sum) / float64(count)))
        }
    }
    return re
}
func checkPossibility(nums []int) bool {
    if len(nums) <=1 {
        return true
    }
    count := 0
    for i:=1;i<len(nums) && count < 2;i ++{
        if nums[i-1] <= nums[i]{
            continue
        }
        count ++
        if i-2 >=0 &&nums[i-2] > nums[i] {
            nums[i] = nums[i-1]
        }else {
            nums[i -1] = nums[i]
        }
    }
    return count <= 1
}

type MyStack struct {
    top int
    p1 []int
}


/** Initialize your data structure here. */
func Constructor1() MyStack {
    return MyStack{p1: []int{}}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
    this.p1 = append(this.p1,x)
    this.top = x
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
    top := this.top
    this.p1 = this.p1[:len(this.p1)-1]
    if len(this.p1) <= 0{
        this.top = -1
    }else {
        this.top = this.p1[len(this.p1)-1]
    }
    return top
}


/** Get the top element. */
func (this *MyStack) Top() int {
    return this.top
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
    return len(this.p1) <= 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    
    // 先计算num2每个元素更大的
    indexMap := make(map[int]int)
    for i:=0;i<len(nums2);i++{
        search := false
        for j:=i+1;j<len(nums2);j++{
            if nums2[j] > nums2[i] {
                indexMap[nums2[i]] = nums2[j]
                search = true
                break
            }
        }
        if !search{
            indexMap[nums2[i]] = -1
        }
    }
    ret := []int{}
    for _,v := range nums1{
       mv,_ := indexMap[v]
       ret = append(ret,mv)
    }
    return ret
}
/**
数字入栈
+ 两个数出栈相加并入栈
D 出栈*2入栈
C 最上面的出栈
 */
func calPoints(ops []string) int {
    stack := make([]int,len(ops))
    re := 0
    top := -1
    for _,v := range ops{
        //fmt.Println(v)
        switch v {
        case "C":
            re -= stack[top]
            top --
            
        case "D":
            tmp := 2 * stack[top]
            top ++
            stack[top] = tmp
            re += tmp
        case "+":
            tmp1 := stack[top]
            tmp2 := 0
            if top > 0 {
                top --
                tmp2 = stack[top]
                top ++
            }
            top ++
            stack[top] = tmp1 + tmp2
            re = re + tmp1 + tmp2
        default:
            d,_ := strconv.Atoi(v)
            top ++
            stack[top] = d
            re += d
        }
    }
    return re
}
func backspaceCompare(S string, T string) bool {
    stack1 := make([]rune,len(S))
    stack2 := make([]rune,len(T))
    top1,top2 := -1,-1
    for _,v := range S {
        if '#' == v{
            if top1 >= 0{
                top1 --
            }
        }else{
            top1 ++
            stack1[top1] = v
        }
    }
    for _,v := range T {
        if '#' == v {
            if top2 >= 0 {
                top2 --
            }
        }else{
            top2 ++
            stack2[top2] = v
        }
    }
    if top1 == top2 {
        for ;top1 >= 0;top1--{
            if stack1[top1] != stack2[top1] {
                return false
            }
        }
        return true
    }
    return false
}
func removeOuterParentheses(S string) string {
    stack := make([]rune,len(S))
    s := make([]rune,len(S))
    index := -1
    top := -1
    //re := ""
    for _,v := range S{
        if v == '('{
            if  top >= 0 {
                // 不是最外层括号,需要打印出
                //re += "("
                index ++
                s[index] = '('
                //是最外层括号,只要进栈
            }
            top ++
            stack[top] = v
        }else {
            if top > 0 {
                // 打印非最外层
                //re += ")"
                index ++
                s[index] = ')'
            }
            top --
        }
    }
    return string(s[:index+1])
}
func removeDuplicates(S string) string {
    // 入栈,相同出栈
    stack := make([]rune,len(S))
    top := -1
    for _,v := range S{
        if top < 0{
            top++
            stack[top] = v
        }else{
            if stack[top] == v {
                top --
            }else{
                top++
                stack[top] = v
            }
        }
    }
    return string(stack[:top+1])
}
func addBinary(a string, b string) string {
    aBit := make([]int,len(a))
    bBit := make([]int,len(b))
    stack := make([]int,len(a)+len(b))
    top := -1
    for i, v := range a {
       d,_ := strconv.Atoi(string(v))
       aBit[i] = d
    }
    for i, v := range b {
        d,_ := strconv.Atoi(string(v))
        bBit[i] = d
    }
    i,j := len(a)-1,len(b)-1
    
    tmpa := aBit[0]
    tmpb := bBit[0]
    carry := 0
    for i >= 0 || j >=0 {
        if i < 0 {
            tmpa = 0
        }else{
            tmpa = aBit[i]
        }
        if j < 0 {
            tmpb = 0
        }else {
            tmpb = bBit[j]
        }
        sum := tmpa + tmpb + carry
        carry = sum / 2
        top ++
        stack[top] = sum % 2
        i --
        j --
    }
    if carry > 0 {
        top ++
        stack[top] = carry
    }
    ret := make([]rune,top+1)
    stack_size := top
    for top >= 0 {
        ret[stack_size - top] = rune(48 + stack[top])
        top --
    }
    return string(ret[:stack_size + 1])
}

func isPalindrome2(s string) bool {
    if "" == s {
        return true
    }
    stack := make([]rune,len(s))
    top := -1
    
    tmp_s := make([]rune,len(s))
    index_tmp := -1
    for _,v := range s {
        if (v >= '0' && v <= '9') || (v >= 'a' && v <='z') || (v >='A' && v <= 'Z') {
            top ++
            stack[top] = v
            
            index_tmp ++
            tmp_s[index_tmp] = v
        }
    }
    
    for i:=0;i<= top;i++{
        v := tmp_s[i]
        tmp := stack[top - i]
        if (v >= 'a' && v <='z') || (v >='A' && v <= 'Z') {
            sub := math.Abs(float64(tmp - v))
            if sub == 0 || sub == 32{
                continue
            }else {
                return false
            }
        }else{
            if v == tmp {
                continue
            }
            return false
        }
    }
    return true
}

func longestCommonPrefix(strs []string) string {
    prefix := ""
    if len(strs) == 1 {
        return strs[0]
    }
    if len(strs) == 2 {
        first := strs[0]
        second := strs[1][:]
        for i,v := range first {
            if i < len(second){
                if uint8(v) == second[i] {
                    prefix += string(v)
                }else{
                    break
                }
            }else{
                break
            }
        }
        return prefix
    }
    prefix = longestCommonPrefix(strs[:2])
    return longestCommonPrefix(append([]string{prefix},strs[2:]...))
}

func strStr(haystack string, needle string) int {
    if "" == needle {
        return 0
    }
    // 思路,for循环,比较needle。完全匹配才可以
    needle_ := needle[:]
    haystack_ := haystack[:]
    for i:=0;i<len(haystack);i++{
        v := haystack_[i]
        if v == needle_[0]{
            // 开始进行匹配
            j := i
            index_needle := 0
            for index_needle < len(needle) && j < len(haystack) {
                if index_needle == len(needle) -1 && needle_[index_needle] == haystack_[j] {
                    return i
                }
                if haystack_[j] != needle_[index_needle]  {
                    break
                }else{
                    j ++
                    index_needle ++
                }
            }
        }
    }
    return -1
}
func Max(nums []int) int {
    max := math.MinInt32
    for _,v := range nums {
        if v > max {
            max = v
        }
    }
    return max
}
func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) <= 0 {
        return []int{0}
    }
    if len(nums) <= k {
        return []int{Max(nums)}
    }
    slideWindows := make([]int,k)
    baseMax := math.MinInt32
    ret := make([]int, len(nums) - k + 1)
    retIndex := 0
    for i:=0;i<k && i <len(nums);i++{
        if nums[i] > baseMax {
            baseMax = nums[i]
        }
        slideWindows[i] = nums[i]
    }
    ret[retIndex] = baseMax
    for i:=k;i<len(nums);i++{
        slideWindows = append(slideWindows[1:],nums[i])
        baseMax = Max(slideWindows)
        retIndex ++
        ret[retIndex] = baseMax
    }
    return ret
}
func isHappy(n int) bool {
    strn := strconv.FormatInt(int64(n),10)
    result := 0
    for _,v := range strn{
        tmp,_ := strconv.Atoi(string(v))
        result += tmp * tmp
    }
    if result == 1 {
        return true
    }else if result == 4{
        return false
    }
    return isHappy(result)
}
func countAndSay(n int) string {
    if n <= 1 {
        return "1"
    }
    result := ""
    tmpStr := countAndSay(n - 1)
    tmp := tmpStr[:]
    count := 0
    tmpChar := tmp[0]
    // 数数
    for i := 0; i < len(tmp); i++{
        if tmpChar == tmp[i] {
            count ++
        }else{
            // 出现不同
            result += strconv.Itoa(count) + string(tmpChar)
            tmpChar = tmp[i]
            count = 1
        }
    }
    if count > 0 {
        result += strconv.Itoa(count) + string(tmpChar)
    }
    return result
}
func mySqrt(x int) int {
    if x == 1 {
        return 1
    }
    half := x/2
    start := 1
    end := half
    mid :=  (start + end) / 2
    for i:=1;i <= 1000;i++{
        if start * start == x {
            return start
        }
        if end * end == x {
            return end
        }
        if start*start < x && end * end > x && end - start == 1 {
            return start
        }
        mid = (start + end) / 2
        if mid * mid < x {
            start = mid
        }else {
            end = mid
        }
    }
    return half
}
func isPrimes(n int) bool{
    for i:=2; i<n; i++{
        if (float64(n)/float64(i) - float64(n/i)) == 0 {
            return false
        }
    }
    return true
}
func countPrimes(n int) int {
    count := 0
    if n < 2 {
        return 0
    }
    if n == 2 {
        return 1
    }
    count ++
    for i := 3;i<n;i++{
        if isPrimes(i) {
            count++
        }
    }
    return count
}
func countPrimes2(n int) int {
    count := 0
    if n <= 2 {
        return 0
    }
    count ++
    arr := make([]bool,n)
    for i := 3;i < n;i += 2{
        if !arr[i] {
            for j:=3;i*j <n;j+= 2{
                arr[i*j] = true
            }
            count ++
        }
    }
    return count
}
func sumArray(arr []int) int {
    result := 0
    for _,v := range arr{
        result += v
    }
    return result
}
func sumOddLengthSubarrays(arr []int) int {
    result := 0
    if len(arr) <= 0 {
        return 0
    }
    // 设置滑动窗口
    slide_window := make([]int,1)
    slide_index := 0
    slide_size := 1
    for ;slide_size <= len(arr);slide_size += 2{
        slide_index = 0
        slide_window = make([]int,slide_size)
        for _,v := range arr {
            // 先填充滑动窗口
            if slide_index <= slide_size -1 {
                slide_window[slide_index] = v
                slide_index ++
            }
            if slide_index == slide_size {
                // 填满之后计算结果
                result += sumArray(slide_window)
                slide_window = append(slide_window[1:],0)
                slide_index -= 1
            }
        }
    }
    return result
}

func minSubarray(nums []int, p int) int {
    // 一个数x能够被p除尽,x % p == 0
    // 一个数x不能被p除尽,x % p == y 可以知道(x - y) % p = 0,其中x = n * p
    // 由已知可以求得y,然后在数组中寻找y,或者加起来得数据除上p，余数等于y的数组
    sum := 0
    for _,v := range nums{
        sum += v
    }
    y := sum % p
    // 在数组中寻求y值
    if y == 0 {
        return 0
    }
    for i := 0; i<len(nums); i++{
        nums[i] = nums[i] % p
    }
    for i := 1; i < len(nums);i++ {
        // i滑动窗口的大小
        // [1] [2] [3] [4] ...
        // [1,2] [2,3]
        // [1,2,3]
        init := false // 是否初始化滑动窗口
        windows := 0
        for j := i;j < len(nums);j ++{
            // 根据滑动窗口的大小,计算滑动窗口中的值
            if !init {
                for k := 0; k < i; k ++{
                    windows += nums[k]
                }
                if windows % p == y {
                    return i
                }
                init = true
            }
            if init {
                // 初始化:[1]
                // 右移:[1,2] [2]
                if j < len(nums){
                    windows -= nums[j - i]
                    windows += nums[j]
                }
                if windows % p == y {
                    return i
                }
            }
        }
        
    }
    return -1
}
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func Root2Leaves(root *TreeNode) int {
    if root == nil || (root.Left == nil && root.Right == nil){
        return 0
    }
    left := Root2Leaves(root.Left)
    leftEqual := false
    right := Root2Leaves(root.Right)
    rightEqual := false
    
    if root.Left != nil && root.Left.Val == root.Val {
        left += 1
        leftEqual = true
    }
    if root.Right != nil && root.Right.Val == root.Val {
        right += 1
        rightEqual = true
    }
    if !leftEqual && !rightEqual {
        // 都不相同,返回0
        return 0
    }else if(leftEqual && !rightEqual){
        return left
    }else if (!leftEqual && rightEqual){
        return right
    }
    return int(math.Max(float64(left),float64(right)))
}

func Leave2Leave(root *TreeNode) int {
    // 叶子到叶子节点
    if root == nil || (root.Left == nil && root.Right == nil){
        return 0
    }
    r := 0
    // 两边的子树值和本身相等才算
    if root.Left != nil && root.Right != nil && root.Left.Val == root.Val && root.Right.Val == root.Val {
        // 取得两个子树的直达长度
        left := Root2Leaves(root.Left)
        right := Root2Leaves(root.Right)
        r += left + right + 2
    }else {
        r = int(math.Max(float64(Leave2Leave(root.Left)),float64(Leave2Leave(root.Right))))
    }
    return r
}
func longestUnivaluePath(root *TreeNode) int {
    // 两种情况,1.从根节点叶子的路径长度 2.叶子节点到叶子节点的路径长度
    root2LeaveLength := Root2Leaves(root)
    leave2LeaveLength := Leave2Leave(root)
    
    if root2LeaveLength > leave2LeaveLength{
        return root2LeaveLength
    }
    return leave2LeaveLength
}


func fib(n int) int {
    if n <= 1{
        return n
    }
    nums := make([]int64,n+1)
    for i:=0;i <= n;i++ {
        if i == 0 {
            nums[i] = 0
        }else if i == 1 {
            nums[i] = 1
        }else{
            nums[i] = nums[i-1]%1000000007 + nums[i-2]%1000000007
        }
    }
    
    return int(nums[n]%1000000007)
}
func move(n int, A *[]int, B *[]int, C *[]int){
    if n == 0 {
        return
    }
    if n == 1 {
        // 一个盘子,只需要从A->C
        *C = append(*C,(*A)[len(*A)-1])
        (*A) = (*A)[:len(*A)-1]
    }else{
        // 2以上个盘子 A->B A->C B->C
        move(n-1,A,C,B)
        move(1,A,B,C)
        move(n-1,B,A,C)
    }
}
func hanota(A []int, B []int, C []int) []int {
    if A == nil{
        return nil
    }
    move(len(A),&A,&B,&C)
    return C
}

func missingNumber(nums []int) int {
    first,last,mid := 0,len(nums)-1,(len(nums)/2)
    for first <= last {
        mid = (first + last)/2
        if mid == nums[mid] {
            first = mid + 1
        }else{
            last = mid -1
        }
    }
    return first
}
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }else if p != nil && q != nil {
        if p.Val == q.Val {
            return isSameTree(p.Left,q.Left) && isSameTree(p.Right,q.Right)
        }
    }
    return false
}
func isSymmetricTrees(p *TreeNode,q *TreeNode) bool{
    if p ==nil && q == nil {
        return true;
    }else if p!= nil && q != nil {
        if  p.Val == q.Val {
            return isSymmetricTrees(p.Left,q.Right) && isSymmetricTrees(p.Right,q.Left)
        }
    }
    return false
}
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return isSymmetricTrees(root.Left,root.Right)
}
func max(a,b int) int{
    if a > b {
        return a
    }
    return b
}
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + max(maxDepth(root.Left),maxDepth(root.Right))
}


func levelOrderBottom(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    ret := [][]int{}
    left := levelOrderBottom(root.Left)
    right := levelOrderBottom(root.Right)
    i,j := 0,0
    for i < len(left) || j < len(right) {
        if (len(left)-1 - i) == (len(right)-1 - j) {
            ret = append(ret,append(left[i],right[j]...))
            i ++
            j ++
        }else if (len(left)-1 - i) > (len(right)-1 - j){
            ret = append(ret,left[i])
            i ++
        }else {
            ret = append(ret,right[j])
            j ++
        }
        
    }
    ret = append(ret,[]int{root.Val})
    return ret
}
func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
    if root.Left == nil && root.Right == nil && root.Val == sum {
        return true
    }
    target := sum - root.Val
    return hasPathSum(root.Left,target) || hasPathSum(root.Right,target)
}

func pathSum(root *TreeNode, sum int) [][]int {
    if root == nil {
        return [][]int{}
    }
    if root.Left == nil && root.Right == nil && root.Val == sum {
        return [][]int{{root.Val}}
    }
    target := sum - root.Val
    left := pathSum(root.Left,target)
    right := pathSum(root.Right,target)
    ret := [][]int{}
    for i:= 0;i<len(left);i++{
        ret = append(ret,append([]int{root.Val},left[i]...))
    }
    for i:= 0;i<len(right);i++{
        ret = append(ret,append([]int{root.Val},right[i]...))
    }
    return ret
}

func sortedArrayToBST(nums []int) *TreeNode {
    return dfs(nums,0, len(nums) -1)
}
func dfs(nums []int, first int, last int) *TreeNode{
    if first > last {
        return nil
    }
    mid := (first +last) / 2
    root := TreeNode{Val: nums[mid]}
    root.Left = dfs(nums, first, mid -1)
    root.Right = dfs(nums,mid +1 ,last)
    return &root
}

func min(x,y int) int {
    if x > y{
        return y
    }
    return x
}
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    if nil == root.Left && nil == root.Right {
        return 1
    }
    minDep := math.MaxInt32
    if root.Left != nil {
        minDep = min(minDepth(root.Left), minDep)
    }
    if root.Right != nil{
        minDep =  min(minDepth(root.Right),minDep)
    }
    return minDep + 1
}

func uncommonFromSentences(A string, B string) []string {
    m1 := make(map[string]int)
    m2 := make(map[string]int)
    for _,v := range strings.Split(A, " ") {
        m1[v] += 1
    }
    for _,v := range strings.Split(B, " ") {
        m2[v] += 1
    }
    ret := []string{}
    for k,v := range m1 {
        if v == 1 {
            if _, ok := m2[k]; !ok{
                ret = append(ret,k)
            }
        }
    }
    for k,v := range m2 {
        if v == 1 {
            if _, ok := m1[k]; !ok{
                ret = append(ret,k)
            }
        }
    }
    return ret
}
func mapping(s,t string) bool  {
    n := len(s)
    map1 := make(map[uint8]uint8)
    for i:=0 ;i< n;i ++{
        s1 := s[i]
        t1 := t[i]
        if _, ok := map1[s1];ok {
            if v,_ := map1[s1];v != t1 {
                return false
            }
        } else {
            map1[s1] = t1
        }
    }
    return true
}
func isIsomorphic(s string, t string) bool {
   return mapping(s, t) && mapping(t, s)
}
func canConstruct(ransomNote string, magazine string) bool {
    // 两个map,记录字母个数
    rMap := make(map[rune]int)
    mMap := make(map[rune]int)
    for _, v := range ransomNote {
        rMap[v] += 1
    }
    for _, v := range magazine {
        mMap[v] += 1
    }
    for k,v := range rMap{
        if mV,ok := mMap[k];ok && mV >= v{
        }else{
            return false
        }
    }
    return true
}
func firstUniqChar(s string) int {
    numMap:= make(map[rune]int)
    for _, v := range s {
        numMap[v] += 1
    }
    for i,v := range s {
       if n,_ := numMap[v]; n == 1 {
           return i
       }
    }
    return -1
}

func reverseString(s []byte)  {
    // 思路,一个指向末尾的指针，一个指向首个地址的指针。相互交换数据
    firstIndex,lastIndex := 0, len(s) -1
    for firstIndex < lastIndex {
        tmp := s[firstIndex]
        s[firstIndex] = s[lastIndex]
        s[lastIndex] = tmp
        firstIndex ++
        lastIndex --
    }
}
func isVowel(c rune) bool  {
    if c == 'a' || c == 'A' || c == 'e' || c == 'E'  || c == 'i' || c == 'I'|| c == 'o' || c == 'O' || c == 'u' || c=='U'{
        return true
    }
    return false
}
func reverseVowels(s string) string {
    sb := []rune(s)
    firstIndex,lastIndex := 0, len(sb) -1
    for firstIndex < lastIndex {
        if isVowel(sb[firstIndex]) && isVowel(sb[lastIndex]){
            tmp := sb[firstIndex]
            sb[firstIndex] = sb[lastIndex]
            sb[lastIndex] = tmp
            firstIndex ++
            lastIndex --
        }
        if !isVowel(sb[firstIndex]) {
            firstIndex ++
        }
        if !isVowel(sb[lastIndex]) {
            lastIndex --
        }
    }
    return string(sb)
}
func addStrings(num1 string, num2 string) string {
    // 从个位数开始相加,保存进位,一个个相加
    num1B := []rune(num1)
    num2B := []rune(num2)
    carry := 0
    l1,l2:= len(num1)-1,len(num2)-1
    resultB := []rune{}
    for l1 >= 0 || l2 >= 0 {
        tmp1 := 0
        tmp2 := 0
        if l1 >= 0 {
            tmp1,_ = strconv.Atoi(string(num1B[l1]))
            l1 --
        }
        if l2 >= 0 {
            tmp2, _ = strconv.Atoi(string(num2B[l2]))
            l2 --
        }
        result := (tmp1 + tmp2 + carry) % 10
        carry = (tmp1 + tmp2 + carry) / 10
        resultB = append([]rune{rune(result + '0')}, resultB...)
    }
    if carry > 0 {
        resultB = append([]rune{rune(carry + '0')}, resultB...)
    }
    return string(resultB)
}
func countSegments(s string) int {
    num := 0
    if s == "" {
        return num
    }
    flag := false
    for _,v := range s {
        if v != ' ' && !flag {
            flag = true
            num ++
        }
        if v == ' ' {
            flag = false
        }
    }
    return num
}
func repeatedSubstringPattern(s string) bool {
    if "" == s {
        return true
    }
    return strings.Index(string((s+s)[1:]), s) < len(s) -1
}
func detectCapitalUse(word string) bool {
    if len(word) <= 1 {
        return true
    }
    capitalNum := 0
    index := -1
    for i,v := range word{
        if v < 'a' {
            capitalNum ++
        }
        // 大写,只有第一个或者全为大写的时候才会变更index
        if v < 'a' && (i == 0 || i == index){
            index ++
        }
    }
    // 全大写,全小写,第一个为大写
    return capitalNum == len(word) || capitalNum == 0 || (capitalNum == 1 && index == 0)
}
func findLUSlength(a string, b string) int {
    if a == b {
        return -1
    }
    return max(len(a), len(b))
}


func reverseStr(s string, k int) string {
    // 按k分割, 0k 2k 4k 6k 8k 10k 反转
    ret := []byte(s)
    // ab cd ef g
    // 0  2  4  6
    for i :=0; i< len(ret) -1;i += 2*k {
        l,r := i, min(i + k -1, len(ret) -1)
        for l < r {
            tmp := ret[l]
            ret[l] = ret[r]
            ret[r] = tmp
            l ++
            r --
        }
    }
    return string(ret)
}

func checkRecord(s string) bool {
    absentNum := 0
    for _,v := range s {
        if v == 'A' {
            if absentNum >= 1 {
                return false
            }
            absentNum ++
        }
    }
    return !strings.Contains(s,"LLL")
}

func maxLengthBetweenEqualCharacters(s string) int {
    // 两个循环,计算相同元素之间的长度
    tmp := []rune(s)
    maxLength := -1
    for i:= 0;i < len(tmp); i++{
        for j := i + 1;j < len(tmp);j ++{
            if tmp[i] == tmp[j] {
                if j - i - 1 > maxLength {
                    maxLength = j - i - 1
                }
            }
        }
    }
    return maxLength
}
// 累加偶数位操作
func aF(s string, a int) string {
    tmp2 := []byte(s)
    for i := 1 ; i < len(s); i += 2 {
        tmp2N,_ := strconv.Atoi(string(tmp2[i]))
        tmp2[i] = byte('0' + (tmp2N + a)%10)
    }
    return string(tmp2)
}
// 累加奇数位
func aF2(s string, a int) string  {
    tmp2 := []byte(s)
    for i := 0 ; i < len(s); i += 2 {
        tmp2N,_ := strconv.Atoi(string(tmp2[i]))
        tmp2[i] = byte('0' + (tmp2N + a)%10)
    }
    return string(tmp2)
}
// 轮转操作
func bF(s string, b int) string {
    tmp2 := []byte(s)
    tmp2B := []byte{}
    for _,v := range tmp2[len(tmp2) - b:] {
        tmp2B = append(tmp2B, v)
    }
    tmp2B = append(tmp2B, tmp2[:len(tmp2) - b]...)
    return string(tmp2B)
}
func findLexSmallestString(s string, a int, b int) string {
    minNumstr := s
    // 一个数 经过10次累加之后,回到初始状态
    // 一个数 经过len(s)次轮转之后,回到初始状态
    // 对于偶数轮转来说: 累加只能改变奇数位,偶数位不会被累加。所有的枚举结果: len(s) * 10
    // 对于奇数轮转来说: 虽然累加只能改变奇数位,但是经过轮转之后,奇数位变成偶数位,偶数位变成奇数位。这时候,奇数位，偶数位都应该累加。
    
    tmp := s
    for i := 0;i < len(s);i ++{
        // 轮转
        tmp = bF(tmp, b)
        if tmp < minNumstr {
            minNumstr = tmp
        }
        //fmt.Println(tmp)
        //// 对轮转之后的结果进行累加
        addStr := tmp
        for j := 0;j < 10;j ++{
           // 累加完偶数位
           addStr= aF(addStr, a)
           if addStr < minNumstr {
               minNumstr = addStr
           }
           // 如果是奇数,把奇数位也累加
           if b % 2 != 0 {
               for k := 0;k < 10;k ++ {
                   addStr = aF2(addStr, a)
                   if addStr < minNumstr{
                       minNumstr = addStr
                   }
               }
           }
        }
       
    }
    return minNumstr
}
func reverseWords(s string) string {
    ret := []rune{}
    arr := strings.Split(s, " ")
    for i,v := range arr{
        l,r := 0, len(v) -1
        tmpV := []rune(v)
        for l < r {
            tmp := tmpV[l]
            tmpV[l] = tmpV[r]
            tmpV[r] = tmp
            l ++
            r --
        }
        if i != len(arr) -1 {
            tmpV = append(tmpV, ' ')
        }
        ret = append(ret, tmpV...)
    }
    return string(ret)
}

func tree2str(t *TreeNode) string {
    if nil == t {
        return ""
    }
    if nil != t.Left && nil != t.Right {
        return strconv.Itoa(t.Val) +"(" + tree2str(t.Left)+")" +"("+ tree2str(t.Right)+")"
    }else if nil != t.Left && nil == t.Right {
        return strconv.Itoa(t.Val) +"("+ tree2str(t.Left) + ")"
    }else if nil == t.Left && nil != t.Right {
        return strconv.Itoa(t.Val) + "()"+ "(" + tree2str(t.Right) +")"
    }else{
        return strconv.Itoa(t.Val)
    }
}
func isPalindrome3(s string) bool{
    right := len(s) -1
    sb := []rune(s)
    for i := 0;i < right; {
        if sb[i] == sb[right] {
            right --
            i ++
        }else{
            return false
        }
    }
    return true
}
func validPalindrome(s string) bool {
    // 双指针,夹逼。
    // 出现不一致,尝试删除左边或者右边看是否相等
    right := len(s) -1
    sb := []rune(s)
    for i := 0; i < right ;{
        if sb[i] == sb[right] {
            right --
            i ++
            continue
        } else {
            // 去除左边数据
            return isPalindrome3(string(sb[i+1: right + 1])) || isPalindrome3(string(sb[i: right]))
        }
    }
    return true
}
func isFlipedString(s1 string, s2 string) bool {
    if s1 == s2 {
        return true
    }
    s1bytes := []rune(s1)
    for i := 0; i < len(s1bytes); i++{
        // 轮转
        tmp := s1bytes[:i+1]
        ret := s1bytes[i+1:]
        for j:=0;j<len(tmp);j++ {
            ret = append(ret, tmp[j])
        }
        if string(ret) == s2 {
            return true
        }
    }
    return false
    
}
func intersection(nums1 []int, nums2 []int) []int {
    map1 := make(map[int]int)
    map2 := make(map[int]int)
    for _,v := range nums1{
        map1[v] += 1
    }
    for _,v := range nums2 {
        map2[v] += 1
    }
    ret := []int{}
    for k,_ := range map1{
        if _,ok := map2[k];ok {
            ret = append(ret, k)
        }
    }
    return ret
}
func intersect(nums1 []int, nums2 []int) []int {
    map1 := make(map[int]int)
    map2 := make(map[int]int)
    for _,v := range nums1{
        map1[v] += 1
    }
    for _,v := range nums2 {
        map2[v] += 1
    }
    ret := []int{}
    for k,v1 := range map1{
        if v2,ok := map2[k];ok {
            v := v1
            if v1 > v2 {
                v = v2
            }
            for i:=v; v >0; i--{
                ret = append(ret, k)
            }
        }
    }
    return ret
}
func findTheDifference(s string, t string) byte {
    map1 := make(map[byte]int)
    map2 := make(map[byte]int)
    for _,v := range s {
        map1[byte(v)] += 1
    }
    for _,v := range t {
        map2[byte(v)] += 1
    }
    for k,v := range map2{
        if v2,ok := map1[k];!ok || v != v2 {
            return k
        }
    }
    return 0
}
func CountCharNum()  {
    in := bufio.NewReader(os.Stdin)
    input,_,_ := in.ReadLine()
    target,_,_ := in.ReadLine()
    
    map1 := make(map[byte]int)
    for _,v := range input {
        map1[v] += 1
    }
    v,_ := map1[target[0]]
    fmt.Println(v)
}

func wordPattern(pattern string, s string) bool {
    // 分割字符串,先建立起pattern和s的映射,然后再进行顺序判断
    patternMap := make(map[rune]string)
    sMap := make(map[string]rune)
    
    arr := strings.Split(s, " ")
    if len(arr) != len(pattern){
        return false
    }
    for i, v := range pattern {
        if _,ok := patternMap[v];!ok {
            patternMap[v] = arr[i]
            sMap[arr[i]] = v
        }
    }
    for i, v := range pattern {
        if v2, ok := sMap[arr[i]]; !ok || v2 != v {
            return false
        }
    }
    return true
}
func longestPalindrome(s string) int {
  mapP := make(map[rune]int)
  for _,v := range s {
      mapP[v] += 1
  }
  ret := 0
  for _,v := range mapP{
        ret += (v/2)*2
  }
  if ret == len(s) {
      return ret
  }
  return ret + 1
}
func islandPerimeter(grid [][]int) int {
    // 如果是陆地，计算相邻的四个格子，如果相邻的格子是0或者是边界,添加一条边
    m := len(grid)
    n := len(grid[0])
    ret := 0
    for i:=0;i< m;i++{
        for j:=0;j<n;j++{
            if grid[i][j] == 1 {
                // 左 上 右 下
                if j - 1 < 0 || (j - 1 >= 0 && grid[i][j -1] == 0) {
                    ret += 1
                }
                if i - 1 < 0 || (i - 1 >= 0 && grid[i-1][j] == 0) {
                    ret += 1
                }
                if j + 1 >= n || (j + 1 < n && grid[i][j + 1] == 0){
                    ret += 1
                }
                if i + 1 >= m || (i + 1 < m && grid[i+ 1][j] == 0) {
                    ret += 1
                }
            }
        }
    }
    return ret
}
func findWords(words []string) []string {
    m1Str := "qwertyuiopQWERTYUIOP"
    m2Str := "asdfghjklASDFGHJKL"
    m3Str := "zxcvbnmZXCVBNM"
    
    m1 := make(map[rune]int)
    m2 := make(map[rune]int)
    m3 := make(map[rune]int)
    for _,v := range m1Str{
        m1[v] = 1
    }
    for _,v := range m2Str{
        m2[v] = 2
    }
    for _, v := range m3Str {
        m3[v] = 3
    }
    ret := []string{}
    for _, item := range words {
        
        flag := -1
        count := 0
        if _,ok := m1[rune(item[0])];ok {
            flag = 1
        }
        if _,ok := m2[rune(item[0])];ok {
            flag = 2
        }
        if _,ok := m3[rune(item[0])];ok {
            flag = 3
        }
        for _,v := range item {
            switch flag {
            case 1:
                if _,ok := m1[v];ok{
                    count ++
                }
                continue
            case 2:
                if _,ok := m2[v];ok{
                    count ++
                }
                continue
            case 3:
                if _,ok := m3[v];ok{
                    count ++
                }
                continue
            default:
            }
        }
        
        if flag != -1 && count == len(item) {
            ret = append(ret, item)
        }
    }
    return ret
}
func frequencySort(nums []int) []int {
    // 计算出现频率
    nMap := make(map[int]int)
    for _,v := range nums {
        nMap[v] += 1
    }
    nLen := len(nMap)
    ret := []int{}
    
    for i :=0; i < nLen; i++{
        minVal := 101
        tmp := []int{}
        for _, v := range nMap{
            if v < minVal {
                minVal = v
            }
        }
        for k, v := range nMap {
            if v == minVal {
                for j := 0 ; j < minVal; j++{
                    tmpV := k
                    tmp = append(tmp, tmpV)
                    delete(nMap, k)
                }
            }
        }
        sort.Ints(tmp)
        for y := len(tmp) -1; y >= 0; y--{
            ret = append(ret, tmp[y])
        }
    }
    return ret
}
func countSubstrings(s string, t string) int {
    return 0
}
func distributeCandies(candies []int) int {
    m := make(map[int]struct{})
    for _,v := range candies {
        m[v] = struct{}{}
    }
    return min(len(m), len(candies) / 2)
}
func findLHS(nums []int) int {
    // 和谐数组
    nMap := make(map[int]int)
    max := -1 << (32 -1)
    for _,v := range nums {
        nMap[v] += 1
    }
    for k,v := range nMap {
        if v2,ok := nMap[k - 1];ok {
            if v2 + v > max {
                max = v2 + v
            }
        }
        if v2, ok := nMap[k + 1];ok {
            if v2 + v > max {
                max = v2 + v
            }
        }
    }
    return max
}
func findRestaurant(list1 []string, list2 []string) []string {
    //
    m1 := make(map[string]int)
    m2 := make(map[string]int)
    for i,v := range list1 {
        index := i
        m1[v] = index
    }
    for i,v := range list2 {
        index := i
        m2[v] = index
    }
    ret := []string{}
    minIndex := 1 << (32 -1)
    for k,v := range m1 {
        if v2, ok := m2[k];ok {
            if minIndex > v + v2 {
                minIndex = v + v2
            }
        }
    }
    for k,v := range m1 {
        if v2, ok := m2[k];ok && (v2 + v == minIndex) {
            ret = append(ret, k)
        }
    }
    return ret
}
func findErrorNums(nums []int) []int {
    // 寻找重复值
    nMap := make(map[int]int)
    ret := []int{}
    sum := 0
    for _,v := range nums {
        nMap[v] += 1
        sum += v
        if mv,_ := nMap[v];mv == 2 {
            tmp := v
            ret = append(ret, tmp)
        }
    }
    // 计算差值
    standard := (len(nums) * (len(nums) + 1)) / 2
    ret = append(ret, (ret[0] + (standard - sum)))
    return ret
}
func validMountainArray(A []int) bool {
    // 两端往中间,最后重合,说明是山峰
    if len(A) <= 2 {
        return false
    }
    start := 0
    end := len(A) -1
    for start < end {
        if start + 1 <= end && A[start + 1] > A[start]{
            start ++
        } else if end - 1 >= start && A[end] < A[end -1] {
            end --
        } else {
            break
        }
    }
    return start == end && start != 0 && end != len(A) -1
}


type Employee struct {
     Id int
     Importance int
     Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
    eMap := make(map[int]*Employee)
    for _,v := range employees{
        eMap[v.Id] = v
    }
    e,_ := eMap[id]
    ret := 0
    ret += e.Importance
    for _,v := range e.Subordinates{
       ret += getImportance(employees,v)
    }
    return ret
}
func longestWord(words []string) string {
    sMap := make(map[string]int)
    for _, v := range words {
        sMap[v] = len(v)
    }
    result := ""
    haveResult := false
    for k,_ := range sMap {
        setable := true
        tmp := k
        for i := 1;i < len(tmp);i ++{
            if _,ok := sMap[string(tmp[:i])]; !ok{
                setable = false
            }
        }
        if setable && !haveResult {
            result = tmp
            haveResult = true
        }
        if setable && haveResult {
            if len(result) < len(tmp){
                result = tmp
            }
            if len(result) == len(tmp) && result > tmp {
                result = tmp
            }
        }
    }
    if haveResult {
        return result
    }
    return ""
}
func shortestCompletingWord(licensePlate string, words []string) string {
    lMap := make(map[rune]int)
    ulicensePlate  := strings.ToUpper(licensePlate)
    for _,v := range ulicensePlate {
        if v >= 'A' && v <= 'Z' {
            lMap[v] += 1
        }
    }
    result := ""
    minLen := 1 << (32 -1)
    for _, str := range words {
        uStr := strings.ToUpper(str)
        uStrMap := make(map[rune]int)
        for _, v := range uStr {
            if v >= 'A' && v <= 'Z' {
                uStrMap[v] += 1
            }
        }
        setAble := true
        for k, v := range lMap {
            if v2,ok := uStrMap[k]; !ok || v2 < v {
                setAble = false
            }
        }
        if setAble {
            if len(str) < minLen {
                minLen = len(str)
                result = str
            }
        }
    }
    return result
}
type MyHashSet struct {
    keys []int
}
func Constructor2() MyHashSet {
    return MyHashSet{}
}
func (this *MyHashSet) Add(key int)  {
    for _,v := range this.keys{
        if key == v {
            return
        }
    }
    this.keys = append(this.keys, key)
}


func (this *MyHashSet) Remove(key int)  {
    tmp := []int{}
    for _, v := range this.keys {
        if key == v {
            continue
        }else{
            tmp = append(tmp, v)
        }
    }
    this.keys = tmp
}


/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
    for _,v := range this.keys {
        if v == key {
            return true
        }
    }
    return false
}

type MyHashMap struct {
    // 双数组
    keys []int
    values []int
}


/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int)  {
    for i,v := range this.keys{
        if v == key {
            this.values[i] = value
            return
        }
    }
    this.keys = append(this.keys, key)
    this.values = append(this.values, value)
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
    for i,v := range this.keys{
        if v == key {
            return this.values[i]
        }
    }
    return -1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int)  {
    newKeys := []int{}
    newValues := []int{}
    for i,v := range this.keys {
        if v == key{
            continue
        }
        newKeys = append(newKeys,v)
        newValues = append(newValues, this.values[i])
    }
    this.keys = newKeys
    this.values = newValues
}

func subdomainVisits(cpdomains []string) []string {
    // 依次解析
    cMap := make(map[string]int)
    for _,v := range cpdomains{
        count,_ :=strconv.Atoi(strings.Split(v," ")[0])
        cMap[strings.Split(v," ")[1]] += count
        tmp := v
        index := strings.Index(tmp, ".")
        for index != -1 {
            tmp = string(tmp[index + 1:])
            cMap[tmp] += count
            index = strings.Index(tmp, ".")
        }
    }
    ret := []string{}
    for k, v := range cMap{
        tmp := strconv.Itoa(v)
        ret = append(ret, tmp + " " + k)
    }
    return ret
}
func getMaximumGenerated(n int) int {
    nums := make([]int, n + 1)
    max := -1 << (32 -1)
    for i := 0;i<len(nums);i++{
        if i ==0 {
            nums[i] = 0
        }
        if i == 1{
            nums[i] = 1
        }
        if n >= 2 * i && 2 * i >= 2 {
            if 2 * i < len(nums) {
                nums[2 * i] = nums[i]
            }
        }
        if n >= (2 * i + 1) && (2 * i + 1) >= 2{
            if (2 * i + 1) < len(nums) {
                nums[2 * i + 1] = nums[i] + nums[i + 1]
            }
        }
        if nums[i] > max {
            max = nums[i]
        }
    }
    return max
}
func minDeletions(s string) int {
    sMap := make(map[rune]int)
    for _,v := range s {
        sMap[v] += 1
    }
    nums := make([]int, len(sMap))
    //
    index := 0
    for _,v := range sMap {
        nums[index] = v
        index ++
    }
    // 排序
    sort.Ints(nums)
    // 删除
    step := 0
    for i := len(nums) -2; i >= 0; i -- {
        for nums[i] > 0 && nums[i] >= nums[i + 1]{
            step += 1
            nums[i] -= 1
        }
    }
    return step
}
func repeatedNTimes(A []int) int {
    n := len(A) / 2
    aMap := make(map[int]int)
    for _, v := range A {
        aMap[v] += 1
        v2,_ := aMap[v]
        if v2 >= n {
            return v
        }
    }
    return 0
}
func isAlienSorted(words []string, order string) bool {
    // 字典序,
    indexMap := make(map[rune]int)
    for i,v := range order {
        indexMap[v] = i
    }
    for i := 1;i<len(words);i++{
        first := words[i - 1]
        second := words[i]
        wlen := len(first)
        wflag := false
        if len(second) < wlen {
            wlen = len(second)
            wflag = true
        }
        for j := 0; j < wlen;j ++ {
            indexFirst,_ := indexMap[rune(first[j])]
            indexSecond,_:= indexMap[rune(second[j])]
            if indexFirst > indexSecond {
                return false
            }else if indexFirst < indexSecond {
                // 字典序
                wflag = false
                break
            }
        }
        
        //// 相同的地方[aab aa]
        if wflag {
           return false
        }
    }
    return true
}
func powerfulIntegers(x int, y int, bound int) []int {
    rMap := make(map[int]struct{})
    for i := 0; i < 20 && int(math.Pow(float64(y),float64(i))) <= bound;i++{
        for j := 0;j < 20 && int(math.Pow(float64(x),float64(j)))<=bound;j ++{
            a := int(math.Pow(float64(y),float64(i))) + int(math.Pow(float64(x),float64(j)))
            if a <= bound {
                rMap[a] = struct{}{}
            }
        }
    
    }
    ret := []int{}
    for k,_ := range rMap{
        ret = append(ret, k)
    }
    return ret
}
func commonChars(A []string) []string {
    // 为每个字符串建立map,然后取map中元素的并集
    mapArr := []map[rune]int{}
    for _,v := range A {
        tempMap := make(map[rune]int)
        for _,c := range v {
            tempMap[c] += 1
        }
        mapArr = append(mapArr, tempMap)
    }
    tmpM := mapArr[0]
    retMap := make(map[rune]int)
    for k,v := range tmpM {
        minVal := v
        addAble := true
        for i := 1;i<len(mapArr);i++{
            if v2, ok := (mapArr[i])[k];!ok {
                addAble = false
            }else{
                if v2 < minVal {
                    minVal = v2
                }
            }
        }
        if addAble {
            retMap[k] = minVal
        }
    }
    ret := []string{}
    for k,v := range retMap{
        for i := 0;i < v; i++{
            ret = append(ret, string(k))
        }
    }
    return ret
}
func countCharacters(words []string, chars string) int {
    cMap := make(map[rune]int)
    for _,v := range chars {
        cMap[v] += 1
    }
    ret := 0
    for _, v := range words {
        vMap := make(map[rune]int)
        for _, c := range v {
            vMap[c] += 1
        }
        addAble := true
        for k2, v2 := range vMap{
            if v3,ok := cMap[k2]; !ok || v3 < v2 {
                addAble = false
                break
            }
        }
        if addAble {
            ret += len(v)
        }
    }
    return ret
}
func findOcurrences(text string, first string, second string) []string {
    words := strings.Split(text, " ")
    ret := []string{}
    for i := 1;i < len(words);{
        if words[i -1] == first && words[i] == second {
            if i + 1 < len(words) {
                ret = append(ret, words[i + 1])
                i += 1
                continue
            }
        }
        i ++
    }
    return ret
}
func maxNumberOfBalloons(text string) int {
    bMap := make(map[rune]int)
    for _,v := range "balloon"{
        bMap[v] += 1
    }
    ret := 1 << (32 -1)
    tMap := make(map[rune]int)
    for _,v := range text {
        if _, ok := bMap[v];ok {
            tMap[v] += 1
        }
    }
    for k,v := range bMap{
        if v2,ok := tMap[k];!ok {
            return 0
        }else{
            if v2 / v < ret {
                ret = v2 / v
            }
        }
    }
    return ret
}
func canPermutePalindrome(s string) bool {
    // 有个key值大于等于2便可以判定
    if len(s) < 2 {
        return true
    }
    sMap := make(map[rune]int)
    for _, v := range s {
        sMap[v] += 1
    }
    //
    oddNum := 1
    for _,v := range sMap{
        if v % 2 != 0 {
            oddNum -= 1
        }
        if oddNum < 0 {
            return false
        }
    }
    return true
}
func uniqueOccurrences(arr []int) bool {
    aMap := make(map[int]int)
    nMap := make(map[int]struct{})
    for _,v := range arr {
        aMap[v] += 1
    }
    for _,v := range aMap{
        if _,ok := nMap[v];!ok {
            nMap[v] = struct{}{}
        }else{
            return false
        }
    }
    return true
}
func smallerNumbersThanCurrent(nums []int) []int {
    nMap := make(map[int]int)
    for _, v := range nums {
        nMap[v] += 1
    }
    ret := []int{}
    for _,v := range nums {
        t := 0
        for i := v -1;i >= 0;i -- {
            if n, ok := nMap[i];ok {
                t += n
            }
        }
        ret = append(ret, t)
    }
    return ret
}
func numIdenticalPairs(nums []int) int {
    ret := 0
    for i := 0;i<len(nums);i++{
        for j := i + 1;j < len(nums);j++{
            if nums[i] == nums[j] {
                ret += 1
            }
        }
    }
    return ret
}
func relativeSortArray(arr1 []int, arr2 []int) []int {
    a1Map := make(map[int]int)
    for _,v := range arr1{
        a1Map[v] += 1
    }
    ret := []int{}
    for _,v := range arr2 {
        if v2,ok := a1Map[v];ok {
            for i := 0;i < v2 ;i ++{
                ret = append(ret, v)
            }
        }
        delete(a1Map, v)
    }
    tmp := []int{}
    for k,v := range a1Map{
        for i := 0;i<v;i ++{
            tmp = append(tmp, k)
        }
    }
    sort.Ints(tmp)
    ret = append(ret, tmp...)
    return ret
}
func decrypt(code []int, k int) []int {
    //
    tmp := []int{}
    if k > 0 {
        for i := 0; i < len(code);i++{
            replace := 0
            for j := i + 1; j < k + i + 1;j ++{
                replace += code[j % len(code)]
            }
            tmp = append(tmp, replace)
        }
    }else if k < 0 {
        for i := 0;i < len(code);i ++{
            replace := 0
            for j := i - 1;j > i + k - 1;j --{
                replace += code[(j + len(code)) % len(code)]
            }
            tmp = append(tmp, replace)
        }
    }else if k == 0{
        for i := 0;i <len(code);i ++{
            tmp = append(tmp,0)
        }
    }
    return tmp
}
func minimumDeletions(s string) int {
    minDelete := len(s)
    // 计算分割点0....n时a的个数
    // 计算分割点0....n时b的个数
    // 对于a,当前值为a,对应的分割点a值等于(i-1) + 1
    
    // 对于b,从后往前数,计算过程和a一致
    alen := make([]int, len(s))
    blen := make([]int, len(s))
    //计算a
    if s[0] == 'a'{
        alen[0] = 1
    }
    for i := 1; i < len(s);i ++{
        if s[i] == 'a' {
            alen[i] = alen[i -1] + 1
        }else{
            alen[i] = alen[i -1]
        }
    }
    // 计算b
    if s[len(s) - 1] == 'b'{
        blen[len(s) -1] = 1
    }
    for i := len(s) - 2; i >= 0;i --{
        if s[i] == 'b'{
            blen[i] = blen[i + 1] + 1
        }else{
            blen[i] = blen[i + 1]
        }
    }
    for i := 0; i < len(s); i++{
        if len(s) - (alen[i] + blen[i]) < minDelete{
            minDelete = len(s) - (alen[i] + blen[i])
        }
    }
    
    return minDelete
}
func removeKdigits(num string, k int) string {
    tmp := []rune(num)
    for i := 0;i < k;i ++{
        // 删除高位肯定是赚的
        for j := 0;j < len(tmp);j++{
            if j == len(tmp) - 1 {
                tmp = tmp[:j]
                break
            }else {
                if tmp[j] > tmp[j + 1]{
                    tmp = append(tmp[:j], tmp[j + 1:]...)
                    break
                }
            }
        }
    }
    // 去零输出
    ret := []rune{}
    deleteable := true
    for _,v := range tmp{
        if v == '0' && deleteable{
            continue
        }else{
            deleteable = false
            ret = append(ret, v)
        }
    }
    if len(ret) <= 0{
        return "0"
    }
    return string(ret)
}
func threesum()  {
    
}
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
    // 距离计算: |r0 - Rn| + |c0 - Cn|。其中Rn(0-R), Cn(0 - C)。
    ret := [][]int{}
    tmp := [][]int{}
    dis := []int{} // 距离
    for i := 0;i < R;i ++{
        for j := 0;j < C;j ++{
            d := abs(i - r0) + abs(j - c0)
            tmp = append(tmp, []int{i, j})
            dis = append(dis, d)
        }
    }
    for i := 0; i < len(dis);i ++{
        minIndex := 0
        minDis := 1 << (32 - 1)
        for j := 0;j < len(dis); j ++{
            if dis[j] < minDis {
                minDis = dis[j]
                minIndex = j
            }
        }
        ret = append(ret, tmp[minIndex])
        dis[minIndex] = 1 << (32 - 1)
    }
    
    return ret
}
func threeSum(nums []int) [][]int {
    ret := [][]int{}
    sort.Ints(nums)
    nMap := make(map[int]int)
    for i,v := range nums{
        nMap[v] = i
    }
    for first := 0; first < len(nums); first++{
        // 定位第一个元素
        if first > 0 && nums[first] == nums[first - 1]{
            // 去掉重复的
            continue
        }
        for second := first + 1;second < len(nums); second ++{
            // 定位第二个元素
            if second > first + 1 && nums[second] == nums[second - 1]{
                continue
            }
            k := -1 * (nums[first] + nums[second])
            if v,ok := nMap[k];ok && v > second {
                ret = append(ret, []int{nums[first], nums[second], k})
            }
        }
    }
    return ret
}
func fourSum(nums []int, target int) [][]int {
    n := len(nums)
    sort.Ints(nums)
    ret := [][]int{}
    nMap := make(map[int]int)
    for i, v := range nums {
        nMap[v] = i
    }
    for first := 0; first < n; first ++{
        for second := first + 1;second < n;second ++{
            for third := second + 1;third < n; third ++{
                if v, ok := nMap[target - (nums[first] + nums[second] + nums[third])];ok && v > third {
                    addAble := true
                    tmp := []int{nums[first], nums[second], nums[third], nums[v]}
                    for _,v := range ret {
                        if v[0] == tmp[0] && v[1] == tmp[1] && v[2] == tmp[2] && v[3] == tmp[3] {
                            addAble = false
                        }
                    }
                    if addAble {
                        ret = append(ret, tmp)
                    }
                }
            }
        }
    }
    return ret
}
func sortList(head *ListNode) *ListNode {
    tmp := []int{}
    for cur :=head;cur != nil;cur = cur.Next {
        tmp = append(tmp, cur.Val)
    }
    sort.Ints(tmp)
    index := 0
    for cur := head;cur != nil;cur = cur.Next{
        cur.Val = tmp[index]
        index ++
    }
    return head
}
func sumNums(n int) int {
    ret := 0
    var sum func(int)bool
    sum = func(i int) bool {
        ret += n
        return n > 0 && sum(n - 1)
    }
    return ret
}
func reverseWords2(s string) string {
    tmp := strings.Split(s, " ")
    ret := ""
    for i := len(tmp) - 1;i >=0;i --{
        if tmp[i] == "" {
            continue
        }
        ret += " " + tmp[i]
    }
    return strings.TrimSpace(ret)
}
func findMinArrowShots(points [][]int) int {
    // 求数组的交集,通过交集计算之后,独立的数组每一个都需要一根箭
    // pMap 存储气球的信息,k 为开始,v 为结束
    // 排序
    sort.Slice(points, func(i, j int) bool { return points[i][1] < points[j][1] })
    pMap := make(map[int]int)
    for _, p := range points {
        addAble := true
        for k,v := range pMap{
            if p[0] < k && p[1] >= k && p[1] <= v {
                //  [2 ,  8]
                //[1, 6]
                //[1.     8]
                // 新的数组为:(k, p[1)
                pMap[k] = p[1]
                addAble = false
                break
            } else if p[0] <= k && p[1] >= v {
                //  [2,8]
                //[1,     9]
                // 交集:[2,8]
                addAble = false
                break
            } else if p[0] >= k && p[0] < v && p[1] <= v {
                //[2,   8]
                //  [3,7]
                delete(pMap,k)
                pMap[p[0]] = p[1]
                addAble = false
                break
            } else if p[0] > k && p[0] <= v && p[1] > v {
                // [2, 8]
                //    [8, 12]
                delete(pMap, k)
                pMap[p[0]] = v
                addAble = false
            }
        }
        if addAble {
            pMap[p[0]] = p[1]
        }
    }
    return len(pMap)
}
func modifyString(s string) string {
    sb := []rune(s)
    for i := 0; i < len(sb);i ++{
        if sb[i] == '?' {
            for j := 'a'; j < 'z';j ++{
               if i > 0 && sb[i - 1] != j {
                   if (i + 1 < len(sb) && sb[i + 1] != j) || i + 1 >= len(sb) {
                       // a?b || a?
                       sb[i] = j
                       break
                   }
               }else if i == 0 {
                   if i + 1 >= len(sb){
                       sb[i] = j
                       break
                   }else{
                       if sb[i + 1] != j{
                           sb[i] = j
                           break
                       }
                   }
               }
            }
        }
    }
    return string(sb)
}

func isPalindrome4(s string) bool {
    start,end := 0, len(s) - 1
    for start < end {
        if s[start] != s[end] {
            return false
        }else{
            start ++
            end --
        }
    }
    return true
}
func longestPalindrome2(s string) string {
    max := ""
    if isPalindrome4(s) {
        return s
    }
    n := len(s)
    dp := make([][]bool, n)
    for i := 0;i < n;i++{
        dp[i] = make([]bool, n)
    }
    for i := 1;i < n;i -- {
        // 滑动窗口
        initWindows := false
        windows := []byte{}
        for j := i;j < len(s);{
            if !initWindows {
                for k := 0;k < i;k ++{
                    windows = append(windows, s[k])
                }
                initWindows = true
            }else{
                // 滑动
                windows = append([]byte{}, windows[1:]...)
                windows = append(windows, s[j])
                j++
            }
            if isPalindrome4(string(windows)) && len(windows) > len(max) {
                max = string(windows)
                break
            }
        }
    }
    return max
}

func fourSumCount(A []int, B []int, C []int, D []int) int {
    ans := 0
    abMap := make(map[int]int)
    for _,v := range A {
        for _,v2 := range B {
            abMap[v + v2] += 1
        }
    }
    for _,c := range C {
        for _,d := range D {
            if v,ok := abMap[-(c + d)];ok {
                ans += v
            }
        }
    }
    return ans
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
    tree := &TreeNode{
        Val:   0,
        Left:  nil,
        Right: nil,
    }
    if t1 == nil && t2 == nil {
        return nil
    }
    if t1 != nil && t2 != nil {
        tree.Val = t1.Val + t2.Val
        tree.Left = mergeTrees(t1.Left,t2.Left)
        tree.Right = mergeTrees(t1.Right,t2.Right)
    }else if t1 != nil && t2 == nil{
        tree.Val = t1.Val
        tree.Left = mergeTrees(t1.Left,nil)
        tree.Right = mergeTrees(t1.Right, nil)
    }else if t1 == nil && t2 != nil {
        tree.Val = t2.Val
        tree.Left = mergeTrees(nil,t2.Left)
        tree.Right = mergeTrees(nil, t2.Right)
    }
    return tree
}

func buildTree(preorder []int, inorder []int) *TreeNode {
    // 前序定节点,中序定左右
    if len(preorder) <= 0 {
        return nil
    }
    nodeVal := preorder[0]
    left := []int{}
    right := []int{}
    // 取中序左右
    for i := 0;i < len(inorder);i ++{
        if inorder[i] == nodeVal {
            left = inorder[:i]
            right = inorder[i + 1:]
        }
    }
    // 取前序左右
    tree := &TreeNode{
        Val:   nodeVal,
        Left:  buildTree(preorder[1: len(left) + 1], left),
        Right: buildTree(preorder[1 + len(left):], right),
    }
    return tree
}

func reversePairs(nums []int) int {
    // map记录值和索引,
    n := len(nums)
    ans := 0
    // 将第i位的最大值记录下来,
    for i := 0;i < n;i ++{
        
        
        for j := i + 1;j < n;j ++{
            if nums[i] > 2 *nums[j] {
                ans += 1
            }
        }
    }
    return ans
}
func mostCompetitive(nums []int, k int) []int {
    if len(nums) <= k {
       return nums
    }
    ans := []int{}
    stack := make([]int, len(nums))
    top := -1
    for i := 0;i < len(nums);i ++{
        for top >= 0 && nums[i] < stack[top] && len(nums) - i + top >= k/*剩余的要够K*/  {
            top --
        }
        top ++
        stack[top] = nums[i]
    }
    for i := 0;i < k;i ++{
       ans = append(ans, stack[i])
    }
    return ans
}
type MaxQueue struct {
    queue []int
    maxQueue []int
}


func Constructor9() MaxQueue {
    return MaxQueue{queue: []int{}, maxQueue: []int{}}
}


func (this *MaxQueue) Max_value() int {
    if len(this.maxQueue) < 1 {
        return -1
    }
    return this.maxQueue[0]
}


func (this *MaxQueue) Push_back(value int)  {
    this.queue = append(this.queue, value)
    // 最大值
    for len(this.maxQueue) >= 0 && this.maxQueue[len(this.maxQueue) - 1] < value {
        this.maxQueue = this.maxQueue[ : len(this.maxQueue) -1]
    }
    this.maxQueue = append(this.maxQueue, value)
    this.queue = append(this.queue, value)
}


func (this *MaxQueue) Pop_front() int {
    if len(this.queue) < 1 {
        return - 1
    }
    ans := this.queue[0]
    this.queue = this.queue[1:]
    if ans == this.maxQueue[0] {
        this.maxQueue = this.maxQueue[1:]
    }
    return ans
}