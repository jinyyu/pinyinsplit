package pinyinsplit

const (
	numberTrieNode int = 256 //子节点的个数，以字符为索引，最多256个
)

type trieNode struct {
	char     byte //当前节点
	depth    int  //深度
	isFinal  bool
	allNodes [numberTrieNode]*trieNode //所有子节点
}

func newTrieNode(depth int) *trieNode {
	return &trieNode{
		depth:   depth,
		isFinal: false,
	}
}

func (n *trieNode) insert(data []byte) {
	if n.depth == len(data) {
		n.isFinal = true
		return
	}
	c := data[n.depth]
	node := n.next(c)
	if node == nil {
		//new node
		node = newTrieNode(n.depth + 1)
		node.char = c
		n.allNodes[int(c)] = node
	}
	node.insert(data)
}

func (n *trieNode) search(data []byte) *trieNode {
	if len(data) == 0 {
		if n.isFinal {
			return n
		} else {
			return nil
		}
	}

	c := data[0]
	node := n.next(c)
	if node == nil {
		return nil
	}
	data = data[1:]
	return node.search(data)
}

func (n *trieNode) next(c byte) *trieNode {
	return n.allNodes[int(c)]
}

type trieTree struct {
	root *trieNode
}

func (t *trieTree) search(str string) bool {
	node := t.root.search([]byte(str))
	return node != nil
}

func buildTrie() *trieTree {
	root := newTrieNode(0)
	for _, pinyin := range pingyinData {
		root.insert([]byte(pinyin))
	}
	return &trieTree{
		root: root,
	}
}

func Split(spell string) []string {
	return []string{}
}
