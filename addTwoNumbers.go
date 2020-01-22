/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	 var digit, carry int
	 ret := l1
	 for l1.Next != nil && l2.Next != nil {
		digit = (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10
		l1.Val = digit
		l1 = l1.Next
		l2 = l2.Next
	 }
	 if l1.Next == nil {
		digit = (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10
		l1.Val = digit
		l2 = l2.Next
		l1.Next = l2
		for l2 != nil {
			digit = (l2.Val + carry) % 10
			carry = (l2.Val + carry) / 10
			l2.Val = digit
			l1 = l2
			l2 = l2.Next	
		}
	 } else if l2.Next == nil {
		digit = (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10
		l1.Val = digit
		l2 = l1.Next
		for l2 != nil {
			digit = (l2.Val + carry) % 10
			carry = (l2.Val + carry) / 10
			l2.Val = digit
			l1 = l2
			l2 = l2.Next
		}
	 }
	 if carry > 0 {
		l1.Next = &ListNode{carry, nil} 
	 }
	 return ret
}