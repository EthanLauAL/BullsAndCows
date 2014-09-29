package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Puzzle interface {
	Try(try [N]int) (a,b int)
}

type RandomicPuzzle struct {
	secret [N]int
}

func NewRandomicPuzzle(all [][N]int) Puzzle {
	rand.Seed(time.Now().UnixNano())
	return &RandomicPuzzle{
		all[rand.Int() % len(all)],
	}
}

func (p *RandomicPuzzle) Try(try [N]int) (a,b int) {
	return score(try, p.secret)
}

type ScaningPuzzle struct {
}

func NewScanningPuzzle() Puzzle {
	return &ScaningPuzzle{}
}

func (p *ScaningPuzzle) Try(try [N]int) (a,b int) {
	fmt.Print(try, " A B: ")
	fmt.Scan(&a,&b)
	return
}
