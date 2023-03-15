package binarytree

import "fmt"

type node struct {
	val         int
	left, right *node
}

type BinaryTree struct {
	root *node
}

func (binaryTree *BinaryTree) Add(val int) {
	if binaryTree.root == nil {
		binaryTree.root = &node{val: val}
		return
	}
	binaryTree.add(binaryTree.root, val)
}

func (binaryTree *BinaryTree) add(root *node, val int) {
	if root.val >= val {
		if root.left == nil {
			root.left = &node{val: val}
			return
		}
		binaryTree.add(root.left, val)
	} else {
		if root.right == nil {
			root.right = &node{val: val}
			return
		}
		binaryTree.add(root.right, val)
	}
}

func (binaryTree *BinaryTree) Remove(val int) {
	binaryTree.remove(binaryTree.root, val)
}

func (binaryTree *BinaryTree) remove(root *node, val int) *node {
	if root == nil {
		return nil
	}
	if val == root.val {
		if root.right == nil && root.left == nil {
			root = nil
		} else if root.right == nil {
			root = root.left
		} else if root.left == nil {
			root = root.right
		} else {
			root.val = binaryTree.rightMostElementAndRemove(root.left)
			root.left = binaryTree.remove(root.left, root.val)
		}

	} else if val < root.val {
		root.left = binaryTree.remove(root.left, val)
	} else {
		root.right = binaryTree.remove(root.right, val)
	}
	return root
}

func (binaryTree *BinaryTree) rightMostElementAndRemove(root *node) int {
	tmp := root
	for tmp.right != nil {
		tmp = tmp.right
	}
	i := tmp.val
	return i
}

func (binaryTree *BinaryTree) Search(val int) node {
	tmp := binaryTree.search(binaryTree.root, val)
	return node{val: tmp.val}
}

func (binaryTree *BinaryTree) search(root *node, val int) *node {
	if root == nil {
		return root
	}
	if val == root.val {
		return root
	} else if val < root.val {
		return binaryTree.search(root.left, val)
	} else {
		return binaryTree.search(root.right, val)
	}
}

func (binaryTree *BinaryTree) Print() {
	binaryTree.printInOrder(binaryTree.root)
	fmt.Println("")
	binaryTree.printPreOrder(binaryTree.root)
	fmt.Println("")
	binaryTree.printPostOrder(binaryTree.root)
	fmt.Println("")
}

func (binaryTree *BinaryTree) printInOrder(root *node) {
	if root == nil {
		return
	}
	binaryTree.printInOrder(root.left)
	fmt.Print(root.val, " ")
	binaryTree.printInOrder(root.right)

}

func (binaryTree *BinaryTree) printPreOrder(root *node) {
	if root == nil {
		return
	}
	fmt.Print(root.val, " ")
	binaryTree.printPreOrder(root.left)
	binaryTree.printPreOrder(root.right)
}

func (binaryTree *BinaryTree) printPostOrder(root *node) {
	if root == nil {
		return
	}
	binaryTree.printPostOrder(root.left)
	binaryTree.printPostOrder(root.right)
	fmt.Print(root.val, " ")
}
