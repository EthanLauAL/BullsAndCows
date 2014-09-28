package main

import (
	"math/rand"
	"time"
)

type Puzzle struct {
	secret [N]int
}

func NewPuzzle() Puzzle {
	rand.Seed(time.Now().UnixNano())
	var p Puzzle
	for i,_ := range p.secret {
		p.secret[i] = randNum()
	}
	return p
}

func (p *Puzzle) Try(try [N]int) (a,b int) {
	return score(try, p.secret)
}

func randNum() int {
	return rand.Int() % (MaxNum + 1 - MinNum) + MinNum
}
