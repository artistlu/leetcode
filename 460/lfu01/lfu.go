package lfu01

import "container/list"

type LFUCache struct {
	cap, minFre int
	k2v         map[int]int
	k2fre       map[int]int
	f2ks        map[int]*linkedHashSet
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:    capacity,
		minFre: 0,
		k2v:    make(map[int]int),
		k2fre:  make(map[int]int),
		f2ks:   make(map[int]*linkedHashSet),
	}
}

func (this *LFUCache) Get(key int) int {
	if v, k := this.k2v[key]; k {
		this.incrFre(key)
		return v
	}

	return -1

}

func (this *LFUCache) Put(key int, value int) {
	if this.cap <= 0 {
		return
	}

	if _, k := this.k2v[key]; k {
		this.incrFre(key)
		this.k2v[key] = value
		return
	}

	if this.cap == len(this.k2v) {
		this.deleteLeast()
	}
	this.append(key, value)
}

func (this *LFUCache) incrFre(key int) {
	f := this.k2fre[key]
	nf := f + 1
	this.k2fre[key] = nf
	this.f2ks[f].delete(key)
	if this.f2ks[f].empty() {
		delete(this.f2ks, f)
		if this.minFre == f {
			this.minFre = nf
		}
	}

	if _, ok := this.f2ks[nf]; !ok {
		this.f2ks[nf] = newLinkedHashSet()
	}
	this.f2ks[nf].append(key)
}

func (this *LFUCache) append(key int, value int) {
	this.k2v[key] = value
	this.k2fre[key] = 1
	this.minFre = 1
	if _, ok := this.f2ks[1]; !ok {
		this.f2ks[1] = newLinkedHashSet()
	}
	this.f2ks[1].append(key)
}

func (this *LFUCache) deleteLeast() {
	k := this.f2ks[this.minFre].pop()
	if this.f2ks[this.minFre].empty() {
		delete(this.f2ks, this.minFre)
	}
	delete(this.k2fre, k)
	delete(this.k2v, k)
}

type linkedHashSet struct {
	m   map[int]*list.Element
	set *list.List
}

func newLinkedHashSet() *linkedHashSet {
	return &linkedHashSet{
		m:   make(map[int]*list.Element),
		set: list.New(),
	}
}

func (ls *linkedHashSet) append(key int) {
	//e := &list.Element{Value: key}
	e := ls.set.PushBack(key)
	ls.m[key] = e
}

func (ls *linkedHashSet) pop() int {
	e := ls.set.Front()
	ls.set.Remove(e)
	k, _ := e.Value.(int)
	delete(ls.m, k)
	return k
}

func (ls *linkedHashSet) delete(key int) {
	e := ls.m[key]
	ls.set.Remove(e)
	delete(ls.m, key)
}

func (ls *linkedHashSet) empty() bool {
	return len(ls.m) == 0
}
