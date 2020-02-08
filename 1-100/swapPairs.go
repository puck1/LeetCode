/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	pre := dummy
	for head != nil {
		next := head.Next
		if next != nil {
			pre.Next = next
			pre = head
			next.Next, head = head, next.Next
		} else {
			pre.Next = head
			pre = head
			break
		}
	}
	pre.Next = nil
	return dummy.Next
}