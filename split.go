package pinyinsplit

var (
	trie *trieTree
)

type trieNode struct {
	char        byte           //当前节点
	depth       int            //节点深度
	isFinal     bool           //终结节点
	childNodes  [256]*trieNode //非空表示子节点，以字符为索引，最多256个
	numberNodes int            //有多少个子节点，仅为了方便调试
}

func newTrieNode(depth int) *trieNode {
	return &trieNode{
		depth:   depth,
		isFinal: false,
	}
}

func (n *trieNode) insertPath(path []byte) {
	if n.depth == len(path) {
		n.isFinal = true
		return
	}
	c := path[n.depth]
	node := n.nextNode(c)
	if node == nil {
		//new node
		n.numberNodes += 1
		node = newTrieNode(n.depth + 1)
		node.char = c
		n.childNodes[int(c)] = node
	}
	node.insertPath(path)
}

// matchNode 路径查找
func (n *trieNode) matchNode(path []byte) *trieNode {
	currentNode := n
	for _, c := range path {
		currentNode = currentNode.nextNode(c)
		if currentNode == nil {
			//查找失败
			break
		}
	}
	if currentNode != nil && currentNode.isFinal {
		//查找成功
		return currentNode
	} else {
		return nil
	}
}

func (n *trieNode) nextNode(c byte) *trieNode {
	return n.childNodes[int(c)]
}

type trieTree struct {
	root *trieNode
}

// match
func (t *trieTree) match(path string) bool {
	node := t.root.matchNode([]byte(path))
	return node != nil
}

// searchLongest 找到成功匹配的最长路径
func (t *trieTree) searchLongest(path []byte) *trieNode {
	var longestNode *trieNode = nil
	nextNode := t.root
	for _, c := range path {
		nextNode = nextNode.nextNode(c)
		if nextNode == nil {
			//路径匹配失败，直接退出
			break
		}

		if nextNode.isFinal {
			//终结节点
			longestNode = nextNode
		}
	}
	return longestNode
}

func buildTrie(paths []string) *trieTree {
	root := newTrieNode(0)
	for _, path := range paths {
		root.insertPath([]byte(path))
	}
	return &trieTree{
		root: root,
	}
}

func (t *trieTree) split(spell string) (words []string) {
	path := []byte(spell)

	for len(path) > 0 {
		node := t.searchLongest(path)
		if node == nil {
			//查找失败时，忽略当前字节
			path = path[1:]
			continue
		}

		word := path[:node.depth]
		path = path[node.depth:]
		words = append(words, string(word))
	}
	return words
}

func init() {
	trie = buildTrie(pingYinTable)
}

func Split(spell string) []string {
	return trie.split(spell)
}
