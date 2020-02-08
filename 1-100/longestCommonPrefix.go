// vertical scanning
// time complexity : O(S), where S is the sum of all characters in all strings
func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	} else if n == 1 {
		return strs[0]
	}
	i := 0
	for ; i < len(strs[0]); i++{
		r := strs[0][i]
		for j := 1; j < n; j++ {
			if i >= len(strs[j]) || strs[j][i] != r {
				return strs[0][:i]
			}
		}
	}
	return strs[0][:i]
}


// divide and conquer *
// time complexity : O(S), where S is the sum of all characters in all strings
func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	} else if n == 1 {
		return strs[0]
	}
	prefix_left := longestCommonPrefix(strs[:n/2])
	prefix_right := longestCommonPrefix(strs[n/2:])
	return commonPrefix(prefix_left, prefix_right)
}
func commonPrefix(left string, right string) string {
	i := 0
	for ; i < len(left) && i < len(right) && left[i] == right[i] ; i++ {}
	return left[:i]
}


// binary search *
// time complexity : O(S⋅logn), where S is the sum of all characters in all strings
func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	} else if n == 1 {
		return strs[0]
	}
	minLen := math.MaxInt64
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}
	low, high := 0, minLen - 1
	for low <= high {
		mid := (low + high) / 2
		if isCommonPrefix(strs, mid) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return strs[0][:low]
}
func isCommonPrefix(strs []string, length int) bool {
	for i := 0; i <= length; i++ {
		ch := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != ch {
				return false
			}
		}
	}
	return true
}


// use trie *
// time complexity : O(S), where S is the sum of all characters in all strings
// build trie: O(S), search for LCP: O(m), where m is the length all strings(that's the worst case)
// space complexity : O(S)
// import "bytes"
type Trie struct {
	root *Node
}
type Node struct {
	r rune
	isEnd bool
	firstChild, nextSibling *Node
}
/** Initialize a trie. */
func Constructor() Trie {
	trie := new(Trie)
	trie.root = new(Node)
	return *trie
}
/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
	if this == nil || this.root == nil {
		return
	}
	parent := this.root
	var node, lastSibling *Node
	for _, v := range word {
		node = parent.firstChild
		lastSibling = nil
		for node != nil && node.r != v {
			lastSibling = node
			node = node.nextSibling
		}
		if node == nil {
			node = new(Node)
			node.r = v
			if lastSibling == nil {
				parent.firstChild = node
			} else {
				lastSibling.nextSibling = node
			}
		}
		parent = node
	}
	if node != nil {
		node.isEnd = true
	}
}
/** Search longest common perfix in the trie. **/
func (this *Trie) SearchLongestPerfix() string {
	if this == nil || this.root == nil {
		return ""
	}
	buffer := bytes.Buffer{}
	node := this.root.firstChild
	for node != nil && node.nextSibling == nil {
		buffer.WriteRune(node.r)
        if node.isEnd {
            break
        }
		node = node.firstChild
	}
	return buffer.String()
}
func longestCommonPrefix(strs []string) string {
	trie := Constructor()
	for _, str := range strs {
		if str == "" {		// trie can't handle empty string
			return ""
		}
		trie.Insert(str)
	}
	return trie.SearchLongestPerfix()
}