package main

type LinkedList struct {
	Size int
	Head *Node
}

type Node struct {
	Val  int
	Next *Node
}

func (l *LinkedList) Add(val int) {
	node := &Node{Val: val}
	head := l.Head
	if head == nil {
		l.Head = node
	} else {
		for head.Next != nil {
			head = head.Next
		}
		head.Next = node
	}
}

func (l *LinkedList) Remove(val int) {
	head := l.Head
	if head == nil {
		return
	} else if head.Val == val {
		l.Head = head.Next
	} else {
		for head.Next != nil {
			if head.Next.Val == val {
				head.Next = head.Next.Next
			}
		}
	}
}
