package intermediate

import "fmt"

// Implement a binary heap by implementing a pointer-linked binary tree.
// Use it for implementing heap-sort.

type binTree struct {
	root *binTreeNode
}

type binTreeNode struct {
	value       int
	left, right *binTreeNode
}

func main() {
	fmt.Println("exercise 12...")
	tree := &binTree{
		&binTreeNode{15,
			&binTreeNode{8,
				&binTreeNode{3,
					nil,
					nil,
				},
				&binTreeNode{1,
					nil,
					nil,
				},
			}, &binTreeNode{5,
				&binTreeNode{2,
					nil,
					nil,
				},
				nil,
			},
		},
	}

	printTree(tree.root)
}

func printTree(node *binTreeNode) {
	fmt.Printf("%3d", node.value)
	if node.left != nil {
		printTree(node.left)
	}
	if node.right != nil {
		printTree(node.right)
	}
}
