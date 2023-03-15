package trie

type node struct {
	next   [26]*node
	isWord bool
}

type Trie struct {
	root *node
}

func (trie *Trie) Init() {
	trie.root = &node{}
}

func (trie *Trie) Add(str string) {
	curr := trie.root
	for _, ch := range str {
		idx := ch-'A'
		if (curr.next[idx]==nil){
			curr.next[idx] = &node{}
		}
		curr = curr.next[idx]
	}
	curr.isWord=true
}

func (trie *Trie) Search(str string) bool{
	curr := trie.root
	for _, ch := range str {
		idx := ch-'A'
		if (curr.next[idx]==nil){
			return false
		}
		curr = curr.next[idx]
	}
	return curr.isWord
}
