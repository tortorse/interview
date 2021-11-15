package linkedlist

import "testing"

func TestCreateNode(t *testing.T) {
	val := 1
	node := CreateNode(val, nil)
	if node == nil {
		t.Fatal("node should not be nil")
	}

	if node.Value != val {
		t.Fatalf("node value should be %d", val)
	}
}

func TestString(t *testing.T) {
	node1 := CreateNode(1, nil)
	want1 := "1"
	str1 := node1.String()
	if str1 != want1 {
		t.Fatalf("one node to string shoud be %s, indeed return %s", want1, str1)
	}

	node2 := CreateNode(2, node1)
	want2 := "2 -> 1"
	str2 := node2.String()
	if str2 != want2 {
		t.Fatalf("two node to string should be %s, indeed return %s", want2, str2)
	}

	node3 := CreateNode(3, node2)
	want3 := "3 -> 2 -> 1"
	str3 := node3.String()
	if str3 != want3 {
		t.Fatalf("three node to string should be %s, indeed return %s", want3, str3)
	}
}

func TestReverse(t *testing.T) {
	wants := []struct {
		Node *Node
		Str  string
	}{
		{
			Node: CreateNode(1, nil),
			Str:  "1",
		},
		{
			Node: CreateNode(2, CreateNode(1, nil)),
			Str:  "1 -> 2",
		},
		{
			Node: CreateNode(3, CreateNode(2, CreateNode(1, nil))),
			Str:  "1 -> 2 -> 3",
		},
	}

	for i, want := range wants {
		str := want.Node.Reverse().String()
		if str != want.Str {
			t.Fatalf("return%d %s want %s", i, str, want.Str)
		}
	}
}

func TestReverse1(t *testing.T) {
	wants := []struct {
		Node *Node
		Str  string
	}{
		{
			Node: CreateNode(1, nil),
			Str:  "1",
		},
		{
			Node: CreateNode(2, CreateNode(1, nil)),
			Str:  "1 -> 2",
		},
		{
			Node: CreateNode(3, CreateNode(2, CreateNode(1, nil))),
			Str:  "1 -> 2 -> 3",
		},
	}

	for i, want := range wants {
		str := want.Node.Reverse1().String()
		if str != want.Str {
			t.Fatalf("return%d %s want %s", i, str, want.Str)
		}
	}
}

func TestReverseNearby(t *testing.T) {
	wants := []struct {
		Node *Node
		Str  string
	}{
		{
			Node: CreateNode(1, nil),
			Str:  "1",
		},
		{
			Node: CreateNode(2, CreateNode(1, nil)),
			Str:  "1 -> 2",
		},
		{
			Node: CreateNode(3, CreateNode(2, CreateNode(1, nil))),
			Str:  "2 -> 3 -> 1",
		},
		{
			Node: CreateNode(4, CreateNode(3, CreateNode(2, CreateNode(1, nil)))),
			Str:  "3 -> 4 -> 1 -> 2",
		},
	}

	for i, want := range wants {
		str := want.Node.ReverseNearby1().String()
		if str != want.Str {
			t.Fatalf("return%d %s want %s", i, str, want.Str)
		}
	}
}

func TestCircle(t *testing.T) {
	node1 := CreateNode(1, nil)
	node2 := CreateNode(2, node1)
	node3 := CreateNode(3, node2)
	node1.Next = node3

	wants := []struct {
		Node   *Node
		Circle bool
	}{
		{
			Node:   CreateNode(1, nil),
			Circle: false,
		},
		{
			Node:   CreateNode(2, CreateNode(1, nil)),
			Circle: false,
		},
		{
			Node:   CreateNode(3, CreateNode(2, CreateNode(1, nil))),
			Circle: false,
		},
		{
			Node:   node3,
			Circle: true,
		},
	}

	for i, want := range wants {
		if want.Node.Circle() != want.Circle {
			t.Fatalf("Circle error %d", i)
		}
	}
}

func TestMiddle(t *testing.T) {
	wants := []struct {
		node   *Node
		middle int
	}{
		{
			node:   CreateNode(1, nil),
			middle: 1,
		},
		{
			node:   CreateNode(2, CreateNode(1, nil)),
			middle: 2,
		},
		{
			node:   CreateNode(3, CreateNode(2, CreateNode(1, nil))),
			middle: 2,
		},
	}

	for i, want := range wants {
		if want.node.Middle().Value != want.middle {
			t.Fatalf("middle errror %d want %d", i, want.middle)
		}
	}
}
