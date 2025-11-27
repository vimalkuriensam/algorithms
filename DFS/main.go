package main

import (
	"fmt"
)

var visited = make(map[rune]bool)
var graph = make(map[rune][]rune)

func main() {
	graph['A'] = []rune{'B', 'C'}
	graph['B'] = []rune{'D'}
	graph['C'] = []rune{'E'}
	graph['D'] = []rune{}
	graph['E'] = []rune{}
	dfs('A')
}

func dfs(at rune) {
	if visited[at] {
		return
	}
	fmt.Println(string(at))
	visited[at] = true
	neighbors := graph[at]

	for _, neighbor := range neighbors {
		dfs(neighbor)
	}
}
