package main

type Trie struct {
	isWord   bool
	children [26]*Trie
}

func Constructor() Trie {
	return Trie{

	}
}

func (this *Trie) Insert(word string) {
	cur := this
	for i, c := range word {
		n := c - 'a'

		if cur.children[n] == nil {
			cur.children[n] = &Trie{}
		}

		cur = cur.children[n]
		if i == len(word)-1 {
			cur.isWord = true
		}
	}
}

func (this *Trie) Search(word string) bool {
	cur := this
	for _, c := range word {
		n := c - 'a'
		if cur.children[n] == nil {
			return false
		}
		cur = cur.children[n]
	}
	return cur.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, c := range prefix {
		n := c - 'a'
		if cur.children[n] == nil {
			return false
		}
		cur = cur.children[n]
	}
	return true
}

func main() {

}
