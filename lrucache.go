package lrucache

import "fmt"

type Node struct {
	prev  *Node
	next  *Node
	key   any
	value any
}

type LRUCache struct {
	capacity int32
	cache    map[any]*Node
	left     *Node
	right    *Node
}

func NewLRUCache(capacity int32) *LRUCache {
	left := &Node{}
	right := &Node{}
	left.next = right
	right.prev = left
	return &LRUCache{
		capacity: capacity,
		cache:    map[any]*Node{},
		left:     left,
		right:    right,
	}
}

func (c *LRUCache) Get(key any) any {
	node, ok := c.cache[key]
	if !ok {
		return nil
	}
	c.removeNode(node)
	c.insertNode(node)
	return node.value
}

func (c *LRUCache) Put(key any, value any) any {
	node, ok := c.cache[key]
	if ok {
		c.removeNode(node)
		delete(c.cache, key)
	}
	if len(c.cache) >= int(c.capacity) {
		lruItem := c.left.next
		delete(c.cache, lruItem.key)
		c.removeNode(lruItem)
	}

	newNode := &Node{key: key, value: value}
	c.insertNode(newNode)
	c.cache[key] = newNode

	return newNode.value
}

func (c *LRUCache) Clear() {
	for k := range c.cache {
		delete(c.cache, k)
	}
	c.left.next = c.right
	c.right.prev = c.left
}

func (c *LRUCache) Len() int32 {
	return int32(len(c.cache))
}

func (c *LRUCache) Print() {
	for k, v := range c.cache {
		fmt.Printf("%v: %v\n", k, v.value)
	}
}

func (c *LRUCache) removeNode(node *Node) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (c *LRUCache) insertNode(node *Node) {
	c.right.prev.next = node
	node.prev = c.right.prev
	node.next = c.right
	c.right.prev = node
}
