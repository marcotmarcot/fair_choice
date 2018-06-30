package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	algorithms := []algorithm{basic{}, byTwo{}, recursive{}}
	total := make([]int, 3, 3)
	for execution := 0; execution <= 10000; execution++ {
		for numPlayers := 5; numPlayers <= 5; numPlayers++ {
			players := makePlayers(numPlayers)
			for i, a := range algorithms {
				d := diff(players, numPlayers, a)
				if d < 0 {
					d *= -1
				}
				total[i] += d
			}
		}
	}
	for _, t := range total {
		fmt.Println(t)
	}
}

func makePlayers(numPlayers int) []int {
	var players []int
	for i := 0; i < numPlayers * 2; i++ {
		players = append(players, rand.Int() % 1000)
	}
	sort.Ints(players)
	return players
}

func diff(players []int, numPlayers int, a algorithm) int {
	score := make(map[bool]int)
	for i, choice := range a.choose(numPlayers) {
		score[choice] += players[i] * players[i] * players[i]
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
