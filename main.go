package main

import "fmt"

type BinarySearchTree struct {
	Root *Node
}
type Node struct {
	LeftNode  *Node
	RightNode *Node
	Value     int
}

func main() {
	example1 := []int{12, 11, 90, 82, 7, 9}
	bst := CreateBST(example1)

	printDeepestNode(bst)
}

//Wrapper function for findDeepestNode output to console
func printDeepestNode(bst BinarySearchTree) {
	var value *int
	var deepestDepth *int
	var d int = 0
	var v int = 0
	deepestDepth = &d
	value = &v
	findDeepestNode(bst.Root, value, 0, deepestDepth)
	fmt.Printf("deepest, %d; depth, %d\n", d, v)
}

//Walks the tree from left to right, looking for the deepest node
func findDeepestNode(node *Node, value *int, depth int, deepestDepth *int) {
	if node != nil {
		if depth > *deepestDepth {
			*deepestDepth = depth
			*value = node.Value
		}
		findDeepestNode(node.LeftNode, value, depth+1, deepestDepth)
		findDeepestNode(node.RightNode, value, depth+1, deepestDepth)
	}
}

//Performs an in-order walk to output a list of all nodes from smallest to largest on the console
func InOrderTreeWalk(node *Node) {
	if node != nil {
		InOrderTreeWalk(node.LeftNode)
		fmt.Println(node.Value)
		InOrderTreeWalk(node.RightNode)
	}
}

//Creates a BST from an integer array
func CreateBST(input []int) BinarySearchTree {
	bst := BinarySearchTree{}
	var node *Node
	for _, element := range input {
		if bst.Root == nil {
			node := Node{}
			node.Value = element
			bst.Root = &node
		} else {
			node = bst.Root
			inserted := false
			for {
				if inserted {
					break
				}
				if element < node.Value {
					if node.LeftNode == nil {
						leftNode := Node{}
						leftNode.Value = element
						node.LeftNode = &leftNode
						inserted = true
					} else {
						node = node.LeftNode
					}
				} else {
					if node.RightNode == nil {
						rightNode := Node{}
						rightNode.Value = element
						node.RightNode = &rightNode
						inserted = true
					} else {
						node = node.RightNode
					}
				}
			}
		}
	}
	return bst
}
