type ListNode struct {
    key, val int
    next, prev *ListNode
}

type LRUCache struct {
    cache map[int]*ListNode
    capacity int
    head, tail *ListNode
}

func Constructor(capacity int) LRUCache {
    lrucache := LRUCache {
        capacity: capacity,
        cache: make(map[int]*ListNode),
        head: &ListNode{},
        tail: &ListNode{},
    }
    lrucache.head.next = lrucache.tail
    lrucache.tail.prev = lrucache.head
    return lrucache
}

func (this *LRUCache) insert(node *ListNode) {
    prev, next := this.tail.prev, this.tail
    prev.next = node
    next.prev = node
    node.next = next
    node.prev = prev
}
func (this *LRUCache) remove(node *ListNode) {
    prev, next := node.prev, node.next
    prev.next = next
    next.prev = prev
}

func (this *LRUCache) Get(key int) int {
    if node, exists := this.cache[key]; exists {
        this.remove(node)
        this.insert(node)
        return node.val
    }
    return -1
}

func (this *LRUCache) Put(key int, value int) {
    if node, exists := this.cache[key]; exists {
        this.remove(node)
        delete(this.cache, key)
    }
    node := &ListNode {
        key: key,
        val: value,
    }
    this.cache[key] = node
    this.insert(node)
    if len(this.cache) > this.capacity {
        lrucache := this.head.next
        this.remove(lrucache)
        delete(this.cache, lrucache.key)
    }
    
}
