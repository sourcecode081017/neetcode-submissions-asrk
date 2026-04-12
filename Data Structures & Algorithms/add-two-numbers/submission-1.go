/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    curr := head
    carry := 0
    for l1 != nil && l2 != nil {
        res := l1.Val + l2.Val + carry
        carry = res / 10
        res = res % 10
        curr.Next = &ListNode{
            Val: res,
        }
        l1 = l1.Next
        l2 = l2.Next
        curr = curr.Next
    }
    for l1 != nil {
        v1 := l1.Val + carry
        carry = v1 / 10
        v1 = v1 % 10
        curr.Next = &ListNode {
            Val: v1,
        }
        l1 = l1.Next
        curr = curr.Next
    }
    for l2 != nil {
         v2 := l2.Val + carry
        carry = v2 / 10
        v2 = v2 % 10
        curr.Next = &ListNode {
            Val: v2,
        }
        l2 = l2.Next
        curr = curr.Next
    }
    if carry > 0 {
        curr.Next = &ListNode {
            Val: carry,
        }
        curr = curr.Next
    }
    return head.Next

}