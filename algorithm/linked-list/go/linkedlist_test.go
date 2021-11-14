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
