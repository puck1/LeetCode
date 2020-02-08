/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/

// use loser tree
// time O(SlogK), space O(K), where S is amount of element in lists
var ls []int
const MINKEY = math.MinInt64
const MAXKEY = math.MaxInt64

func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)
	if lists == nil || k == 0 {
		return nil
	}
	head := &ListNode{}
	tail := head
	createLoserTree(lists, k)
	for lists[ls[0]] != nil {
		minNow := lists[ls[0]].Val
		for lists[ls[0]] != nil && lists[ls[0]].Val == minNow {
			tail.Next = lists[ls[0]]
			tail = tail.Next
			lists[ls[0]] = lists[ls[0]].Next
		}
		adjust(ls[0], lists, k)
	}
	return head.Next
}

func createLoserTree(lists []*ListNode, k int) {
	ls = make([]int, k, k)
	lists = append(lists, &ListNode{
		Val:  MINKEY,
		Next: nil,
	})
	for i := 0; i < k; i++ {
		ls[i] = k
	}
	for i := k-1; i >= 0; i-- {
		adjust(i, lists, k)
	}
}

func adjust(i int, lists []*ListNode, k int) {
	t := (i + k) >> 1
	for t > 0 {
		now, lastLoser := MAXKEY, MAXKEY
		if lists[i] != nil {
			now = lists[i].Val
		}
		if lists[ls[t]] != nil {
			lastLoser = lists[ls[t]].Val
		}
		if now > lastLoser {
			i ,ls[t] = ls[t], i
		}
		t = t >> 1
	}
	ls[0] = i
}

// use heap (priority queue) *
// time O(SlogK), space O(K), where S is amount of element in lists
// maybe is slower than loser-tree method
var heap []*ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	createHeap(lists)
	for len(heap) > 1 {
		minValNow := heap[1].Val
		for heap[1] != nil && heap[1].Val == minValNow {	// repeated element
			tail.Next = heap[1]
			tail = heap[1]
			heap[1] = heap[1].Next
		}
		if heap[1] == nil {				// reduce scale of heap
			heap[1] = heap[len(heap)-1]
			heap = heap[:len(heap)-1]
		}
		adjust(1)
	}
	return head.Next
}

func createHeap(lists []*ListNode) {
	heap = make([]*ListNode, 1)			// abandon element of index-0
	for _, node := range lists {
		if node != nil {
			heap = append(heap, node)
		}
	}
	k := len(heap) - 1
	for i := k >> 1; i > 0; i-- {
		adjust(i)
	}
}

func adjust(i int) {						// top-down adjust heap
	k := len(heap)-1
	for i <= k>>1 {
		minKey := i<<1
		minVal := heap[minKey].Val
		if minKey+1 <= k && heap[minKey+1].Val < minVal {
			minKey++
			minVal = heap[minKey].Val
		}
		if heap[i].Val <= minVal {
			break
		} else {
			heap[i], heap[minKey] = heap[minKey], heap[i]	// swap
		}
		i = minKey
	}
}


 // divide and conquer *
 func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)
	if lists == nil || k == 0 {
		return nil
	}

	// low effiency with recursion
	// if k == 1 {return lists[1]}
 	// l1 := mergeKLists(lists[:k/2])
	// l2 := mergeKLists(lists[k/2:])
	// l = merge2Lists(l1, l2)

	// the following two non-recursive methods are fast
	for interval := 1; interval < k; interval <<= 1 {
		for i := 0; i < k - interval; i += interval<<1 {
			if i + interval < k {
				lists[i] = merge2Lists(lists[i], lists[i+interval])
			}
		}
	}

	// trick *
	// end := len(lists) - 1
	// for end > 0 {
	// 	start := 0
	// 	for end > start {
	// 		lists[start] = merge2Lists(lists[start], lists[end])
	// 		end --
	// 		start ++
	// 	}
	// }

	return lists[0]
}

func merge2Lists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}
	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}
	return head.Next
}