package linkedlist

import "errors"

var ErrEmptyList = errors.New("sorry, the list is empty")

type Node struct {
	prev, next *Node
	Val        interface{}
}

type List struct {
	start, end *Node
}

func (l *List) PushFront(v interface{}) {
	n := &Node{next: l.start, Val: v}
	if l.start != nil {
		l.start.prev = n
	} else {
		l.end = n
	}
	l.start = n
}

func (l *List) PushBack(v interface{}) {
	n := &Node{prev: l.end, Val: v}
	if l.end != nil {
		l.end.next = n
	} else {
		l.start = n
	}
	l.end = n
}

func (l *List) PopFront() (interface{}, error) {
	if l.start == nil {
		return 0, ErrEmptyList
	}
	v := l.start.Val
	l.start = l.start.next
	if l.start == nil {
		l.end = nil
	} else {
		l.start.prev = nil
	}
	return v, nil
}

func (l *List) PopBack() (interface{}, error) {
	if l.end == nil {
		return 0, ErrEmptyList
	}
	v := l.end.Val
	l.end = l.end.prev
	if l.end == nil {
		l.start = nil
	} else {
		l.end.next = nil
	}
	return v, nil
}

func NewList(s ...interface{}) *List {
	l := new(List)
	for _, v := range s {
		l.PushBack(v)
	}
	return l
}

func (l *List) Reverse() {
	c := l.start
	for c != nil {
		c.prev, c.next = c.next, c.prev
		c = c.prev
	}
	l.start, l.end = l.end, l.start
}

func (l *List) First() *Node {
	return l.start
}

func (l *List) Last() *Node {
	return l.end
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}
