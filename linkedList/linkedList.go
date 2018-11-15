package linkedList

import "strings"

type List struct {
	first  *Node
	length int
	last   *Node
}

type Node struct {
	next *Node
	Val  string
	prev *Node
}

func (l *List) Length() int {

	return l.length
}

func (n *Node) Prev() *Node { return n.prev }

func (n *Node) Next() *Node { return n.next }

func (l *List) First() *Node { return l.first }

func (l *List) Last() *Node { return l.last }

func (l *List) Append(s string) {
	n := &Node{
		prev: l.last,
		Val:  s,
	}
	if l.last == nil {
		l.last = n
		l.first = n
	} else {
		l.last.next = n
	}
	l.last = n
	l.length++
}

//Search should return all strings that contain the string passed in
func (l *List) Search(s string) []string {
	results := []string{}
	node := l.first
	s = strings.ToLower(s)
	for node != nil {
		lower := strings.ToLower(node.Val)
		if strings.Contains(lower, s) {
			results = append(results, node.Val)
		}
		node = node.Next()
	}

	return results
}
