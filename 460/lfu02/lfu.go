package lfu02

import "container/list"

type LFUCache struct {
	cap, minFre int
	k2v         map[int]int
	k2f         map[int]int
	f2ks        map[int]*linkedHashMap
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:    capacity,
		minFre: 0,
		k2v:    make(map[int]int),
		k2f:    make(map[int]int),
		f2ks:   make(map[int]*linkedHashMap),
	}
}

func (this *LFUCache) Get(key int) int {
	if v, ok := this.k2v[key]; ok {
		this.incrFre(key)
		return v
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if v, ok := this.k2v[key]; ok {
		this.incrFre(key)
		if value != v {
			this.k2v[key] = value
		}
		return
	}

	if len(this.k2v) == this.cap {
		this.deleteLeast()
	}

	this.add(key, value)
}

func (this *LFUCache) incrFre(key int) {
	fre := this.k2f[key]
	newFre := fre + 1
	this.k2f[key] = newFre

	oldLhm := this.f2ks[fre]
	oldLhm.remove(key)
	if oldLhm.empty() {
		delete(this.f2ks, fre)
		if fre == this.minFre {
			this.minFre = newFre
		}
	}

	if _, ok := this.f2ks[newFre]; !ok {
		this.f2ks[newFre] = NewLinkedHashMap()
	}

	this.f2ks[newFre].pushBack(key)
}

func (this *LFUCache) deleteLeast() {
	lhm := this.f2ks[this.minFre]
	least := lhm.popHead()
	key := least.Value.(int)
	delete(this.k2v, key)
	delete(this.k2f, key)

	if lhm.empty() {
		delete(this.f2ks, this.minFre)
	}
}

func (this *LFUCache) add(key, value int) {
	this.k2f[key] = 1
	this.k2v[key] = value
	this.minFre = 1
	if _, ok := this.f2ks[1]; !ok {
		this.f2ks[1] = NewLinkedHashMap()
	}

	this.f2ks[1].pushBack(key)
}

type linkedHashMap struct {
	m    map[int]*list.Element
	list *list.List
}

func NewLinkedHashMap() *linkedHashMap {
	return &linkedHashMap{
		m:    make(map[int]*list.Element),
		list: list.New(),
	}
}

func (lh *linkedHashMap) remove(key int) {
	e := lh.m[key]
	lh.list.Remove(e)
	delete(lh.m, key)
}

func (lh *linkedHashMap) pushBack(key int) {
	e := lh.list.PushBack(key)
	lh.m[key] = e
}

func (lh *linkedHashMap) popHead() *list.Element {
	head := lh.list.Front()
	lh.list.Remove(head)
	delete(lh.m, head.Value.(int))
	return head
}

func (lh *linkedHashMap) empty() bool {
	return lh.list.Len() == 0
}
