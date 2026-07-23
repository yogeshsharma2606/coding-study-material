// Package tries contains trie (prefix tree) interview problems.
//
// A trie stores strings by shared prefixes: each node is a character, each root
// path spells a prefix. Lookups/inserts are O(L) in the word length - independent
// of how many words are stored - which is why tries power autocomplete,
// dictionaries, and prefix queries. A "bit trie" over the binary digits of
// integers solves max-XOR style problems.
package tries

// Trie supports insert, exact search, and prefix search over lowercase words.
type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func NewTrie() *Trie { return &Trie{} }

func (t *Trie) Insert(word string) {
	node := t
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if node.children[c] == nil {
			node.children[c] = &Trie{}
		}
		node = node.children[c]
	}
	node.isEnd = true
}

func (t *Trie) find(word string) *Trie {
	node := t
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if node.children[c] == nil {
			return nil
		}
		node = node.children[c]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	n := t.find(word)
	return n != nil && n.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.find(prefix) != nil
}

// WordDictionary supports adding words and searching with '.' wildcards.
type WordDictionary struct {
	root *Trie
}

func NewWordDictionary() *WordDictionary { return &WordDictionary{root: &Trie{}} }

func (w *WordDictionary) AddWord(word string) {
	node := w.root
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if node.children[c] == nil {
			node.children[c] = &Trie{}
		}
		node = node.children[c]
	}
	node.isEnd = true
}

// Search matches '.' against any single character via DFS branching.
func (w *WordDictionary) Search(word string) bool {
	var dfs func(node *Trie, i int) bool
	dfs = func(node *Trie, i int) bool {
		if i == len(word) {
			return node.isEnd
		}
		c := word[i]
		if c == '.' {
			for _, child := range node.children {
				if child != nil && dfs(child, i+1) {
					return true
				}
			}
			return false
		}
		child := node.children[c-'a']
		return child != nil && dfs(child, i+1)
	}
	return dfs(w.root, 0)
}

// ReplaceWords replaces each word in the sentence with the shortest dictionary
// root that is a prefix of it. Build a trie of roots, then for each word walk
// until hitting a root end.
func ReplaceWords(dictionary []string, sentence []string) []string {
	trie := NewTrie()
	for _, root := range dictionary {
		trie.Insert(root)
	}
	shortestRoot := func(word string) string {
		node := trie
		for i := 0; i < len(word); i++ {
			c := word[i] - 'a'
			if node.children[c] == nil {
				return word
			}
			node = node.children[c]
			if node.isEnd {
				return word[:i+1]
			}
		}
		return word
	}
	out := make([]string, len(sentence))
	for i, w := range sentence {
		out[i] = shortestRoot(w)
	}
	return out
}

// LongestWord returns the longest word buildable one character at a time, where
// every prefix is also a word; ties broken lexicographically smallest.
func LongestWord(words []string) string {
	trie := NewTrie()
	for _, w := range words {
		trie.Insert(w)
	}
	best := ""
	// DFS only through nodes that are word-ends (so every prefix is a word)
	var dfs func(node *Trie, path string)
	dfs = func(node *Trie, path string) {
		if len(path) > len(best) || (len(path) == len(best) && path < best) {
			best = path
		}
		for i := 0; i < 26; i++ {
			child := node.children[i]
			if child != nil && child.isEnd {
				dfs(child, path+string(rune('a'+i)))
			}
		}
	}
	dfs(trie, "")
	return best
}

// SuggestedProducts returns, for each prefix of searchWord, up to 3 lexicographically
// smallest products sharing that prefix.
func SuggestedProducts(products []string, searchWord string) [][]string {
	trie := NewTrie()
	for _, p := range products {
		trie.Insert(p)
	}
	var res [][]string
	prefix := ""
	for i := 0; i < len(searchWord); i++ {
		prefix += string(searchWord[i])
		node := trie.find(prefix)
		var suggestions []string
		if node != nil {
			collect(node, prefix, &suggestions)
		}
		res = append(res, suggestions)
	}
	return res
}

// collect gathers up to 3 words in lexicographic order under node.
func collect(node *Trie, path string, out *[]string) {
	if len(*out) == 3 {
		return
	}
	if node.isEnd {
		*out = append(*out, path)
	}
	for i := 0; i < 26 && len(*out) < 3; i++ {
		if node.children[i] != nil {
			collect(node.children[i], path+string(rune('a'+i)), out)
		}
	}
}

// MapSum inserts key-value pairs and sums the values of all keys with a prefix.
type MapSum struct {
	vals map[string]int
	root *sumNode
}

type sumNode struct {
	children [26]*sumNode
	sum      int
}

func NewMapSum() *MapSum { return &MapSum{vals: map[string]int{}, root: &sumNode{}} }

func (m *MapSum) Insert(key string, val int) {
	delta := val - m.vals[key] // support overwrites
	m.vals[key] = val
	node := m.root
	node.sum += delta
	for i := 0; i < len(key); i++ {
		c := key[i] - 'a'
		if node.children[c] == nil {
			node.children[c] = &sumNode{}
		}
		node = node.children[c]
		node.sum += delta
	}
}

func (m *MapSum) Sum(prefix string) int {
	node := m.root
	for i := 0; i < len(prefix); i++ {
		c := prefix[i] - 'a'
		if node.children[c] == nil {
			return 0
		}
		node = node.children[c]
	}
	return node.sum
}

// FindMaximumXOR returns the maximum a[i] XOR a[j] using a binary trie of the
// 32-bit numbers. For each number, greedily walk toward the opposite bit to
// maximize XOR.
func FindMaximumXOR(nums []int) int {
	type bitNode struct{ child [2]*bitNode }
	root := &bitNode{}
	const bits = 31
	insert := func(num int) {
		node := root
		for b := bits; b >= 0; b-- {
			bit := (num >> b) & 1
			if node.child[bit] == nil {
				node.child[bit] = &bitNode{}
			}
			node = node.child[bit]
		}
	}
	best := 0
	for _, num := range nums {
		insert(num)
	}
	for _, num := range nums {
		node := root
		cur := 0
		for b := bits; b >= 0; b-- {
			bit := (num >> b) & 1
			want := bit ^ 1 // prefer the opposite bit
			if node.child[want] != nil {
				cur |= 1 << b
				node = node.child[want]
			} else {
				node = node.child[bit]
			}
		}
		if cur > best {
			best = cur
		}
	}
	return best
}
