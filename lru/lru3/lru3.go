package lru3

import (

	"fmt"

)

type Node struct{

	Key interface{}

	Value interface{}

	pre *Node

	next *Node

}

type LRUCache struct {

	limit int

	HashMap map[interface{}]*Node

	head *Node

	end *Node

}

func(l *LRUCache)removeNode(node *Node)interface{}{

	if node == l.end{

		l.end = l.end.pre

		l.end.next = nil

	} else if node == l.head{

		l.head = l.head.next

		l.head.pre = nil

	} else {

		node.pre.next = node.next

		node.next.pre = node.pre

	}

	return node.Key

}

func (l *LRUCache)addNode(node *Node){

	if l.end != nil {

		l.end.next = node

		node.pre = l.end

		node.next = nil

	}

	l.end = node

	if l.head == nil {

		l.head = node

	}

}

func (l *LRUCache)refreshNode(node *Node){

	if node == l.end{

		return

	}

	l.removeNode(node) // 从链表中的任意位置移除原来的位置

	l.addNode(node) // 添加到链表的尾部

}

// 构造

func Constructor(capacity int)LRUCache{

	lruCache := LRUCache{limit: capacity}

	lruCache.HashMap = make(map[interface{}]*Node, capacity)

	return lruCache

}

// 获取

func (l *LRUCache)Get(key interface{})interface{}{

	if v, ok := l.HashMap[key]; ok {

		l.refreshNode(v)

		return v.Value

	} else {

		return -1

	}

}

func (l *LRUCache)Put(key, value interface{}){

	if v, ok := l.HashMap[key]; !ok{

		if len(l.HashMap) >= l.limit {

			oldkey := l.removeNode(l.head)

			delete(l.HashMap, oldkey)

		}

		node := Node{Key:key, Value:value}

		l.addNode(&node)

		l.HashMap[key] = &node

	} else {

		v.Value = value

		l.refreshNode(v)

	}

}

func (l *LRUCache)getCache(){

	for n := l.head; n != nil; n = n.next{

		fmt.Println(n.Key, n.Value)

	}

}

func main(){

	cache := Constructor(3)

	cache.Put(11, 1)

	cache.Put(22, 2)

	cache.Put(33, 3)

	cache.Put(44, 4)

	v := cache.Get(33)

	fmt.Println(v)

	fmt.Println("========== 获取数据之后 ===============")

	cache.getCache()

}