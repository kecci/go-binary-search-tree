package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	value int
	left  *node
	right *node
}

type bst struct {
	root *node
	size int
}

func (n node) String() string {
	return strconv.Itoa(n.value)
}

func (b bst) String() string {
	sb := strings.Builder{}
	b.inOrderTraversal(&sb)
	return sb.String()
}

func (b bst) inOrderTraversal(sb *strings.Builder) {
	b.inOrderTraversalByNode(sb, b.root)
}

func (b bst) inOrderTraversalByNode(sb *strings.Builder, root *node) {
	if root == nil {
		return
	}
	b.inOrderTraversalByNode(sb, root.left)
	sb.WriteString(fmt.Sprintf("%s", root))
	b.inOrderTraversalByNode(sb, root.right)
}

func (b bst) add(value int) {
	b.root = b.addByNode(b.root, value)
	b.size++
}

func (b bst) addByNode(root *node, value int) *node {
	if root == nil {
		return &node{value: value}
	}

	if value < root.value {
		root.left = b.addByNode(root.left, value)
	} else {
		root.right = b.addByNode(root.right, value)
	}

	return root
}

func (b bst) search(value int) (*node, bool) {
	return b.searchByNode(b.root, value)
}

func (b bst) searchByNode(root *node, value int) (*node, bool) {
	if root == nil {
		return nil, false
	}

	if value == root.value {
		return root, true
	} else if value < root.value {
		return b.searchByNode(root.left, value)
	} else {
		return b.searchByNode(root.right, value)
	}
}

func (b bst) remove(value int) {
	b.root = b.removeByNode(b.root, value)
}

func (b bst) removeByNode(root *node, value int) *node {
	if root == nil {
		return root
	}

	if value > root.value {
		root.right = b.removeByNode(root.right, value)
	} else if value < root.value {
		root.left = b.removeByNode(root.left, value)
	} else {
		if root.left == nil {
			return root.right
		} else {
			temp := root.left
			for temp.right != nil {
				temp = temp.right
			}
			root.value = temp.value

			root.left = b.removeByNode(root.left, value)
		}
	}

	return root
}

func main() {
	// Setup
	n := &node{1, nil, nil}
	n.left = &node{0, nil, nil}
	n.right = &node{2, nil, nil}

	b := bst{n, 3}
	fmt.Println(b)

	// Add new data
	b.add(5)
	b.add(4)
	b.add(6)

	fmt.Println(b)

	// Search data
	fmt.Println(b.search(6))
	fmt.Println(b.search(3))

	// Remove
	b.remove(6)
	fmt.Println(b)
	b.remove(3)
	fmt.Println(b)
}
