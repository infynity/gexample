package lru

type LRUNode struct {
	key  string
	val  interface{}
	prev *LRUNode
	next *LRUNode
}

type LRU struct {
	dataMap  map[string]*LRUNode
	head     *LRUNode
	tail     *LRUNode
	capacity int
	count    int
}

func NewLRU(capacity int) LRU {
	head := &LRUNode{}
	tail := &LRUNode{}
	head.next = tail
	tail.prev = head
	return LRU{head: head, tail: tail, capacity: capacity, count: 0, dataMap: make(map[string]*LRUNode)}
}

func (L *LRU) Get(key string) interface{} {
	v, ok := L.dataMap[key]
	if !ok {
		return nil
	}
	//detach node
	L.detachNode(v)
	//insert first
	L.insertFront(v)
	return v.val
}
func (L *LRU) detachNode(node *LRUNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
func (L *LRU) insertFront(node *LRUNode) {
	node.next = L.head.next
	L.head.next = node
	node.prev = L.head
}

func (L *LRU) delLast() {
	tmp := L.tail.prev
	tmp.prev.next = L.tail
	L.tail.prev = tmp.prev
	tmp.next = nil
	tmp.prev = nil
	L.count = L.count - 1
	delete(L.dataMap, tmp.key)
}
func (L *LRU) Set(key string, val interface{}) {
	v, ok := L.dataMap[key]
	if !ok {
		node := &LRUNode{key: key, val: val}
		if L.count == L.capacity {
			L.delLast()
		}
		L.dataMap[key] = node
		L.insertFront(node)
		L.count = L.count + 1
	} else {
		L.detachNode(v)
		L.insertFront(v)
	}
}
