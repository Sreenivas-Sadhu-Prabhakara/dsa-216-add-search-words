package main

import "fmt"

// WordDictionary stores lowercase words and answers searches where '.'
// matches any single letter. A trie; dots fan the search out over all
// children at that depth.
type WordDictionary struct {
	children [26]*WordDictionary
	isWord   bool
}

// NewWordDictionary returns an empty dictionary.
func NewWordDictionary() *WordDictionary { return &WordDictionary{} }

// AddWord inserts a word (a-z only).
func (d *WordDictionary) AddWord(word string) {
	n := d
	for i := 0; i < len(word); i++ {
		c := word[i] - 'a'
		if n.children[c] == nil {
			n.children[c] = &WordDictionary{}
		}
		n = n.children[c]
	}
	n.isWord = true
}

// Search reports whether any stored word matches the pattern, where '.'
// stands for exactly one arbitrary letter.
func (d *WordDictionary) Search(pattern string) bool {
	return d.match(pattern, 0)
}

func (d *WordDictionary) match(pattern string, i int) bool {
	if i == len(pattern) {
		return d.isWord
	}
	if pattern[i] == '.' {
		for _, child := range d.children {
			if child != nil && child.match(pattern, i+1) {
				return true
			}
		}
		return false
	}
	child := d.children[pattern[i]-'a']
	return child != nil && child.match(pattern, i+1)
}

func main() {
	d := NewWordDictionary()
	d.AddWord("bad")
	d.AddWord("dad")
	d.AddWord("mad")
	fmt.Println(d.Search("pad"), d.Search("bad"), d.Search(".ad"), d.Search("b.."))
}
