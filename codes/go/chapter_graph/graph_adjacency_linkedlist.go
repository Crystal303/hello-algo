// File: graph_adjacency_list.go
// Created Time: 2023-01-31
// Author: Reanon (793584285@qq.com)

package chapter_graph

import (
	"fmt"
	"strconv"
	"strings"

	. "github.com/krahets/hello-algo/pkg"
)

type VertexNode struct {
	Vertex
	next *VertexNode
}

/* 基于邻接表实现的无向图类 */
type graphAdjLinkedlist struct {
	// 邻接表，key: 顶点，value：该顶点的所有邻接顶点
	adjList map[Vertex]*VertexNode
}

/* 构造函数 */
func newGraphAdjLinkedlist(edges [][]Vertex) *graphAdjLinkedlist {
	g := &graphAdjLinkedlist{
		adjList: make(map[Vertex]*VertexNode, 0),
	}
	// 添加所有顶点和边
	for _, edge := range edges {
		if len(edge) == 2 {
			g.addVertex(edge[0])
			g.addVertex(edge[1])
			g.addEdge(edge[0], edge[1])
		}
	}
	return g
}

/* 获取顶点数量 */
func (g *graphAdjLinkedlist) size() int {
	return len(g.adjList)
}

/* 添加边 */
func (g *graphAdjLinkedlist) addEdge(vet1 Vertex, vet2 Vertex) {
	_, ok1 := g.adjList[vet1]
	_, ok2 := g.adjList[vet2]
	if !ok1 || !ok2 || vet1 == vet2 {
		panic("error")
	}
	// 添加边 vet1 - vet2, 添加匿名 struct{},
	g.adjList[vet1] = &VertexNode{
		Vertex: vet2,
		next:   g.adjList[vet1],
	}
	g.adjList[vet2] = &VertexNode{
		Vertex: vet1,
		next:   g.adjList[vet2],
	}
}

/* 删除边 */
func (g *graphAdjLinkedlist) removeEdge(vet1 Vertex, vet2 Vertex) {
	_, ok1 := g.adjList[vet1]
	_, ok2 := g.adjList[vet2]
	if !ok1 || !ok2 || vet1 == vet2 {
		panic("error")
	}
	// 删除边 vet1 - vet2\
	g.adjList[vet1] = deleteVertexNode(g.adjList[vet1], vet2)
	g.adjList[vet2] = deleteVertexNode(g.adjList[vet2], vet1)
}

func deleteVertexNode(node *VertexNode, vertex Vertex) *VertexNode {
	if node == nil {
		return nil
	}
	root := node
	var prev *VertexNode
	for node != nil {
		if node.Vertex == vertex {
			if prev != nil {
				prev.next = node.next
			} else {
				return root.next
			}
		}
		prev = node
		node = node.next
	}
	return root
}

/* 添加顶点 */
func (g *graphAdjLinkedlist) addVertex(vet Vertex) {
	_, ok := g.adjList[vet]
	if ok {
		return
	}
	// 在邻接表中添加一个新链表
	g.adjList[vet] = nil
}

/* 删除顶点 */
func (g *graphAdjLinkedlist) removeVertex(vet Vertex) {
	_, ok := g.adjList[vet]
	if !ok {
		panic("error")
	}
	// 在邻接表中删除顶点 vet 对应的链表
	delete(g.adjList, vet)
	// 遍历其他顶点的链表，删除所有包含 vet 的边
	for v, node := range g.adjList {
		g.adjList[v] = deleteVertexNode(node, vet)
	}
}

/* 打印邻接表 */
func (g *graphAdjLinkedlist) print() {
	var builder strings.Builder
	fmt.Printf("邻接表 = \n")
	for k, v := range g.adjList {
		builder.WriteString("\t\t" + strconv.Itoa(k.Val) + ": ")
		for v != nil {
			builder.WriteString(strconv.Itoa(v.Val) + " ")
			v = v.next
		}
		fmt.Println(builder.String())
		builder.Reset()
	}
}
