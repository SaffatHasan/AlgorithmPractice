package algorithm

import (
	"github.com/SaffatHasan/AlgorithmPractice/BombChainReaction/pkg/model"
)

type BombGraph map[model.Bomb][]model.Bomb

// see README.md for problem description
func FindTotalBombsExploded(bombs []model.Bomb, startBombIndex int) int {
	return FindTotalBombsExplodedHelper(bombs, startBombIndex, false)
}

func FindTotalBombsExplodedHelper(bombs []model.Bomb, startBombIndex int, useChannels bool) int {
	if startBombIndex > len(bombs) || startBombIndex < 0 {
		return 0
	}
	start := bombs[startBombIndex]
	graph := BuildGraph(bombs)
	visited := make(map[model.Bomb]bool)

	if useChannels {
		BFSWithChannels(start, graph, visited)
	} else {
		BFS(start, graph, visited)
	}
	return len(visited)
}

func BuildGraph(bombs []model.Bomb) BombGraph {
	result := make(map[model.Bomb][]model.Bomb)

	for i := 0; i < len(bombs); i++ {
		b1 := bombs[i]
		for j := i + 1; j < len(bombs); j++ {
			b2 := bombs[j]

			distance := int(b1.Distance(b2))

			if b1.R >= distance {
				result[b1] = append(result[b1], b2)
			}

			if b2.R >= distance {
				result[b2] = append(result[b2], b1)
			}
		}
	}
	return result
}

func BFS(start model.Bomb, graph BombGraph, visited map[model.Bomb]bool) {
	queue := make([]model.Bomb, 0)

	queue = append(queue, start)

	for len(queue) > 0 {
		// pop
		current := queue[0]
		visited[current] = true
		queue = queue[1:]

		for _, neighbor := range graph[current] {
			if visited[neighbor] {
				continue
			}
			queue = append(queue, neighbor)
		}
	}
}

// This deadlocks needs more research
func BFSWithChannels(start model.Bomb, graph BombGraph, visited map[model.Bomb]bool) {
	queue := make(chan model.Bomb)

	queue <- start

	select {
	case current := <-queue:
		for _, neighbor := range graph[current] {
			if visited[neighbor] {
				continue
			}
			visited[neighbor] = true
			queue <- neighbor
		}
	default:
		close(queue)
		return
	}
}
