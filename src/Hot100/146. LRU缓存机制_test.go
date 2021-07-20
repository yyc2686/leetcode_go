package Hot100

import "testing"

func TestLruCache3(t *testing.T) {
	lRUCache := Constructor(2)
	lRUCache.Put(2, 1)     // 缓存是 {2=1}
	lRUCache.Put(2, 2)     // 缓存是 {2=2}
	t.Log(lRUCache.Get(2)) // 返回 2
	lRUCache.Put(1, 1)     // 缓存是 {1=1,2=2}
	lRUCache.Put(4, 1)     // 该操作会使得关键字 2 作废，缓存是 {4=1, 1=1}
	t.Log(lRUCache.Get(2)) // 返回 -1 (未找到)
}

func TestLruCache2(t *testing.T) {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 0)     // 缓存是 {1=1}
	lRUCache.Put(2, 2)     // 缓存是 {1=1, 2=2}
	t.Log(lRUCache.Get(1)) // 返回 1
	lRUCache.Put(3, 3)     // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	t.Log(lRUCache.Get(2)) // 返回 -1 (未找到)
	lRUCache.Put(4, 4)     // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	t.Log(lRUCache.Get(1)) // 返回 -1 (未找到)
	t.Log(lRUCache.Get(3)) // 返回 3
	t.Log(lRUCache.Get(4)) // 返回 4
}

func TestLruCache1(t *testing.T) {
	lRUCache := Constructor(2)
	lRUCache.Put(1, 1)     // 缓存是 {1=1}
	lRUCache.Put(2, 2)     // 缓存是 {1=1, 2=2}
	t.Log(lRUCache.Get(1)) // 返回 1
	lRUCache.Put(3, 3)     // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	t.Log(lRUCache.Get(2)) // 返回 -1 (未找到)
	lRUCache.Put(4, 4)     // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	t.Log(lRUCache.Get(1)) // 返回 -1 (未找到)
	t.Log(lRUCache.Get(3)) // 返回 3
	t.Log(lRUCache.Get(4)) // 返回 4
}
