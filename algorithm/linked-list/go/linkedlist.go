package linkedlist

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func CreateNode(value int, next *Node) *Node {
	node := Node{
		value,
		next,
	}
	return &node
}

func (node *Node) String() string {
	if node.Next == nil {
		return fmt.Sprintf("%d", node.Value)
	} else {
		return fmt.Sprintf("%d -> %s", node.Value, node.Next.String())
	}
}

func (node *Node) Reverse() *Node {
	var pre *Node
	cur := node
	next := cur.Next
	cur.Next = pre

	for {
		if next == nil {
			break
		} else {
			pre = cur
			cur = next
			next = next.Next

			cur.Next = pre
		}
	}
	return cur
}

func (node *Node) Reverse1() *Node {
	var pre *Node
	cur := node

	for {
		if cur == nil {
			return pre
		} else {
			cur.Next, pre, cur = pre, cur, cur.Next
		}
	}
}

func (node *Node) ReverseNearby() *Node {
	var pre *Node
	var pre1 *Node
	head := node
	cur := node
	index := 0
	for {
		if cur == nil {
			break
		} else {
			index++
			next := cur.Next
			if index%2 == 0 {
				cur.Next = pre
				if index == 2 {
					head = cur
				}
				pre.Next = next
				if pre1 != nil {
					pre1.Next = cur
				}
				pre1 = pre
			}

			pre = cur
			cur = next
		}
	}
	return head
}
