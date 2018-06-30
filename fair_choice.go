package main

import (
	"fmt"
	"math"
)

func main() {
	algorithms := []algorithm{basic{}, byTwo{}, recursive{}}
	fmt.Printf("numPlayers,basic,byTwo,recursive,\n")
	for numPlayers := 1; numPlayers <= 100; numPlayers++ {
		fmt.Printf("%v,", numPlayers)
		for _, a := range algorithms {
			fmt.Printf("%v,", diff(numPlayers, a))
		}
		fmt.Printf("\n")
	}
}

func diff(numPlayers int, a algorithm) int {
	score := make(map[bool]int)
	for i, choice := range a.choose(numPlayers) {
		score[choice] += i
	}
	return score[false] - score[true]
}

type algorithm interface {
	choose(numPlayers int) []bool
}

type basic struct {}

func (basic) choose(numPlayers int) []bool {
	var choice []bool
	for i := 0; i < numPlayers; i++ {
		choice = append(choice, false, true)
	}
	return choice
}

type byTwo struct {}

func (byTwo) choose(numPlayers int) []bool {
	var choice []bool
	start := false
	for i := 0; i < numPlayers; i++ {
		choice = append(choice, start, !start)
		start = !start
	}
	return choice
}

type recursive struct {}

func (recursive) choose(numPlayers int) []bool {
	return chooseRecursive(int(math.Ceil(math.Log2(float64(numPlayers * 2))) - 1), false)[:numPlayers * 2]
}

func chooseRecursive(level int, start bool) []bool {
	if level <= 0 {
		return []bool{start, !start}
	}
	return append(chooseRecursive(level - 1, start), chooseRecursive(level - 1, !start)...)
}
