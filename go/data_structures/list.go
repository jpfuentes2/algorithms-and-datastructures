package data_structures

type List interface {
	Head() *Node
	Tail() *Node
	Insert(interface{})
	Delete() *Node
	Reverse()
	Each(func(*Node))
	Len() int
	IsEmpty() bool
}

type Node struct {
	value interface{}
	next  *Node
}

type SinglyLinkedList struct {
	head   *Node
	length int
}

var _ List = &SinglyLinkedList{}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{head: nil}
}

func (l *SinglyLinkedList) IsEmpty() bool {
	return l.length == 0
}

func (l *SinglyLinkedList) Len() int {
	return l.length
}

func (l *SinglyLinkedList) Insert(v interface{}) {
	node := &Node{value: v, next: l.head}
	l.head = node
	l.length++
}

func (l *SinglyLinkedList) Delete() *Node {
	head := l.head
	l.head = head.next
	l.length--
	return head
}

func (l *SinglyLinkedList) Head() *Node {
	return l.head
}

func (l *SinglyLinkedList) Tail() *Node {
	var n *Node
	tail := n
	for n = l.head; n != nil; n = n.next {
		tail = n
		continue
	}
	return tail
}

func (l *SinglyLinkedList) Reverse() {
	if l.IsEmpty() || l.head.next == nil {
		return
	}

	var prev, current, next *Node
	current = l.head

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	l.head = prev
}

func (l *SinglyLinkedList) Each(f func(n *Node)) {
	if l.IsEmpty() {
		return
	}

	var n *Node
	for n = l.head; n != nil; n = n.next {
		f(n)
	}
}
