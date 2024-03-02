package lru01

type LRUCache struct {
	dl  *doubleList
	m   map[int]*node
	cap int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		dl:  newDoubleList(),
		m:   make(map[int]*node),
		cap: capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if n, ok := this.m[key]; ok {
		this.makeRecently(key)
		return n.value
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.m[key]; ok {
		this.deleteKey(key)
		this.addRecently(key, value)
		return
	}

	if this.cap == this.dl.length() {
		this.deleteLeastRecently()
	}

	this.addRecently(key, value)
}

func (this *LRUCache) makeRecently(key int) {
	x := this.m[key]
	this.dl.remove(x)
	this.dl.append(x)
}

func (this *LRUCache) addRecently(key, value int) {
	n := &node{key: key, value: value}
	this.m[key] = n
	this.dl.append(n)
}

func (this *LRUCache) deleteKey(key int) {
	n := this.m[key]
	this.dl.remove(n)
	delete(this.m, key)
}

func (this *LRUCache) deleteLeastRecently() {
	first := this.dl.pop()
	delete(this.m, first.key)
}

type node struct {
	prev, next *node
	key, value int
}

type doubleList struct {
	head, tail *node
	size       int
}

func newDoubleList() *doubleList {
	head, tail := &node{}, &node{}
	head.next = tail
	tail.prev = head
	return &doubleList{
		head: head,
		tail: tail,
		size: 0,
	}
}

func (dl *doubleList) append(x *node) {
	x.prev = dl.tail.prev
	x.next = dl.tail
	dl.tail.prev.next = x
	dl.tail.prev = x
	dl.size++
}

func (dl *doubleList) remove(x *node) {
	x.prev.next = x.next
	x.next.prev = x.prev
	dl.size--
}

func (dl *doubleList) pop() *node {
	first := dl.head.next
	dl.remove(first)
	return first
}

func (dl *doubleList) length() int {
	return dl.size
}
