package chapter_stack_and_queue

import (
	"container/list"

	. "github.com/krahets/hello-algo/pkg"
)

/* 基于链表实现的栈 */
type linkedListStack struct {
	// 使用内置包 list 来实现栈
	head, tail *ListNode
	count      int
}

/* 初始化栈 */
func newLinkedListStack() *linkedListStack {
	return &linkedListStack{
		head:  new(ListNode),
		tail:  new(ListNode),
		count: 0,
	}
}

/* 入栈 */
func (s *linkedListStack) push(value int) {
	node := &ListNode{
		Next: nil,
		Val:  value,
	}
	s.tail.Next = node

	head := s.head
	for head.Next != nil {
		head = head.Next
	}
	head.Next = node
	s.count++
}

/* 出栈 */
func (s *linkedListStack) pop() any {
	if s.isEmpty() {
		return nil
	}

	head := s.head
	for head.Next != nil {
		if head.Next.Next == nil {
			s.tail.Next = head
			node := head.Next
			head.Next = nil
			s.count--

			return node
		}
		head = head.Next
	}
	return nil
}

/* 访问栈顶元素 */
func (s *linkedListStack) peek() any {
	if s.isEmpty() {
		return nil
	}
	return s.tail.Next
}

/* 获取栈的长度 */
func (s *linkedListStack) size() int {
	return s.count
}

/* 判断栈是否为空 */
func (s *linkedListStack) isEmpty() bool {
	return s.count == 0
}

/* 获取 List 用于打印 */
func (s *linkedListStack) toList() *list.List {
	l := list.New()
	head := s.head
	for head.Next != nil {
		head = head.Next
		l.PushBack(head.Val)
	}
	return l
}
