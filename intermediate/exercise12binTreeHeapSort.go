package intermediate

import (
	"fmt"
	"log"
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

	tree := &binTree{&binTreeNode{1,
		&binTreeNode{6,
			&binTreeNode{13,
				&binTreeNode{14, nil, nil},
				&binTreeNode{12, nil, nil},
			},
			&binTreeNode{11, nil, nil},
		}, &binTreeNode{54,
			&binTreeNode{2,
				&binTreeNode{16, nil, nil},
				&binTreeNode{11, nil, nil},
			},
			&binTreeNode{13, nil, nil},
		}}}

	recursiveDFSPrintTree(tree.root)
	fmt.Println() // recursive print doesn't leave a newline TODO fix?
	fmt.Println("number of levels: ", countLeftMostLevels(tree.root, 0))
	prettyPrintTree(tree.root)
	parentStack := []*binTreeNode{tree.root, tree.root.left}
	heapify(tree.root.left.left, parentStack)
	prettyPrintTree(tree.root)
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

func prettyPrintTree(root *binTreeNode) {
	numLevels := countLeftMostLevels(root, 0)
	curLevel := 0
	levelNodes := make(map[int][]*binTreeNode)
	levelNodes[0] = append(levelNodes[0], root)
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

func heapify(node *binTreeNode, parentStack []*binTreeNode) {
	largestNode := node
	// if either child is larger than largest so far
	//fmt.Println("root val ", node.value)
	if node.left != nil && node.left.value > largestNode.value {
		largestNode = node.left
	}
	if node.right != nil && node.right.value > largestNode.value {
		largestNode = node.right
	}
	if largestNode != node {
		// fmt.Printf("change left '%d' with largest node '%d'\n",
		// 	node.value, largestNode.value)
		buf := node.value
		node.value = largestNode.value
		largestNode.value = buf
		// fmt.Printf("after change left '%d' with largest node '%d'\n",
		// 	node.value, largestNode.value)
		// prettyPrintTree(largestNode)
		if len(parentStack) > 0 {
			heapify(parentStack[len(parentStack)-1], parentStack[:len(parentStack)-1])
		}
	}
}

func nextNodeParent(root *binTreeNode) *binTreeNode {
	numLevels := countLeftMostLevels(root, 0)
	curLevel := 0
	levelNodes := make(map[int][]*binTreeNode)
	levelNodes[0] = append(levelNodes[0], root)
	// going over each level print the value and check for values to append from below, left to right
	for curLevel < numLevels {
		// figure out how much padding to print here
		for i := 0; i < len(levelNodes[curLevel]); i++ {
			curNode := levelNodes[curLevel][i]
			//fmt.Printf(" %3d", curNode.value)
			if curNode.left != nil || curNode.right != nil {
				return curNode
			}
		}
		// print newline after each level
		fmt.Println()
		// have to increment the current level
		curLevel = curLevel + 1
	}
	log.Fatalln("could not get a new node, no fix in sight")
	return nil
}
