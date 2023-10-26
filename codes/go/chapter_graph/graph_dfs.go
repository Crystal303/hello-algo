// File: graph_dfs.go
// Created Time: 2023-02-18
// Author: Reanon (793584285@qq.com)

package chapter_graph

import (
	. "github.com/krahets/hello-algo/pkg"
)

/* 深度优先遍历 DFS 辅助函数 */
func dfs(g *graphAdjList, visited map[Vertex]struct{}, res *[]Vertex, vet Vertex) {
	// append 操作会返回新的的引用，必须让原引用重新赋值为新slice的引用
	*res = append(*res, vet)
	visited[vet] = struct{}{}
	// 遍历该顶点的所有邻接顶点
	for _, adjVet := range g.adjList[vet] {
		_, isExist := visited[adjVet]
		// 递归访问邻接顶点
		if !isExist {
			dfs(g, visited, res, adjVet)
		}
	}
}

/* 深度优先遍历 DFS */
// 使用邻接表来表示图，以便获取指定顶点的所有邻接顶点
func graphDFS(g *graphAdjList, startVet Vertex) []Vertex {
	// 顶点遍历序列
	res := make([]Vertex, 0)
	// 哈希表，用于记录已被访问过的顶点
	visited := make(map[Vertex]struct{})
	dfs(g, visited, &res, startVet)
	// 返回顶点遍历序列
	return res
}

func graphDFSV2(g *graphAdjList, startVet Vertex) []Vertex {
	res := make([]Vertex, 0)
	if _, ok := g.adjList[startVet]; !ok {
		return res
	}

	visited := make(map[Vertex]struct{}, 0)
	stack := make([]Vertex, 0)
	stack = append(stack, startVet)
	for 0 < len(stack) {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		res = append(res, current)
		visited[current] = struct{}{}

		for i := len(g.adjList[current]) - 1; 0 <= i; i-- {
			if _, ok := visited[g.adjList[current][i]]; !ok {
				stack = append(stack, g.adjList[current][i])
			}
		}
	}

	return res
}
