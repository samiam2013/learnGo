package intermediate

import (
	"fmt"
	"strings"
)

type binTree struct {
	root *binTreeNode
}

type binTreeNode struct {
	value       int
	left, right *binTreeNode
}

// BinTreeHeapSort implements the rules for intermediate exercise 12
// 	Implement a binary heap by implementing a pointer-linked binary tree.
// 	Use it for implementing heap-sort.
func BinTreeHeapSort() {
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

	recursiveDFSPrintTree(tree.root)
	fmt.Println() // recursive print doesn't leave a newline TODO fix?
	fmt.Println("number of levels: ", countLeftMostLevels(tree.root, 0))
	prettyPrintTree(tree)
}

func recursiveDFSPrintTree(node *binTreeNode) {
	fmt.Printf(" %3d", node.value)
	if node.left != nil {
		recursiveDFSPrintTree(node.left)
	}
	if node.right != nil {
		recursiveDFSPrintTree(node.right)
	}
}

func countLeftMostLevels(node *binTreeNode, levels int) int {
	levels = levels + 1
	if node.left != nil {
		levels = countLeftMostLevels(node.left, levels)
	}
	return levels
}

func prettyPrintTree(tree *binTree) {
	numLevels := countLeftMostLevels(tree.root, 0)
	curLevel := 0
	levelNodes := make(map[int][]*binTreeNode)
	levelNodes[0] = append(levelNodes[0], tree.root)
	// going over each level print the value and check for values to append from below, left to right
	for curLevel < numLevels {
		// figure out how much padding to print here
		fmt.Print(strings.Repeat(" ", 4*(numLevels-(curLevel+1))))
		for i := 0; i < len(levelNodes[curLevel]); i++ {
			curNode := levelNodes[curLevel][i]
			fmt.Printf(" %3d", curNode.value)
			if curNode.left != nil {
				levelNodes[curLevel+1] = append(levelNodes[curLevel+1], curNode.left)
			}
			if curNode.right != nil {
				levelNodes[curLevel+1] = append(levelNodes[curLevel+1], curNode.right)
			}
		}
		// print newline after each level
		fmt.Println()
		// have to increment the current level
		curLevel = curLevel + 1
	}
}
