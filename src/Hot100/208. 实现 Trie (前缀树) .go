package Hot100

//208. 实现 Trie (前缀树)
//Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。
//
//请你实现 Trie 类：
//
//Trie() 初始化前缀树对象。
//void insert(String word) 向前缀树中插入字符串 word 。
//boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
//boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。
//
//
//示例：
//
//输入
//["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
//[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
//输出
//[null, null, true, false, true, null, true]
//
//解释
//Trie trie = new Trie();
//trie.insert("apple");
//trie.search("apple");   // 返回 True
//trie.search("app");     // 返回 False
//trie.startsWith("app"); // 返回 True
//trie.insert("app");
//trie.search("app");     // 返回 True
//
//
//提示：
//
//1 <= word.length, prefix.length <= 2000
//word 和 prefix 仅由小写英文字母组成
//insert、search 和 startsWith 调用次数 总计 不超过 3 * 104 次

type Trie struct {
	Next  [26]*Trie
	IsEnd bool
}

/** Initialize your data structure here. */
func TrieConstructor() Trie {
	trie := Trie{}
	trie.Next = [26]*Trie{}
	return trie
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	i := 0
	for ; i < len(word); i++ {
		if this.Next[word[i]-'a'] == nil {
			break
		}
		this = this.Next[word[i]-'a']
	}

	for ; i < len(word); i++ {
		this.Next[word[i]-'a'] = &Trie{}
		this = this.Next[word[i]-'a']
	}
	this.IsEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return this.IsEnd
	}
	if this == nil || this.Next[word[0]-'a'] == nil {
		return false
	}
	return this.Next[word[0]-'a'].Search(word[1:])
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}
	if this == nil || this.Next[prefix[0]-'a'] == nil {
		return false
	}
	return this.Next[prefix[0]-'a'].StartsWith(prefix[1:])
}
