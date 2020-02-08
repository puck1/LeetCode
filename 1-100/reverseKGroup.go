/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

 // non-recursion with O(2n) time
 func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {					// if k==1, algorithm below is inefficient
		return head
	}
	dummy := &ListNode{}
	tail := dummy
	for head != nil {
		begin := head			// beginning pointer of a group
		for i := 1; i < k && head != nil; i++ {
			head = head.Next
		}
		if head != nil {
			p := begin			// work pointer
			next := head.Next	// start of next group
			for p != next {		// insert at front into group
				tmp := p.Next
				p.Next = tail.Next
				tail.Next = p
				p = tmp
			}
			tail = begin		// end of a group
			head = next
		} else {
			tail.Next = begin
		}
	}
	return dummy.Next
}

// recursive method with O(2n) time
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	dummy := head
	for i := 0; i < k; i++ {
		if dummy == nil {
			return head
		}
		dummy = dummy.Next
	}
	current, next, previous := head, head.Next, reverseKGroup(dummy, k)
	for current != dummy {
		next, current.Next = current.Next, previous
		previous, current = current, next		// previous expands forward
	}
	return previous
}