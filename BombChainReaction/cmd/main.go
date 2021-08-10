package main

import (
	"fmt"

	"github.com/SaffatHasan/AlgorithmPractice/BombChainReaction/pkg/algorithm"
	"github.com/SaffatHasan/AlgorithmPractice/BombChainReaction/pkg/model"
)

func main() {
	bombs := []model.Bomb{
		{0, 0, 1},
		{0, 1, 0},
	}
	startBombIndex := 0

	for index, bomb := range bombs {
		fmt.Printf("Bomb %d: %+v\n", index, bomb)
	}

	fmt.Printf("Exploded %d total bombs.\n", algorithm.FindTotalBombsExploded(bombs, startBombIndex))
}
