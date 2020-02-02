type Trie struct {
	root *Node
}

type Node struct {
	r rune
	isEnd bool
	firstChild, nextSibling *Node
}

/** Initialize your data structure here. */
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


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if this == nil || this.root == nil {
		return false
	}
	parent := this.root
	var node *Node
	for _, v := range word {
		node = parent.firstChild
		for node != nil && node.r != v {
			node = node.nextSibling
		}
		if node == nil {
			return false
		} else {
			parent = node
		}
	}
	if node == nil || node.isEnd {
		return true
	} else {
		return false
	}
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if this == nil || this.root == nil {
		return false
	}
	node := this.root.firstChild
	for _, v := range prefix {
		for node != nil && node.r != v {
			node = node.nextSibling
		}
		if node == nil {
			return false
		} else {
			node = node.firstChild
		}
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */



 // anothor method for char 'a' to 'z'
 type Trie struct {
	root *Node
}

type Node struct {
	r byte
	isEnd bool
	children [26]*Node
}

/** Initialize your data structure here. */
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
	node := this.root
	for _, v := range word {
		if node.children[v-'a'] == nil {
			newNode := new(Node)
			newNode.r = byte(v)
			node.children[v-'a'] = newNode
		}
		node = node.children[v-'a']
	}
	node.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if this == nil || this.root == nil {
		return false
	}
	node := this.root
	for _, v := range word {
		node = node.children[v-'a']
		if node == nil {
			return false
		}
	}
	if node.isEnd {
		return true
	} else {
		return false
	}
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if this == nil || this.root == nil {
		return false
	}
	node := this.root
	for _, v := range prefix {
		node = node.children[v-'a']
		if node == nil {
			return false
		}
	}
	return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */