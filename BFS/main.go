package main

import "fmt"

type Graph struct {
	ADJ_List map[int][]int
}

func main() {
	graph := NewGraph()
	graph.Add_Edge(0, 2)
	graph.Add_Edge(0, 1)
	graph.Add_Edge(1, 2)
	graph.Add_Edge(1, 3)
	graph.Add_Edge(2, 4)
	graph.Add_Edge(3, 4)

	graph.BFS(2)
}

func NewGraph() *Graph {
	return &Graph{
		ADJ_List: make(map[int][]int),
	}
}

func (g *Graph) Add_Edge(u, v int) {
	g.ADJ_List[u] = append(g.ADJ_List[u], v)
	g.ADJ_List[v] = append(g.ADJ_List[v], u)
}

func (g *Graph) BFS(start int) {
	queue := []int{start}
	visited := make(map[int]bool)
	for len(queue) > 0 {
		current_node := Dequeu(&queue)
		if !visited[current_node] {
			visited[current_node] = true
			Visit(current_node)
			for _, neighbour := range g.ADJ_List[current_node] {
				Enqueu(&queue, neighbour)
			}
		}
	}
}

func Enqueu(queue *[]int, node int) {
	*queue = append(*queue, node)
}

func Dequeu(queue *[]int) int {
	node := (*queue)[0]
	*queue = (*queue)[1:]
	return node
}

func Visit(node int) {
	fmt.Printf("%d\t", node)
}
