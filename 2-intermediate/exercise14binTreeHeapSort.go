package intermediate

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type binTreeNode struct {
	value       int
	left, right *binTreeNode
}

// BinTreeHeapSort implements the rules for intermediate exercise 12
// 	Implement a binary heap by implementing a pointer-linked binary tree.
// 	Use it for implementing heap-sort.
func BinTreeHeapSort() {
	input := []int{6, 10, 9, 1, 302, 100, 3, 1025, 513, 257, 129, 65, 33, 17}
	//8, 16, 32, 64, 128, 256, 512, 1024}
	newRoot := binHeapSort(input)
	prettyPrintTree(newRoot)
	for newRoot.left != nil {
		prettyPrintTree(newRoot)
		maxVal := binHeapUnSort(newRoot)
		fmt.Println("heap top number: ", maxVal)
	}
}

func binHeapSort(values []int) *binTreeNode {
	root := &binTreeNode{value: values[0], left: nil, right: nil}
	values = values[1:]
	for _, value := range values {
		newNode := &binTreeNode{value, nil, nil}
		nextNode := nextNodeParent(root)
		//fmt.Println("nextNodeParent: ", nextNode)
		if nextNode.left == nil {
			nextNode.left = newNode
		} else if nextNode.right == nil {
			nextNode.right = newNode
		}
		parentSlice := GetParents(newNode, []*binTreeNode{root})
		//fmt.Println("parent slice: ", parentSlice)
		//prettyPrintTree(root)
		heapify(parentSlice[len(parentSlice)-2], parentSlice[:len(parentSlice)-2])
		//prettyPrintTree(root)
	}
	return root
}

// func recursiveDFSPrintTree(node *binTreeNode) {
// 	fmt.Printf(" %3d", node.value)
// 	if node.left != nil {
// 		recursiveDFSPrintTree(node.left)
// 	}
// 	if node.right != nil {
// 		recursiveDFSPrintTree(node.right)
// 	}
// }

func countLeftMostLevels(node *binTreeNode) int {
	levels, err := countOutsideMostLevels(node, "left", 0)
	if err != nil {
		log.Fatal(err)
	}
	return levels
}

func countRightMostLevels(node *binTreeNode) int {
	levels, err := countOutsideMostLevels(node, "right", 0)
	if err != nil {
		log.Fatal(err)
	}
	return levels
}

func countOutsideMostLevels(node *binTreeNode, direction string, levels int) (int, error) {
	levels = levels + 1
	if direction == "left" {
		if node.left != nil {
			levels, err := countOutsideMostLevels(node.left, "left", levels)
			if err != nil {
				log.Fatal(err)
			}
			return levels, nil
		}
		return levels, nil
	} else if direction == "right" {
		if node.right != nil {
			levels, err := countOutsideMostLevels(node.right, "right", levels)
			if err != nil {
				log.Fatal(err)
			}
			return levels, nil
		}
		return levels, nil
	}
	return levels, errors.New("countOutsideMostLevels expects string direction 'right' or 'left'")
}

func prettyPrintTree(root *binTreeNode) {
	numLevels := countLeftMostLevels(root)
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
	//fmt.Println("heapify largest node", largestNode.value)
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
	numLevels := countLeftMostLevels(root)
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
	//fmt.Println("value of target node: ", target.value, "parent stack: ", parentStack)
	if parentStack[len(parentStack)-1] != nil {
		left := parentStack[len(parentStack)-1].left
		right := parentStack[len(parentStack)-1].right
		if left == target {
			return append(parentStack, left)
		} else if right == target {
			return append(parentStack, right)
		} else {
			leftParentStack := GetParents(target, append(parentStack, left))
			if len(leftParentStack) > 1 {
				return leftParentStack
			}
			rightParentStack := GetParents(target, append(parentStack, right))
			if len(rightParentStack) > 1 { //len(parentStack) {
				return rightParentStack
			}
		}

	}
	return []*binTreeNode{nil}
}

func binHeapUnSort(root *binTreeNode) int {
	lastElement := getLastElement(root)
	parents := GetParents(lastElement, []*binTreeNode{root})
	maxVal := root.value
	root.value = lastElement.value
	lastElementParent := parents[len(parents)-2]
	if lastElementParent.right != nil {
		lastElementParent.right = nil
	} else if lastElementParent.left != nil {
		lastElementParent.left = nil
	} else {
		log.Fatal("last parent doesn't have a non nil node?")
	}
	lastElement = getLastElement(root)
	parents = GetParents(lastElement, []*binTreeNode{root})
	heapify(lastElement, parents)
	return maxVal
}

func getLastElement(node *binTreeNode) *binTreeNode {
	leftLevels, rightLevels := countLeftMostLevels(node), countRightMostLevels(node)
	if node.left != nil && node.right == nil {
		return node.left
	}
	if leftLevels == rightLevels {
		return getRightMost(node)
	}
	return getLastElement(node.left)
}

func getRightMost(node *binTreeNode) *binTreeNode {
	if node.right != nil {
		return getRightMost(node.right)
	}
	return node
}
