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
