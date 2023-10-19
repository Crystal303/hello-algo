package chapter_heap

import (
	"cmp"
	"container/heap"
)

type MinHeap[T cmp.Ordered] []T

func (m *MinHeap[T]) Len() int {
	return len(*m)
}

func (m *MinHeap[T]) Less(i, j int) bool {
	tmp := *m
	return tmp[i] < tmp[j]
}

func (m *MinHeap[T]) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *MinHeap[T]) Push(x any) {
	*m = append(*m, x.(T))
}

func (m *MinHeap[T]) Pop() any {
	node := (*m)[m.Len()-1]
	*m = (*m)[:m.Len()-1]
	return node
}

func (m *MinHeap[T]) Peek() T {
	return (*m)[0]
}

func (m *MinHeap[T]) toAny() []any {
	res := make([]any, 0, m.Len())
	for _, item := range *m {
		res = append(res, item)
	}
	return res
}

func TopKHeap[T cmp.Ordered](nums []T, k int) MinHeap[T] {
	// build heap
	if len(nums) < k {
		k = len(nums)
	}
	mh := new(MinHeap[T])
	*mh = append(*mh, nums[:k]...)
	heap.Init(mh)

	for ; k < len(nums); k++ {
		if mh.Peek() < nums[k] {
			heap.Pop(mh)
			heap.Push(mh, nums[k])
		}
	}
	return *mh
}
