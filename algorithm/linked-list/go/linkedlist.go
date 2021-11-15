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

// 加一个dumy节点
func (node *Node) ReverseNearby1() *Node {
	dumy := CreateNode(0, node)
	cur := dumy

	for {
		if cur == nil || cur.Next == nil || cur.Next.Next == nil {
			break
		} else {
			cur.Next, cur.Next.Next, cur.Next.Next.Next = cur.Next.Next, cur.Next.Next.Next, cur.Next
			cur = cur.Next.Next
		}
	}

	return dumy.Next
}

func (node *Node) Circle() bool {
	slow := node
	fast := node
	for {
		if slow.Next == nil || fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
}

func (node *Node) Middle() *Node {
	slow := node
	fast := node

	for {
		if slow.Next == nil || fast.Next == nil || fast.Next.Next == nil {
			return slow
		}

		slow = slow.Next
		fast = fast.Next.Next
	}
}

func (node *Node) LastN(n int) *Node {
	slow := node
	fast := node

	for i := 1; i < n; i++ {
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next
	}

	for {
		if slow.Next == nil || fast.Next == nil {
			return slow
		}
		slow = slow.Next
		fast = fast.Next
	}
}

func Merge(node1, node2 *Node) *Node {
	var head, insert *Node
	if node1.Value > node2.Value {
		head, insert = node2, node1
	} else {
		head, insert = node1, node2
	}

	cur := head
	for {
		if insert == nil {
			return head
		}

		if cur.Next == nil {
			cur.Next = insert
			return head
		}

		if cur.Next.Value < insert.Value {
			cur = cur.Next
		} else {
			cur.Next, insert.Next, insert = insert, cur.Next, insert.Next
		}
	}
}
