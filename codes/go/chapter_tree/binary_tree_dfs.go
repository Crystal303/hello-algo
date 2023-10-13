// File: binary_tree_dfs.go
// Created Time: 2022-11-26
// Author: Reanon (793584285@qq.com)

package chapter_tree

import (
	list "container/list"

	. "github.com/krahets/hello-algo/pkg"
)

var nums []any

/* 前序遍历 */
func preOrder(node *TreeNode) {
	if node == nil {
		return
	}
	// 访问优先级：根节点 -> 左子树 -> 右子树
	nums = append(nums, node.Val)
	preOrder(node.Left)
	preOrder(node.Right)
}

/* 中序遍历 */
func inOrder(node *TreeNode) {
	if node == nil {
		return
	}
	// 访问优先级：左子树 -> 根节点 -> 右子树
	inOrder(node.Left)
	nums = append(nums, node.Val)
	inOrder(node.Right)
}

/* 后序遍历 */
func postOrder(node *TreeNode) {
	if node == nil {
		return
	}
	// 访问优先级：左子树 -> 右子树 -> 根节点
	postOrder(node.Left)
	postOrder(node.Right)
	nums = append(nums, node.Val)
}

/* 前序遍历迭代 */
func preOrderIterate(node *TreeNode, data *[]any) {
	if node == nil {
		return
	}

	// 根 -> 左 -> 右
	stack := list.New()
	stack.PushBack(node)
	for 0 < stack.Len() {
		tmp := stack.Remove(stack.Back()).(*TreeNode)
		*data = append(*data, tmp.Val)
		if tmp.Right != nil {
			stack.PushBack(tmp.Right)
		}
		if tmp.Left != nil {
			stack.PushBack(tmp.Left)
		}
	}
}

func postOrderIterate(node *TreeNode, data *[]any) {
	if node == nil {
		return
	}
	var prev *TreeNode

	// 左 -> 右 -> 根
	stack := list.New()
	currentNode := node
	for currentNode != nil || stack.Len() > 0 {
		for currentNode != nil {
			stack.PushBack(currentNode)
			currentNode = currentNode.Left
		}

		if stack.Len() > 0 {
			tmpNode := stack.Back().Value.(*TreeNode)
			if tmpNode.Right != nil && tmpNode.Right != prev {
				currentNode = tmpNode.Right
				continue
			}

			stack.Remove(stack.Back())
			*data = append(*data, tmpNode.Val)
			prev = tmpNode
		}
	}
}

func inOrderIterate(node *TreeNode, data *[]any) {
	if node == nil {
		return
	}

	// 左 -> 根 -> 右
	stack := list.New()
	root := node
	for root != nil || stack.Len() > 0 {
		for root != nil && root.Left != nil {
			stack.PushBack(root)
			root = root.Left
		}
		if root != nil {
			stack.PushBack(root)
		}

		if stack.Len() > 0 {
			currentNode := stack.Remove(stack.Back()).(*TreeNode)
			*data = append(*data, currentNode.Val)
			root = currentNode.Right
		}
	}
}

func postOrderIterateV2(node *TreeNode, data *[]any) {
	if node == nil {
		return
	}
	// 左 -> 右 -> 根
	stack := list.New()
	root := node
	var prev *TreeNode
	for root != nil || stack.Len() > 0 {
		for root != nil {
			stack.PushBack(root)
			root = root.Left
		}

		if stack.Len() > 0 {
			currentNode := stack.Back().Value.(*TreeNode)
			if currentNode.Right == nil || currentNode.Right == prev {
				stack.Remove(stack.Back())
				*data = append(*data, currentNode.Val)

				prev = currentNode
				continue
			}

			root = currentNode.Right
		}
	}
}
