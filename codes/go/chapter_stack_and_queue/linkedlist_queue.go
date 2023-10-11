// File: linkedlist_queue.go
// Created Time: 2022-11-28
// Author: Reanon (793584285@qq.com)

package chapter_stack_and_queue

import (
	"container/list"
)

/* 基于链表实现的队列 */
type linkedListQueue struct {
	*linkedListStack
}

/* 初始化队列 */
func newLinkedListQueue() *linkedListQueue {
	return &linkedListQueue{
		linkedListStack: newLinkedListStack(),
	}
}

/* 入队 */
func (s *linkedListQueue) push(value int) {
	s.linkedListStack.push(value)
}

/* 出队 */
func (s *linkedListQueue) pop() any {
	if s.isEmpty() {
		return nil
	}
	node := s.head.Next
	if s.head.Next != nil {
		s.head.Next = s.head.Next.Next
	}
	s.count--

	if s.isEmpty() {
		s.head.Next = nil
		s.tail.Next = nil
	}
	return node
}

/* 访问队首元素 */
func (s *linkedListQueue) peek() any {
	if s.isEmpty() {
		return nil
	}
	return s.head.Next
}

/* 获取队列的长度 */
func (s *linkedListQueue) size() int {
	return s.linkedListStack.size()
}

/* 判断队列是否为空 */
func (s *linkedListQueue) isEmpty() bool {
	return s.linkedListStack.isEmpty()
}

/* 获取 List 用于打印 */
func (s *linkedListQueue) toList() *list.List {
	return s.linkedListStack.toList()
}
