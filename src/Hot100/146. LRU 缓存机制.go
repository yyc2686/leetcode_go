package Hot100

//146. LRU 缓存机制
//运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。
//实现 LRUCache 类：
//
//LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
//int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
//void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。
//
//
//进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？
//
//
//
//示例：
//
//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4
//
//
//提示：
//
//1 <= capacity <= 3000
//0 <= key <= 10000
//0 <= value <= 105
//最多调用 2 * 105 次 get 和 put

// 哈希表+双向链表（哈希表尽可定位key是否存在，双向链表每个节点保存完整数据信息）
type DLinkNode struct {
	key, val  int
	pre, next *DLinkNode
}

type LRUCache struct {
	size       int
	capacity   int
	head, tail *DLinkNode
	cache      map[int]*DLinkNode
}

func Constructor(capacity int) LRUCache {
	lRUCache := LRUCache{capacity: capacity}
	lRUCache.cache = make(map[int]*DLinkNode, 0)
	// 引入头尾节点，使得插入和删除得到统一，不需要对头尾节点单独处理
	lRUCache.head, lRUCache.tail = &DLinkNode{}, &DLinkNode{}
	lRUCache.head.next = lRUCache.tail
	lRUCache.tail.pre = lRUCache.head
	return lRUCache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; !ok {
		return -1
	} else {
		// 删除key
		this.Pop(node)
		// 将key插入到最后
		this.Append(node)
		return node.val
	}
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; !ok {
		// 如果当前长度达到最大容量，删除头部节点
		if this.size == this.capacity {
			delete(this.cache, this.head.next.key)
			this.Pop(this.head.next)
			this.size--
		}

	} else {
		// 删除key
		this.Pop(node)
		this.size--
	}
	// 将key插入到最后
	node := &DLinkNode{key: key, val: value}
	this.Append(node)
	this.size++
	// 更新cache
	this.cache[key] = node
}

func (this *LRUCache) Pop(node *DLinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre

}

func (this *LRUCache) Append(node *DLinkNode) {
	this.tail.pre.next = node
	node.next = this.tail
	node.pre = this.tail.pre
	this.tail.pre = node
}
