package main

import (
	"fmt"

	"github.com/SaffatHasan/AlgorithmPractice/BombChainReaction/pkg/algorithm"
	"github.com/SaffatHasan/AlgorithmPractice/BombChainReaction/pkg/model"
)

func main() {
	bombs := []model.Bomb{
		{X: 0, Y: 0, R: 5},
		{X: 0, Y: 1, R: 5},
	}
	startBombIndex := 0

	for index, bomb := range bombs {
		fmt.Printf("Bomb %d: %+v\n", index, bomb)
	}

	fmt.Printf("Exploded %d total bombs.\n", algorithm.FindTotalBombsExploded(bombs, startBombIndex))
}
