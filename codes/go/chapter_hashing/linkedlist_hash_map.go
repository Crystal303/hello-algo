package chapter_hashing

import (
	"fmt"
	"strconv"
	"strings"
)

/* 链式地址哈希表 */
type linkedlistHashMap struct {
	size        int     // 键值对数量
	capacity    int     // 哈希表容量
	loadThres   float64 // 触发扩容的负载因子阈值
	extendRatio int     // 扩容倍数
	buckets     []*node // 桶数组
}

type node struct {
	Next *node
	Val  pair
}

/* 构造方法 */
func newLinkedlistHashMap() *linkedlistHashMap {
	buckets := make([]*node, 4)
	for i := 0; i < 4; i++ {
		buckets[i] = new(node)
	}
	return &linkedlistHashMap{
		size:        0,
		capacity:    4,
		loadThres:   2.0 / 3.0,
		extendRatio: 2,
		buckets:     buckets,
	}
}

/* 哈希函数 */
func (m *linkedlistHashMap) hashFunc(key int) int {
	return key % m.capacity
}

/* 负载因子 */
func (m *linkedlistHashMap) loadFactor() float64 {
	return float64(m.size / m.capacity)
}

/* 查询操作 */
func (m *linkedlistHashMap) get(key int) string {
	idx := m.hashFunc(key)
	bucket := m.buckets[idx]
	// 遍历桶，若找到 key 则返回对应 val
	for bucket.Next != nil {
		bucket = bucket.Next
		if bucket.Val.key == key {
			return bucket.Val.val
		}
	}
	// 若未找到 key 则返回空字符串
	return ""
}

/* 添加操作 */
func (m *linkedlistHashMap) put(key int, val string) {
	// 当负载因子超过阈值时，执行扩容
	if m.loadFactor() > m.loadThres {
		m.extend()
	}
	idx := m.hashFunc(key)
	// 遍历桶，若遇到指定 key ，则更新对应 val 并返回
	bucket := m.buckets[idx]
	for bucket.Next != nil {
		bucket = bucket.Next
		if bucket.Val.key == key {
			bucket.Val.val = val
			return
		}
	}
	// 若无该 key ，则将键值对添加至尾部
	p := pair{
		key: key,
		val: val,
	}
	for bucket.Next != nil {
		bucket = bucket.Next
	}
	bucket.Next = &node{
		Next: nil,
		Val:  p,
	}
	m.size += 1
}

/* 删除操作 */
func (m *linkedlistHashMap) remove(key int) {
	idx := m.hashFunc(key)
	// 遍历桶，从中删除键值对
	bucket := m.buckets[idx]
	for bucket.Next != nil {
		if bucket.Next.Val.key == key {
			bucket.Next = bucket.Next.Next
			m.size -= 1
			break
		}
	}
}

/* 扩容哈希表 */
func (m *linkedlistHashMap) extend() {
	// 暂存原哈希表
	tmpBuckets := make([]*node, len(m.buckets))
	for i := 0; i < len(m.buckets); i++ {
		tmpBuckets[i] = m.buckets[i]
	}
	// 初始化扩容后的新哈希表
	m.capacity *= m.extendRatio
	m.buckets = make([]*node, m.capacity)
	for i := 0; i < m.capacity; i++ {
		m.buckets[i] = new(node)
	}
	m.size = 0

	// 将键值对从原哈希表搬运至新哈希表
	for _, bucket := range tmpBuckets {
		for bucket.Next != nil {
			bucket = bucket.Next
			m.put(bucket.Val.key, bucket.Val.val)
		}
	}
}

/* 打印哈希表 */
func (m *linkedlistHashMap) print() {
	var builder strings.Builder

	for _, bucket := range m.buckets {
		builder.WriteString("[")
		for bucket.Next != nil {
			bucket = bucket.Next
			builder.WriteString(strconv.Itoa(bucket.Val.key) + " -> " + bucket.Val.val + " ")
		}
		builder.WriteString("]")
		fmt.Println(builder.String())
		builder.Reset()
	}
}
