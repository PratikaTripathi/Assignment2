package main

import "fmt"

type node struct {
	data string
	prev *node
	next *node
}

type doublyLinkedList struct {
	len  int
	tail *node
	head *node
}

type LRU struct {
	mp map[string]*node
	d  *doublyLinkedList
}

var windowSize int = 2

func initDoublyList() *doublyLinkedList {
	return &doublyLinkedList{}
}

func initMap() *(map[string]*node) {
	m := make(map[string]*node)
	return &m
}

func (d *doublyLinkedList) AddFrontNodeDLL(data string) *node {
	newNode := &node{
		data: data,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	}
	d.len++
	return newNode
}

func (d *doublyLinkedList) DeleteEndNodeDLL() {
	prevNode := d.tail.prev
	currentNode := d.tail
	prevNode.next = nil
	currentNode.prev = nil
	d.tail = prevNode
	d.len--
}

func (d *doublyLinkedList) DeleteNodeWhenFound(v *node) {
	if d.len == 1 {
		d.head = nil
		d.tail = nil
	} else if d.head == v {
		d.head = d.head.next
		d.head.prev = nil
	} else if d.tail == v {
		d.tail = d.tail.prev
		d.tail.next = nil
	} else {
		prevv := v.prev
		nextt := v.next
		prevv.next = nextt
		nextt.prev = prevv
	}
	d.len--
}

func (l *LRU) LRUCache(data string) {
	v, found := l.mp[data]
	if found {
		l.d.DeleteNodeWhenFound(v)
		delete(l.mp, data)
	} else if l.d.len >= windowSize {
		dataa := l.d.tail.data
		l.d.DeleteEndNodeDLL()
		delete(l.mp, dataa)
	}
	newnode := l.d.AddFrontNodeDLL(data)
	l.mp[data] = newnode
}

func (l * LRU) TraverseDLL() {
	fmt.Println("Elements in Doubly Linked List")
	for temp := l.d.head; temp != nil; temp = temp.next {
		fmt.Printf("%s\t", temp.data)
	}
	fmt.Println()
}

func main() {
	var lru LRU
	doublyList := initDoublyList()
	lru.d = doublyList
	lru.mp = *initMap()
	lru.LRUCache("A")
	lru.LRUCache("B")
	lru.TraverseDLL()
	lru.LRUCache("A")
	lru.TraverseDLL()
	lru.LRUCache("C")
	lru.TraverseDLL()
}
