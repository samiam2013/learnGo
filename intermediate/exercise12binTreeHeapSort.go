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

	// tree := &binTree{&binTreeNode{1,
	// 	&binTreeNode{6,
	// 		&binTreeNode{13, nil, nil},
	// 		&binTreeNode{11, nil, nil},
	// 	}, &binTreeNode{54,
	// 		&binTreeNode{2, nil, nil},
	// 		&binTreeNode{13, nil, nil},
	// 	}}}

	//recursiveDFSPrintTree(tree.root)
	// fmt.Println() // recursive print doesn't leave a newline TODO fix?
	// fmt.Println("number of levels: ", countLeftMostLevels(tree.root, 0))
	// prettyPrintTree(tree.root)
	// heapify(tree.root.left, parentStack)

	// parentStack := []*binTreeNode{tree.root}
	// parentStack = GetParents(nextNodeParent(tree.root), []*binTreeNode{tree.root})
	// for _, parent := range parentStack {
	// 	if parent != nil {
	// 		fmt.Println("stack parent value: ", parent.value)
	// 	}
	// }

	newRoot := binHeapSort([]int{6, 10, 9, 1, 302, 2})
	fmt.Println("final outpout because golang is a bitch and won't let me have an unused value:")
	prettyPrintTree(newRoot)

}

func binHeapSort(values []int) *binTreeNode {
	root := &binTreeNode{value: values[0], left: nil, right: nil}
	values = values[1:]
	for _, value := range values {
		newNode := &binTreeNode{value, nil, nil}
		nextNode := nextNodeParent(root)
		fmt.Println("nextNodeParent: ", nextNode)
		if nextNode.left == nil {
			nextNode.left = newNode
		} else if nextNode.right == nil {
			nextNode.right = newNode
		}
		parentSlice := GetParents(newNode, []*binTreeNode{root})
		fmt.Println("parent slice: ", parentSlice)
		prettyPrintTree(root)
		heapify(root, parentSlice)
		prettyPrintTree(root)
	}
	return root
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
	fmt.Println("heapify largest node", largestNode.value)
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
			// fmt.Printf("curnode %3d left %d right %d\n",
			// 	curNode.value, curNode.left.value, curNode.right.value)
			if curNode.left == nil || curNode.right == nil {
				return curNode
			}
			if curNode.left != nil {
				levelNodes[curLevel+1] = append(levelNodes[curLevel+1], curNode.left)
			}
			if curNode.right != nil {
				levelNodes[curLevel+1] = append(levelNodes[curLevel+1], curNode.right)
			}

		}
		// print newline after each level
		//fmt.Println()
		// have to increment the current level
		curLevel = curLevel + 1
	}
	log.Fatalln("could not get a new node, no fix in sight")
	return nil
}

// GetParents recursivley searches for the node handed back
//	by nextNodeParent()
func GetParents(target *binTreeNode, parentStack []*binTreeNode) []*binTreeNode {
	fmt.Println("value of target node: ", target.value, "parent stack: ", parentStack)
	if parentStack[len(parentStack)-1] != nil {
		left := parentStack[len(parentStack)-1].left
		right := parentStack[len(parentStack)-1].right
		if left == target {
			return append(parentStack, left)
		} else {
			leftParentStack := GetParents(target, append(parentStack, left))
			if len(leftParentStack) > len(parentStack) {
				return leftParentStack
			}
		}
		if right == target {
			return append(parentStack, right)
		} else {
			rightParentStack := GetParents(target, append(parentStack, right))
			if len(rightParentStack) > len(parentStack) {
				return rightParentStack
			}
		}

	}
	return []*binTreeNode{nil}
}
