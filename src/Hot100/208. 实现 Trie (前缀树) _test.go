package Hot100

import "testing"

func TestTrie1(t *testing.T) {
	trie := TrieConstructor()
	trie.Insert("apple")
	t.Log(trie.Search("apple"))   // 返回 True
	t.Log(trie.Search("app"))     // 返回 False
	t.Log(trie.StartsWith("app")) // 返回 True
	trie.Insert("app")
	t.Log(trie.Search("app")) // 返回 True
}
